/*
 * Copyright (c) 2025. Marlin Kuhn
 */

package gopdfattach

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttach_WithValidZugFeRD(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	pdfData, err := AttachZUGFeRD(xmlFile, pdfFile, nil)
	assert.NoError(t, err)
	assert.NotNil(t, pdfData)
}

func TestAttach_WithValidFacturX(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	pdfData, err := AttachFacturX(xmlFile, pdfFile, nil)
	assert.NoError(t, err)
	assert.NotNil(t, pdfData)
}

func TestAttach_WithMissingPDF(t *testing.T) {
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	pdfData, err := AttachFacturX(xmlFile, nil, nil)
	assert.Error(t, err)
	assert.Nil(t, pdfData)
}

func TestAttach_WithMissingXML(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()

	pdfData, err := AttachFacturX(nil, pdfFile, nil)
	assert.Error(t, err)
	assert.Nil(t, pdfData)
}

func TestAttach_WithInvalidPDF(t *testing.T) {
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()
	invalidPDF := bytes.NewReader([]byte("invalid pdf content"))

	pdfData, err := AttachFacturX(xmlFile, invalidPDF, nil)
	assert.Error(t, err)
	assert.Nil(t, pdfData)
}

func Test_ZugFeRD(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	config := &AttachConfig{
		DocumentType:     "INVOICE",
		FileName:         "factur-x.xml",
		Version:          "2.0",
		ConformanceLevel: "EN 16931",
	}

	pdfData, err := AttachZUGFeRD(xmlFile, pdfFile, config)
	assert.NoError(t, err)
	assert.NotNil(t, pdfData)

	xml, infos, err := Extract(bytes.NewReader(pdfData))
	assert.NoError(t, err)
	assert.NotNil(t, xml)
	assert.Equal(t, FileTypeZugferd, infos.FileType)
	assert.Equal(t, config.DocumentType, infos.DocumentType)
	assert.Equal(t, config.FileName, infos.FileName)
	assert.Equal(t, config.Version, infos.Version)
	assert.Equal(t, config.ConformanceLevel, infos.ConformanceLevel)
}

func Test_FacturX(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	config := &AttachConfig{
		DocumentType:     "INVOICE",
		FileName:         "factur-x.xml",
		Version:          "1.0",
		ConformanceLevel: "EN 16931",
	}

	pdfData, err := AttachFacturX(xmlFile, pdfFile, config)
	assert.NoError(t, err)
	assert.NotNil(t, pdfData)

	xml, infos, err := Extract(bytes.NewReader(pdfData))
	assert.NoError(t, err)
	assert.NotNil(t, xml)
	assert.Equal(t, FileTypeFacturX, infos.FileType)
	assert.Equal(t, config.DocumentType, infos.DocumentType)
	assert.Equal(t, config.FileName, infos.FileName)
	assert.Equal(t, config.Version, infos.Version)
	assert.Equal(t, config.ConformanceLevel, infos.ConformanceLevel)
}

func TestAttach_WithCustomAFRelationship(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	config := &AttachConfig{
		DocumentType:     "INVOICE",
		FileName:         "factur-x.xml",
		Version:          "1.0",
		ConformanceLevel: "EN 16931",
		AFRelationship:   AFData, // Custom value
	}

	pdfData, err := AttachFacturX(xmlFile, pdfFile, config)
	assert.NoError(t, err)
	assert.NotNil(t, pdfData)

	// Verify the PDF can be extracted successfully
	xml, infos, err := Extract(bytes.NewReader(pdfData))
	assert.NoError(t, err)
	assert.NotNil(t, xml)
	assert.Equal(t, FileTypeFacturX, infos.FileType)
}

func TestAttach_WithAlternativeAFRelationship(t *testing.T) {
	pdfFile, _ := os.Open("testdata/invoice.pdf")
	defer pdfFile.Close()
	xmlFile, _ := os.Open("testdata/factur-x.xml")
	defer xmlFile.Close()

	config := &AttachConfig{
		DocumentType:     "INVOICE",
		FileName:         "factur-x.xml",
		Version:          "1.0",
		ConformanceLevel: "EN 16931",
		AFRelationship:   AFAlternative, // Spec-compliant value
	}

	pdfData, err := AttachZUGFeRD(xmlFile, pdfFile, config)
	assert.NoError(t, err)
	assert.NotNil(t, pdfData)

	// Verify the PDF can be extracted successfully
	xml, infos, err := Extract(bytes.NewReader(pdfData))
	assert.NoError(t, err)
	assert.NotNil(t, xml)
	assert.Equal(t, FileTypeZugferd, infos.FileType)
}
