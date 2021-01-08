package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	
    pdf := gopdf.GoPdf{}
    mm6ToPx := 22.68
    
    // Base trim-box
    pdf.Start(gopdf.Config{
        PageSize: *gopdf.PageSizeA4, //595.28, 841.89 = A4
        TrimBox: gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
    })

    // Page trim-box
    opt := gopdf.PageOption{
        PageSize: gopdf.PageSizeA4, //595.28, 841.89 = A4
        TrimBox: &gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
    }
    pdf.AddPageWithOption(opt)

    if err := pdf.AddTTFFont("wts11", "./ttf/wts11.ttf"); err != nil {
        log.Print(err.Error())
        return
    }
    
    if err := pdf.SetFont("wts11", "", 14); err != nil {
        log.Print(err.Error())
        return
    }
	
	pdf.Cell(nil,"Hi")
	pdf.SetX(250) //move current location
	pdf.SetY(200)
	pdf.Cell(nil, "gopher and gopher") //print text

	pdf.SetX(30)
	pdf.SetY(40)
	pdf.Text("Link to example.com")
	pdf.AddExternalLink("http://example.com/", 27.5, 28, 125, 15)

	pdf.SetX(30)
	pdf.SetY(70)
	pdf.Text("Link to second page")
	pdf.AddInternalLink("anchor", 27.5, 58, 120, 15)

	pdf.AddPage()
	pdf.SetX(30)
	pdf.SetY(100)
	pdf.SetAnchor("anchor")
	pdf.Text("Anchor position")

	pdf.Image("./ean13.png", 20, 150, nil) //print image

    pdf.WritePdf("hello.pdf")
}