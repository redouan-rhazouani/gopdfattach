/*
 * Copyright (c) 2025. Marlin Kuhn
 */

package fx

import (
	"fmt"

	"github.com/trimmer-io/go-xmp/xmp"
)

var (
	NsFacturX = xmp.NewNamespace("fx", "urn:factur-x:pdfa:CrossIndustryDocument:invoice:1p0#", NewModel)
)

func init() {
	xmp.Register(NsFacturX, xmp.XmpMetadata)
}

func NewModel(name string) xmp.Model {
	return &CrossIndustryDocument{}
}

func MakeModel(d *xmp.Document) (*CrossIndustryDocument, error) {
	m, err := d.MakeModel(NsFacturX)
	if err != nil {
		return nil, err
	}
	x, _ := m.(*CrossIndustryDocument)
	return x, nil
}

func FindModel(d *xmp.Document) *CrossIndustryDocument {
	if m := d.FindModel(NsFacturX); m != nil {
		return m.(*CrossIndustryDocument)
	}
	return nil
}

type CrossIndustryDocument struct {
	DocumentType     string `xmp:"fx:DocumentType"`
	DocumentFileName string `xmp:"fx:DocumentFileName"`
	Version          string `xmp:"fx:Version"`
	ConformanceLevel string `xmp:"fx:ConformanceLevel"`
}

func (x CrossIndustryDocument) Can(nsName string) bool {
	return NsFacturX.GetName() == nsName
}

func (x CrossIndustryDocument) Namespaces() xmp.NamespaceList {
	return xmp.NamespaceList{NsFacturX}
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
		return "", fmt.Errorf("%s: %v", NsFacturX.GetName(), err)
	} else {
		return v, nil
	}
}

func (x *CrossIndustryDocument) SetTag(tag, value string) error {
	if err := xmp.SetNativeField(x, tag, value); err != nil {
		return fmt.Errorf("%s: %v", NsFacturX.GetName(), err)
	}
	return nil
}
