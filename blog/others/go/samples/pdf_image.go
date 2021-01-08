package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("wts11", "./ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	//use path
	pdf.Image("./qrcode.png", 200, 50, nil)
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	//use image holder by []byte
	imgH1, err := gopdf.ImageHolderByBytes(getImageBytes())
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.ImageByHolder(imgH1, 200, 250, nil)

	//use image holder by io.Reader
	file, err := os.Open("./gopher.jpeg")
	if err != nil {
		log.Print(err.Error())
		return
	}
	imgH2, err := gopdf.ImageHolderByReader(file)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.ImageByHolder(imgH2, 200, 450, nil)

	pdf.SetX(250)
	pdf.SetY(200)
	pdf.Cell(nil, "gopher and gopher")

	pdf.WritePdf("image.pdf")
}

func getImageBytes() []byte {
	b, err := ioutil.ReadFile("./gopher.jpeg")
	if err != nil {
		panic(err)
	}
	return b
}