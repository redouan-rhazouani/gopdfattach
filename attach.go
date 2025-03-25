package gopdfattach

import (
	"io"

	"github.com/MarlinKuhn/gopdfattach/internal/attach"
)

type AttachConfig struct {
	DocumentType     string // defaults to "INVOICE"
	FileName         string // defaults to "factur-x.xml"
	Version          string // defaults to "1.0" if factur-x, "2p0" if zugferd
	ConformanceLevel string // defaults to "EN 16931"
	Creator          string // defaults to "gopdfattach"
}

func (a *AttachConfig) toConfig() attach.Config {
	if a == nil {
		return attach.Config{}
	}

	return attach.Config{
		DocumentType:     a.DocumentType,
		FileName:         a.FileName,
		Version:          a.Version,
		ConformanceLevel: a.ConformanceLevel,
		Creator:          a.Creator,
	}
}

// AttachZUGFeRD attaches a ZUGFeRD XML file to a PDF document and converts it to a PDF/A-3 document.
func AttachZUGFeRD(zugFeRDXml io.Reader, pdf io.ReadSeeker, config *AttachConfig) ([]byte, error) {
	c := config.toConfig()
	c.XmlType = attach.TypeZugferd
	return attach.Attach(zugFeRDXml, pdf, c)
}

// AttachFacturX attaches a Factur-X XML file to a PDF document and converts it to a PDF/A-3 document.
func AttachFacturX(factorXXml io.Reader, pdf io.ReadSeeker, config *AttachConfig) ([]byte, error) {
	c := config.toConfig()
	c.XmlType = attach.TypeFacturX
	return attach.Attach(factorXXml, pdf, c)
}
