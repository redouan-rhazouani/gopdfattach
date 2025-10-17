package gopdfattach

import (
	"io"

	"github.com/MarlinKuhn/gopdfattach/internal/attach"
)

// AF represents the relationship between the embedded file and the PDF document
// according to the PDF/A-3 specification.
type AF string

// AFRelationship constants define valid relationship types for embedded files.
const (
	// AFAlternative indicates the embedded file is an alternative representation of the PDF content.
	// This is the recommended value for Factur-X/ZUGFeRD hybrid invoices where the XML is a machine-readable
	// alternative to the human-readable PDF.
	AFAlternative AF = "Alternative"

	// AFData indicates the embedded file contains additional data related to the PDF.
	AFData AF = "Data"

	// AFSource indicates the embedded file is the source file from which the PDF was created.
	AFSource AF = "Source"

	// AFSupplement indicates the embedded file contains supplementary information.
	AFSupplement AF = "Supplement"
)

type AttachConfig struct {
	DocumentType     string // defaults to "INVOICE"
	FileName         string // defaults to "factur-x.xml"
	Version          string // defaults to "1.0" if factur-x, "2p0" if zugferd
	ConformanceLevel string // defaults to "EN 16931"
	Creator          string // defaults to "gopdfattach"
	AFRelationship   AF     // defaults to AFAlternative (spec-compliant for Factur-X/ZUGFeRD)
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
		AFRelationship:   string(a.AFRelationship),
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
