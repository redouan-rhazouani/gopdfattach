/*
 * Copyright (c) 2025. Marlin Kuhn
 */

package attach

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	_ "github.com/MarlinKuhn/gopdfattach/internal/xsd"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/fx"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/pdfaExtension"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/pdfaid"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/zf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate"
	pdf2 "github.com/trimmer-io/go-xmp/models/pdf"
	"github.com/trimmer-io/go-xmp/xmp"
)

func init() {
	model.ConfigPath = "disabled"
}

type fileType string

const (
	TypeZugferd fileType = "zugferd"
	TypeFacturX fileType = "factur-x"
)

type Config struct {
	XmlType          fileType
	DocumentType     string
	FileName         string
	Version          string
	ConformanceLevel string
	Creator          string
}

func (c *Config) setDefaults() {
	if c.XmlType == "" {
		c.XmlType = TypeFacturX
	}

	if c.FileName == "" {
		c.FileName = "factur-x.xml"
	}

	if c.DocumentType == "" {
		c.DocumentType = "INVOICE"
	}

	if c.ConformanceLevel == "" {
		c.ConformanceLevel = "EN 16931"
	}

	if c.Creator == "" {
		c.Creator = "gopdfattach"
	}

	switch c.XmlType {
	case TypeFacturX:
		if c.Version == "" {
			c.Version = "1.0"
		}

	case TypeZugferd:
		if c.Version == "" {
			c.Version = "2p0"
		}
	}
}

func Attach(zugFeRD io.Reader, pdf io.ReadSeeker, config Config) ([]byte, error) {
	if zugFeRD == nil {
		return nil, fmt.Errorf("missing XML file")
	}

	if pdf == nil {
		return nil, fmt.Errorf("missing PDF file")
	}

	config.setDefaults()
	configuration := model.NewDefaultConfiguration()
	ctx, err := api.ReadContext(pdf, configuration)
	if err != nil {
		return nil, fmt.Errorf("could not read PDF file: %w", err)
	}

	// Needs to be done for attachments to work!
	if err = validate.XRefTable(ctx); err != nil {
		return nil, fmt.Errorf("could not validate XRefTable: %w", err)
	}

	catalog, err := ctx.Catalog()
	if err != nil {
		return nil, fmt.Errorf("could not get catalog: %w", err)
	}

	{
		metadata, err := pdfcpu.ExtractMetadata(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not extract metadata: %w", err)
		}

		doc := xmp.NewDocument()

		for _, meta := range metadata {
			rawMetaXMP, err := io.ReadAll(meta)
			if err != nil {
				continue
			}

			// XMP metadata manipulation
			err = xmp.Unmarshal(rawMetaXMP, doc)
			if err == nil {
				break
			}
		}

		info, err := pdf2.MakeModel(doc)
		if err != nil {
			return nil, fmt.Errorf("could not make model: %w", err)
		}

		info.PDFVersion = "1.7"
		info.Keywords = "ZUGFeRD, PDF/A-3"
		info.Creator = xmp.AgentName(config.Creator)
		info.Producer = xmp.AgentName(config.Creator)

		pdfa, err := pdfaid.MakeModel(doc)
		if err != nil {
			return nil, fmt.Errorf("could not make model: %w", err)
		}

		pdfa.Conformance = "U"
		pdfa.Part = "3"

		extension, err := pdfaExtension.MakeModel(doc)
		if err != nil {
			return nil, fmt.Errorf("could not make model: %w", err)
		}

		switch config.XmlType {
		case TypeFacturX:
			makeModel, err := fx.MakeModel(doc)
			if err != nil {
				return nil, fmt.Errorf("could not make model: %w", err)
			}

			makeModel.DocumentType = config.DocumentType
			makeModel.DocumentFileName = config.FileName
			makeModel.Version = config.Version
			makeModel.ConformanceLevel = config.ConformanceLevel
			extension.AddFx()
		case TypeZugferd:
			makeModel, err := zf.MakeModel(doc)
			if err != nil {
				return nil, fmt.Errorf("could not make model: %w", err)
			}

			makeModel.DocumentType = config.DocumentType
			makeModel.DocumentFileName = config.FileName
			makeModel.Version = config.Version
			makeModel.ConformanceLevel = config.ConformanceLevel
			extension.AddZf()
		}

		rawMetaXMP, err := xmp.MarshalIndent(doc, "", "\t")
		if err != nil {
			return nil, fmt.Errorf("could not marshal metadata: %w", err)
		}

		{
			// New XRefModel
			streamDict, err := ctx.XRefTable.NewStreamDictForBuf(rawMetaXMP)
			if err != nil {
				return nil, fmt.Errorf("could not create stream: %w", err)
			}

			streamDict.InsertName("Type", "Metadata")
			streamDict.InsertName("Subtype", "XML")

			err = streamDict.Encode()
			if err != nil {
				return nil, fmt.Errorf("could not encode stream: %w", err)
			}

			indirectRef, err := ctx.XRefTable.IndRefForNewObject(*streamDict)
			if err != nil {
				return nil, fmt.Errorf("could not create indirect reference: %w", err)
			}

			catalog.Update("Metadata", *indirectRef)
		}
	}

	err = attachFileToPfd(ctx, model.Attachment{
		Reader:   zugFeRD,
		ID:       config.FileName,
		FileName: config.FileName,
		Desc:     "Factur-X/ZUGFeRD-Rechnung",
	}, "text/xml")
	if err != nil {
		return nil, fmt.Errorf("could not add attachment: %w", err)
	}

	var data = new(bytes.Buffer)
	err = api.Write(ctx, data, configuration)
	return data.Bytes(), err
}

func attachFileToPfd(ctx *model.Context, a model.Attachment, mimeType string) error {
	xRefTable := ctx.XRefTable
	if err := xRefTable.LocateNameTree("EmbeddedFiles", true); err != nil {
		return err
	}

	catalog, err := ctx.Catalog()
	if err != nil {
		return err
	}

	modTime := time.Now()
	if a.ModTime != nil {
		modTime = *a.ModTime
	}
	sd, err := xRefTable.NewEmbeddedStreamDict(a, modTime)
	if err != nil {
		return err
	}

	streamDict, _, err := xRefTable.DereferenceStreamDict(*sd)
	if err != nil {
		return err
	}

	if mimeType != "" {
		streamDict.InsertName("Subtype", mimeType)
	}

	find, ok := streamDict.Find("Params")
	if ok {
		dict, err := xRefTable.DereferenceDict(find)
		if err != nil {
			return err
		}

		// md5
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, a.Reader); err != nil {
			return err
		}

		bb := buf.Bytes()
		dict.InsertString("CheckSum", getMD5Hash(bb))
	}

	d, err := xRefTable.NewFileSpecDict(a.ID, a.ID, a.Desc, *sd)
	if err != nil {
		return err
	}

	d.InsertName("AFRelationship", "Data")

	ir, err := xRefTable.IndRefForNewObject(d)
	if err != nil {
		return err
	}

	associatedFiles := catalog.ArrayEntry("AF")
	associatedFiles = append(associatedFiles, *ir)
	catalog.Update("AF", associatedFiles)

	m := model.NameMap{a.ID: []types.Dict{d}}

	return xRefTable.Names["EmbeddedFiles"].Add(xRefTable, a.ID, *ir, m, []string{"F", "UF"})
}

func getMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}
