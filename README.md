# Go PDFAttach

A Go library for working with electronic invoice attachments in PDF files. Supports ZUGFeRD and Factur-X XML formats by attaching them to PDFs (converting to PDF/A-3) and extracting them from PDF files.

## Installation

```bash
go get github.com/MarlinKuhn/gopdfattach
```

## Usage

### Attaching XML to PDF

You can attach either ZUGFeRD or Factur-X XML to a PDF:

```go
package main

import (
    "os"
    "github.com/MarlinKuhn/gopdfattach"
)

func main() {
    // Open the PDF and XML files
    pdfFile, _ := os.Open("invoice.pdf")
    xmlFile, _ := os.Open("invoice.xml")
    defer pdfFile.Close()
    defer xmlFile.Close()
    
    // Default configuration
    pdfData, err := gopdfattach.AttachZUGFeRD(xmlFile, pdfFile, nil)
    if err != nil {
        panic(err)
    }
    
    // Write the new PDF to file
    os.WriteFile("invoice-with-zugferd.pdf", pdfData, 0644)
    
    // Or use Factur-X with custom configuration
    pdfFile.Seek(0, 0) // Reset position
    xmlFile.Seek(0, 0) 
    
    config := &gopdfattach.AttachConfig{
        DocumentType: "INVOICE",
        ConformanceLevel: "EN 16931",
        Version: "1.0",
        Creator: "MyInvoiceSystem",
    }
    
    pdfData, err = gopdfattach.AttachFacturX(xmlFile, pdfFile, config)
    if err != nil {
        panic(err)
    }
    
    os.WriteFile("invoice-with-facturx.pdf", pdfData, 0644)
}
```

### Extracting XML from PDF

```go
package main

import (
    "fmt"
    "os"
    "github.com/MarlinKuhn/gopdfattach"
)

func main() {
    // Open the PDF file
    pdfFile, _ := os.Open("invoice-with-xml.pdf")
    defer pdfFile.Close()
    
    // Extract the XML data
    xmlData, info, err := gopdfattach.Extract(pdfFile)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Extracted %s XML (version: %s, conformance: %s)\n", 
        info.FileType, info.Version, info.ConformanceLevel)
        
    // Save the extracted XML
    os.WriteFile(info.FileName, xmlData, 0644)
}
```

## Configuration Options

When attaching XML files, you can customize the process with `AttachConfig`:

```go
type AttachConfig struct {
    DocumentType     string // defaults to "INVOICE"
    FileName         string // defaults to "factur-x.xml"
    Version          string // defaults to "1.0" if factur-x, "2p0" if zugferd
    ConformanceLevel string // defaults to "EN 16931"
    Creator          string // defaults to "gopdfattach"
    AFRelationship   string // defaults to "Alternative" (spec-compliant for Factur-X/ZUGFeRD)
}
```

### AFRelationship Field

The `AFRelationship` field specifies how the embedded XML file relates to the PDF document according to the PDF/A-3 specification. 

**Default value**: `"Alternative"` (recommended for Factur-X/ZUGFeRD)

This default follows the Factur-X/ZUGFeRD specification, which states that the embedded XML is an alternative representation of the PDF content (the XML is machine-readable, while the PDF is human-readable, both representing the same invoice).

**Other valid values**:
- `"Data"` - Additional data (general attachments)
- `"Source"` - Source file (original source documents)
- `"Supplement"` - Supplementary information (supporting documents)

For most Factur-X/ZUGFeRD use cases, you should use the default `"Alternative"` value or leave it unset.

## Return Types

### Extract Function

The `Extract` function returns:
- `[]byte`: The raw XML data
- `*XMLInfo`: Metadata about the extracted XML
- `error`: Any error encountered during extraction

```go
type XMLInfo struct {
    FileType         string // Either "ZUGFeRD" or "Factur-X"
    DocumentType     string // Usually "INVOICE"
    FileName         string // Original filename of the attachment
    Version          string // Standard version
    ConformanceLevel string // Conformance level of the XML
}
```

## Error Handling

Both attachment and extraction functions return detailed errors that should be checked in your code to handle common issues like invalid PDFs, missing XML files, or format errors.