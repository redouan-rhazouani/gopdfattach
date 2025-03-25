/*
 * Copyright (c) 2024-2025. Marlin Kuhn
 */

package pdfaExtension

import "github.com/trimmer-io/go-xmp/xmp"

type Schema struct {
	Schema       string       `xmp:"pdfaSchema:schema"`
	NamespaceURI string       `xmp:"pdfaSchema:namespaceURI"`
	Prefix       string       `xmp:"pdfaSchema:prefix"`
	Property     PropertyList `xmp:"pdfaSchema:property"`
}

type SchemaList []Schema

func (x *SchemaList) UnmarshalText(data []byte) error {
	// TODO: need samples
	return nil
}

func (x SchemaList) Typ() xmp.ArrayType {
	return xmp.ArrayTypeUnordered
}

func (x SchemaList) MarshalXMP(e *xmp.Encoder, node *xmp.Node, m xmp.Model) error {
	return xmp.MarshalArray(e, node, x.Typ(), x)
}

func (x *SchemaList) UnmarshalXMP(d *xmp.Decoder, node *xmp.Node, m xmp.Model) error {
	return xmp.UnmarshalArray(d, node, x.Typ(), x)
}

type Property struct {
	Name        string `xmp:"pdfaProperty:name"`
	ValueType   string `xmp:"pdfaProperty:valueType"`
	Category    string `xmp:"pdfaProperty:category"`
	Description string `xmp:"pdfaProperty:description"`
}

type PropertyList []Property

func (x *PropertyList) UnmarshalText(data []byte) error {
	return nil
}

func (x PropertyList) Typ() xmp.ArrayType {
	return xmp.ArrayTypeOrdered
}

func (x PropertyList) MarshalXMP(e *xmp.Encoder, node *xmp.Node, m xmp.Model) error {
	return xmp.MarshalArray(e, node, x.Typ(), x)
}

func (x *PropertyList) UnmarshalXMP(d *xmp.Decoder, node *xmp.Node, m xmp.Model) error {
	return xmp.UnmarshalArray(d, node, x.Typ(), x)
}
