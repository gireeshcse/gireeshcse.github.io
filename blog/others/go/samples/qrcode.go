package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/ean"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

	//  The last digit being used as a checksum
	barCodeEan8, _ := ean.Encode("1234567")
	barCodeEan8New, _ := barcode.Scale(barCodeEan8, 200, 200)
	// create the output file
	fileNew, _ := os.Create("ean8.png")
	defer fileNew.Close()

	png.Encode(fileNew, barCodeEan8New)

	barCodeEan13, _ := ean.Encode("123456789012")
	barCodeEan13New, _ := barcode.Scale(barCodeEan13, 200, 200)
	// create the output file
	fileNew13, _ := os.Create("ean13.png")
	defer fileNew13.Close()

	png.Encode(fileNew13, barCodeEan13New)
}