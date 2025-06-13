package services

import "github.com/signintech/gopdf"

func GenerateDisputePDF(filename, body string) error {
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
    pdf.AddPage()
    pdf.SetFont("Helvetica", "", 14)
    pdf.Cell(nil, body)
    return pdf.WritePdf(filename)
}
