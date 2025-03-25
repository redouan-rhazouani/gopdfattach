/*
 * Copyright (c) 2025. Marlin Kuhn
 */

package gopdfattach

import (
	"io"

	"github.com/MarlinKuhn/gopdfattach/internal/extract"
)

const (
	FileTypeZugferd = "ZUGFeRD"
	FileTypeFacturX = "Factur-X"
)

type XMLInfo struct {
	FileType         string
	DocumentType     string
	FileName         string
	Version          string
	ConformanceLevel string
}

// Extract extracts the embedded zugferd or x-rechnung from a PDF. Caution make sure to only use PDFs.
func Extract(pdf io.ReadSeeker) (xml []byte, infos *XMLInfo, err error) {
	out, err := extract.FromReader(pdf)
	if err != nil {
		return nil, nil, err
	}

	infos = &XMLInfo{
		DocumentType:     out.DocumentType,
		FileName:         out.FileName,
		Version:          out.Version,
		ConformanceLevel: out.ConformanceLevel,
	}

	switch out.FileType {
	case extract.Zugferd:
		infos.FileType = FileTypeZugferd
	case extract.FacturX:
		infos.FileType = FileTypeFacturX
	}

	return out.Data, infos, nil
}
