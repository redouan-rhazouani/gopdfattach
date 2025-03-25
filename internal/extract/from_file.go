/*
 * Copyright (c) 2025. Marlin Kuhn
 */

package extract

import (
	"fmt"
	"io"

	_ "github.com/MarlinKuhn/gopdfattach/internal/xsd"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/fx"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/zf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/validate"
	"github.com/trimmer-io/go-xmp/xmp"
)

type fileType int

const (
	Zugferd fileType = iota
	FacturX
)

type Output struct {
	FileType         fileType
	DocumentType     string
	FileName         string
	Version          string
	ConformanceLevel string
	Data             []byte
}

// FromReader extracts the embedded zugferd or x-rechnung from a PDF.
func FromReader(reader io.ReadSeeker) (*Output, error) {
	ctx, err := api.ReadContext(reader, model.NewDefaultConfiguration())
	if err != nil {
		return nil, fmt.Errorf("could not read PDF file: %w", err)
	}

	// Needs to be done for attachments to work!
	if err = validate.XRefTable(ctx); err != nil {
		return nil, err
	}

	metadata, err := pdfcpu.ExtractMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not extract metadata: %w", err)
	}

	var out Output

	for _, meta := range metadata {
		rawMetaXMP, err := io.ReadAll(meta)
		if err != nil {
			continue
		}

		// XMP metadata manipulation
		var doc xmp.Document
		err = xmp.Unmarshal(rawMetaXMP, &doc)
		if err != nil {
			continue
		}

		makeModel := fx.FindModel(&doc)
		if makeModel == nil {
			zfModel := zf.FindModel(&doc)
			if zfModel == nil {
				return nil, fmt.Errorf("could not find fx: %w", err)
			}

			out.FileName = zfModel.DocumentFileName
			out.DocumentType = zfModel.DocumentType
			out.ConformanceLevel = zfModel.ConformanceLevel
			out.Version = zfModel.Version
			out.FileType = Zugferd
		} else {
			out.FileName = makeModel.DocumentFileName
			out.DocumentType = makeModel.DocumentType
			out.ConformanceLevel = makeModel.ConformanceLevel
			out.Version = makeModel.Version
			out.FileType = FacturX
		}

		break
	}

	if out.FileName == "" {
		return nil, fmt.Errorf("could not find file name")
	}

	embeddedFiles, err := ctx.ExtractAttachments([]string{out.FileName})
	if err != nil {
		return nil, fmt.Errorf("could not extract attachments: %w", err)
	}

	if len(embeddedFiles) == 0 {
		return nil, fmt.Errorf("no %s files found", out.FileName)
	}

	out.Data, err = io.ReadAll(embeddedFiles[0].Reader)
	if err != nil {
		return nil, fmt.Errorf("could not read attachment: %w", err)
	}

	return &out, nil
}
