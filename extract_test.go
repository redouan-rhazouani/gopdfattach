/*
 * Copyright (c) 2025. Marlin Kuhn
 */

package gopdfattach

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultFileName = "factur-x.xml"
)

const (
	BasicFolder     = "testdata/BASIC"
	BasicWLFolder   = "testdata/BASIC WL"
	EN16931Folder   = "testdata/EN16931"
	ExtendedFolder  = "testdata/EXTENDED"
	MinimumFolder   = "testdata/MINIMUM"
	XRechnungFolder = "testdata/XRECHNUNG"
)

func TestExtractBasic(t *testing.T) {
	dirEntries, err := os.ReadDir(BasicFolder)
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
		return
	}

	for _, entry := range dirEntries {
		filePath := filepath.Join(BasicFolder, entry.Name())
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(filePath, ".pdf") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			pdfFile, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
				return
			}
			defer pdfFile.Close()

			xml, infos, err := Extract(pdfFile)
			assert.NoError(t, err)
			assert.NotNil(t, xml)
			assert.NotNil(t, infos)
			assert.Equal(t, defaultFileName, infos.FileName)
			assert.Equal(t, FileTypeFacturX, infos.FileType)
			assert.Equal(t, "INVOICE", infos.DocumentType)
			assert.Equal(t, "1.0", infos.Version)
			assert.Equal(t, "BASIC", infos.ConformanceLevel)
		})

	}
}

func TestExtractBasicWL(t *testing.T) {
	dirEntries, err := os.ReadDir(BasicWLFolder)
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
		return
	}

	for _, entry := range dirEntries {
		filePath := filepath.Join(BasicWLFolder, entry.Name())
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(filePath, ".pdf") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			pdfFile, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
				return
			}
			defer pdfFile.Close()

			xml, infos, err := Extract(pdfFile)
			assert.NoError(t, err)
			assert.NotNil(t, xml)
			assert.NotNil(t, infos)
			assert.Equal(t, defaultFileName, infos.FileName)
			assert.Equal(t, FileTypeFacturX, infos.FileType)
			assert.Equal(t, "INVOICE", infos.DocumentType)
			assert.Equal(t, "1.0", infos.Version)
			assert.Equal(t, "BASIC WL", infos.ConformanceLevel)
		})
	}
}

func TestExtractEN16931(t *testing.T) {
	dirEntries, err := os.ReadDir(EN16931Folder)
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
		return
	}

	for _, entry := range dirEntries {
		filePath := filepath.Join(EN16931Folder, entry.Name())
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(filePath, ".pdf") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			pdfFile, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
				return
			}
			defer pdfFile.Close()

			xml, infos, err := Extract(pdfFile)
			assert.NoError(t, err)
			assert.NotNil(t, xml)
			assert.NotNil(t, infos)
			assert.Equal(t, defaultFileName, infos.FileName)
			assert.Equal(t, FileTypeFacturX, infos.FileType)
			assert.Equal(t, "INVOICE", infos.DocumentType)
			assert.Equal(t, "1.0", infos.Version)
			assert.Equal(t, "EN 16931", infos.ConformanceLevel)
		})
	}
}

func TestExtractExtended(t *testing.T) {
	dirEntries, err := os.ReadDir(ExtendedFolder)
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
		return
	}

	for _, entry := range dirEntries {
		filePath := filepath.Join(ExtendedFolder, entry.Name())
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(filePath, ".pdf") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			pdfFile, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
				return
			}
			defer pdfFile.Close()

			xml, infos, err := Extract(pdfFile)
			assert.NoError(t, err)
			assert.NotNil(t, xml)
			assert.NotNil(t, infos)
			assert.Equal(t, defaultFileName, infos.FileName)
			assert.Equal(t, FileTypeFacturX, infos.FileType)
			assert.Equal(t, "INVOICE", infos.DocumentType)
			assert.Equal(t, "1.0", infos.Version)
			assert.Equal(t, "EXTENDED", infos.ConformanceLevel)
		})
	}
}

func TestExtractMinimum(t *testing.T) {
	dirEntries, err := os.ReadDir(MinimumFolder)
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
		return
	}

	for _, entry := range dirEntries {
		filePath := filepath.Join(MinimumFolder, entry.Name())
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(filePath, ".pdf") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			pdfFile, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
				return
			}
			defer pdfFile.Close()

			xml, infos, err := Extract(pdfFile)
			assert.NoError(t, err)
			assert.NotNil(t, xml)
			assert.NotNil(t, infos)
			assert.Equal(t, defaultFileName, infos.FileName)
			assert.Equal(t, FileTypeFacturX, infos.FileType)
			assert.Equal(t, "INVOICE", infos.DocumentType)
			assert.Equal(t, "1.0", infos.Version)
			assert.Equal(t, "MINIMUM", infos.ConformanceLevel)
		})
	}
}

func TestExtractXRechnung(t *testing.T) {
	dirEntries, err := os.ReadDir(XRechnungFolder)
	if err != nil {
		t.Fatalf("failed to read directory: %v", err)
		return
	}

	for _, entry := range dirEntries {
		filePath := filepath.Join(XRechnungFolder, entry.Name())
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(filePath, ".pdf") {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			pdfFile, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
				return
			}
			defer pdfFile.Close()

			xml, infos, err := Extract(pdfFile)
			assert.NoError(t, err)
			assert.NotNil(t, xml)
			assert.NotNil(t, infos)
			assert.Equal(t, "xrechnung.xml", infos.FileName)
			assert.Equal(t, FileTypeFacturX, infos.FileType)
			assert.Equal(t, "INVOICE", infos.DocumentType)
			assert.Equal(t, "2.1", infos.Version)
			assert.Equal(t, "XRECHNUNG", infos.ConformanceLevel)
		})
	}
}
