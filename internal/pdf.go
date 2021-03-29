package internal

import (
	"io"

	"github.com/phpdave11/gofpdf"
)

func GeneratePdf(file io.Writer, text string) error {

	margin := 23.5
	LineHeight := 1.0

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTopMargin(25.4)
	pdf.SetLeftMargin(margin)
	pdf.SetRightMargin(margin)
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	_, FontSize := pdf.GetFontSize()

	pdf.MultiCell(210-(2*margin), FontSize+LineHeight, text, "", "L", false)

	return pdf.Output(file)
}
