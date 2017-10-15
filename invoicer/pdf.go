package invoicer

import (
		"bytes"
		"github.com/SebastiaanKlippert/go-wkhtmltopdf"
		"log"
)

func getPDFContent(htmlContent *string) []byte {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Dpi.Set(72)
	pdfg.ImageDpi.Set(72)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// Add one page from an URL
	parseHtmlContent(&*htmlContent)
	buf := []byte(*htmlContent)
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(buf))
	page.DisableSmartShrinking.Set(true)
	page.Zoom.Set(3)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	return pdfg.Bytes()
}