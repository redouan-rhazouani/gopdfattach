/*
 * Copyright (c) 2024-2025. Marlin Kuhn
 */

package zf

import (
	"fmt"

	"github.com/trimmer-io/go-xmp/xmp"
)

var (
	NsZugferd = xmp.NewNamespace("zf", "urn:zugferd:pdfa:CrossIndustryDocument:invoice:2p0#", NewModel)
)

func init() {
	xmp.Register(NsZugferd, xmp.XmpMetadata)
}

func NewModel(name string) xmp.Model {
	return &CrossIndustryDocument{}
}

func MakeModel(d *xmp.Document) (*CrossIndustryDocument, error) {
	m, err := d.MakeModel(NsZugferd)
	if err != nil {
		return nil, err
	}
	x, _ := m.(*CrossIndustryDocument)
	return x, nil
}

func FindModel(d *xmp.Document) *CrossIndustryDocument {
	if m := d.FindModel(NsZugferd); m != nil {
		return m.(*CrossIndustryDocument)
	}
	return nil
}

type CrossIndustryDocument struct {
	DocumentType     string `xmp:"zf:DocumentType"`
	DocumentFileName string `xmp:"zf:DocumentFileName"`
	Version          string `xmp:"zf:Version"`
	ConformanceLevel string `xmp:"zf:ConformanceLevel"`
}

func (x CrossIndustryDocument) Can(nsName string) bool {
	return NsZugferd.GetName() == nsName
}

func (x CrossIndustryDocument) Namespaces() xmp.NamespaceList {
	return xmp.NamespaceList{NsZugferd}
}

func (x *CrossIndustryDocument) SyncModel(d *xmp.Document) error {
	return nil
}

func (x *CrossIndustryDocument) SyncFromXMP(d *xmp.Document) error {
	return nil
}

func (x CrossIndustryDocument) SyncToXMP(d *xmp.Document) error {
	return nil
}

func (x *CrossIndustryDocument) CanTag(tag string) bool {
	_, err := xmp.GetNativeField(x, tag)
	return err == nil
}

func (x *CrossIndustryDocument) GetTag(tag string) (string, error) {
	if v, err := xmp.GetNativeField(x, tag); err != nil {
		return "", fmt.Errorf("%s: %v", NsZugferd.GetName(), err)
	} else {
		return v, nil
	}
}

func (x *CrossIndustryDocument) SetTag(tag, value string) error {
	if err := xmp.SetNativeField(x, tag, value); err != nil {
		return fmt.Errorf("%s: %v", NsZugferd.GetName(), err)
	}
	return nil
}
