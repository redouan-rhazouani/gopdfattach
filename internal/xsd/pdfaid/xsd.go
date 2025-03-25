/*
 * Copyright (c) 2024-2025. Marlin Kuhn
 */

package pdfaid

import (
	"fmt"

	"github.com/trimmer-io/go-xmp/xmp"
)

var (
	NsPdfaid = xmp.NewNamespace("pdfaid", "http://www.aiim.org/pdfa/ns/id/", NewModel)
)

func init() {
	xmp.Register(NsPdfaid, xmp.XmpMetadata)
}

func NewModel(name string) xmp.Model {
	return &Pdfaid{}
}

func MakeModel(d *xmp.Document) (*Pdfaid, error) {
	m, err := d.MakeModel(NsPdfaid)
	if err != nil {
		return nil, err
	}
	x, _ := m.(*Pdfaid)
	return x, nil
}

func FindModel(d *xmp.Document) *Pdfaid {
	if m := d.FindModel(NsPdfaid); m != nil {
		return m.(*Pdfaid)
	}
	return nil
}

type Pdfaid struct {
	Part        string `xmp:"pdfaid:part"`
	Conformance string `xmp:"pdfaid:conformance"`
}

func (x Pdfaid) Can(nsName string) bool {
	return NsPdfaid.GetName() == nsName
}

func (x Pdfaid) Namespaces() xmp.NamespaceList {
	return xmp.NamespaceList{NsPdfaid}
}

func (x *Pdfaid) SyncModel(d *xmp.Document) error {
	return nil
}

func (x *Pdfaid) SyncFromXMP(d *xmp.Document) error {
	return nil
}

func (x Pdfaid) SyncToXMP(d *xmp.Document) error {
	return nil
}

func (x *Pdfaid) CanTag(tag string) bool {
	_, err := xmp.GetNativeField(x, tag)
	return err == nil
}

func (x *Pdfaid) GetTag(tag string) (string, error) {
	if v, err := xmp.GetNativeField(x, tag); err != nil {
		return "", fmt.Errorf("%s: %v", NsPdfaid.GetName(), err)
	} else {
		return v, nil
	}
}

func (x *Pdfaid) SetTag(tag, value string) error {
	if err := xmp.SetNativeField(x, tag, value); err != nil {
		return fmt.Errorf("%s: %v", NsPdfaid.GetName(), err)
	}
	return nil
}
