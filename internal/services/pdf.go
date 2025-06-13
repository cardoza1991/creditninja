package services

import (
	"github.com/signintech/gopdf"
)

func RenderPDF(text, path string) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	if err := pdf.AddTTFFont("Arial", "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"); err != nil {
		return err
	}
	if err := pdf.SetFont("Arial", "", 14); err != nil {
		return err
	}
	pdf.Cell(nil, text)
	return pdf.WritePdf(path)
}
