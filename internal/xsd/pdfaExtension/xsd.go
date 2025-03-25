/*
 * Copyright (c) 2024-2025. Marlin Kuhn
 */

package pdfaExtension

import (
	"fmt"

	"github.com/MarlinKuhn/gopdfattach/internal/xsd/fx"
	"github.com/MarlinKuhn/gopdfattach/internal/xsd/zf"
	"github.com/trimmer-io/go-xmp/xmp"
)

var (
	NsPdfaExtension = xmp.NewNamespace("pdfaExtension", "http://www.aiim.org/pdfa/ns/extension/", NewModel)
	NsPdfaFields    = xmp.NewNamespace("pdfaField", "http://www.aiim.org/pdfa/ns/field#", NewModel)
	NsPdfaProperty  = xmp.NewNamespace("pdfaProperty", "http://www.aiim.org/pdfa/ns/property#", NewModel)
	NsPdfaSchema    = xmp.NewNamespace("pdfaSchema", "http://www.aiim.org/pdfa/ns/schema#", NewModel)
	NsPdfaType      = xmp.NewNamespace("pdfaType", "http://www.aiim.org/pdfa/ns/type#", NewModel)
	nslist          = xmp.NamespaceList{
		NsPdfaExtension,
		NsPdfaFields,
		NsPdfaProperty,
		NsPdfaSchema,
		NsPdfaType,
	}
)

func init() {
	for _, ns := range nslist {
		xmp.Register(ns, xmp.XmpMetadata)
	}
}

func NewModel(name string) xmp.Model {
	return &PdfaExtension{}
}

func MakeModel(d *xmp.Document) (*PdfaExtension, error) {
	m, err := d.MakeModel(NsPdfaExtension)
	if err != nil {
		return nil, err
	}
	x, _ := m.(*PdfaExtension)
	return x, nil
}

func FindModel(d *xmp.Document) *PdfaExtension {
	if m := d.FindModel(NsPdfaExtension); m != nil {
		return m.(*PdfaExtension)
	}
	return nil
}

type PdfaExtension struct {
	Schemas SchemaList `xmp:"pdfaExtension:schemas"`
}

func (x *PdfaExtension) AddFx() {
	uri := fx.NsFacturX.URI
	for _, schema := range x.Schemas {
		if schema.NamespaceURI == uri {
			return
		}
	}

	x.Schemas = append(x.Schemas, Schema{
		Schema:       "Factur-X PDFA Extension Schema",
		NamespaceURI: uri,
		Prefix:       fx.NsFacturX.Name,
		Property: PropertyList{
			{
				Name:        "DocumentFileName",
				ValueType:   "Text",
				Category:    "external",
				Description: "name of the embedded XML invoice file",
			},
			{
				Name:        "DocumentType",
				ValueType:   "Text",
				Category:    "external",
				Description: "INVOICE",
			},
			{
				Name:        "Version",
				ValueType:   "Text",
				Category:    "external",
				Description: "The actual version of the ZUGFeRD data",
			},
			{
				Name:        "ConformanceLevel",
				ValueType:   "Text",
				Category:    "external",
				Description: "The conformance level of the ZUGFeRD data",
			},
		},
	})
}

func (x *PdfaExtension) AddZf() {
	uri := zf.NsZugferd.URI
	for _, schema := range x.Schemas {
		if schema.NamespaceURI == uri {
			return
		}
	}

	x.Schemas = append(x.Schemas, Schema{
		Schema:       "Factur-X PDFA Extension Schema",
		NamespaceURI: uri,
		Prefix:       zf.NsZugferd.Name,
		Property: PropertyList{
			{
				Name:        "DocumentFileName",
				ValueType:   "Text",
				Category:    "external",
				Description: "name of the embedded XML invoice file",
			},
			{
				Name:        "DocumentType",
				ValueType:   "Text",
				Category:    "external",
				Description: "INVOICE",
			},
			{
				Name:        "Version",
				ValueType:   "Text",
				Category:    "external",
				Description: "The actual version of the ZUGFeRD data",
			},
			{
				Name:        "ConformanceLevel",
				ValueType:   "Text",
				Category:    "external",
				Description: "The conformance level of the ZUGFeRD data",
			},
		},
	})
}

func (x PdfaExtension) Can(nsName string) bool {
	return NsPdfaExtension.GetName() == nsName
}

func (x PdfaExtension) Namespaces() xmp.NamespaceList {
	return nslist
}

func (x *PdfaExtension) SyncModel(d *xmp.Document) error {
	return nil
}

func (x *PdfaExtension) SyncFromXMP(d *xmp.Document) error {
	// remove schemas with no properties
	for i := 0; i < len(x.Schemas); {
		if len(x.Schemas[i].Property) == 0 {
			x.Schemas = append(x.Schemas[:i], x.Schemas[i+1:]...)
		} else {
			i++
		}
	}
	return nil
}

func (x PdfaExtension) SyncToXMP(d *xmp.Document) error {
	// remove schemas with no properties
	for i := 0; i < len(x.Schemas); {
		if len(x.Schemas[i].Property) == 0 {
			x.Schemas = append(x.Schemas[:i], x.Schemas[i+1:]...)
		} else {
			i++
		}
	}
	return nil
}

func (x *PdfaExtension) CanTag(tag string) bool {
	_, err := xmp.GetNativeField(x, tag)
	return err == nil
}

func (x *PdfaExtension) GetTag(tag string) (string, error) {
	if v, err := xmp.GetNativeField(x, tag); err != nil {
		return "", fmt.Errorf("%s: %v", NsPdfaExtension.GetName(), err)
	} else {
		return v, nil
	}
}

func (x *PdfaExtension) SetTag(tag, value string) error {
	if err := xmp.SetNativeField(x, tag, value); err != nil {
		return fmt.Errorf("%s: %v", NsPdfaExtension.GetName(), err)
	}
	return nil
}
