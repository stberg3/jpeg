package main

import (
	// "encoding/base64"
	// "fmt"
	"image"
	"log"
	// "strings"
	"io"
	"os"
	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	"image/png"
	// "image/draw"
	"image/color"
	_ "image/jpeg"
)


func main() {
	// Decode the PNG data. If reading from file, create a reader with
	//
	// reader, err := os.Open("media/prism.png")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer reader.Close()
	// m, err := png.Decode(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// x := m.Bounds().Max.X/2;
	// y := m.Bounds().Max.Y/2;
	// fmt.Printf("Pixel [0,0]: %v \n", m.At(0,0));
	// fmt.Printf("Color [%d,%d]: %v \n", x, y, m.At(x,y));
	// var r,g,b,_ uint32 = m.At(x,y).RGBA()
	// var yy,cb,cr uint8 = color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))

	// fmt.Printf("Color [%d,%d,%d]\n", yy, cb, cr);
	// fmt.Printf("Bounds [0,0]: %v \n", m.Bounds());

	// WriteStuff("Hello! Can you hear me??")
	// CopyStuff()
	CopyYChannel()
}

func WriteStuff(stuff string) {
	writer, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer writer.Close()

	if _, err := io.WriteString(writer, stuff);  err != nil {
		log.Fatal(err)
	}
}

func CopyStuff() {
	reader, err := os.Open("media/prism.png")
	if err != nil {
	    log.Fatal(err)
	}
	defer reader.Close()

	writer, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}

	defer writer.Close()

	if _, err := io.Copy(writer, reader); err != nil {
		log.Fatal(err)
	}

}

func CopyYChannel() {
	reader, err := os.Open("media/prism.png")
	if err != nil {
	    log.Fatal(err)
	}
	defer reader.Close()
	
	m, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	img := image.NewYCbCr(m.Bounds(), image.YCbCrSubsampleRatio411)
	bounds := m.Bounds()

	// for i := 0; i < img.CStride; i++ {
		for j := 0; j < len(img.Cb); j++ {
			var r,g,b,_ uint32 = m.At(i,j).RGBA()
			var _,cb,_ uint8 = color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))
			// img.Y[i*bounds.Max.X+j] = yy
			img.Cb[j] = cb
			// img.Cr[i*bounds.Max.X+j] = cr
		}		
	// }


	writer, err := os.Create("test-cb411-channel.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(writer, img); err != nil {
		log.Fatal(err)
	}

	defer writer.Close()

}