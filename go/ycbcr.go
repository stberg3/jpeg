package main

import (
	"image"
	"log"
	"os"
	"image/png"
	"image/jpeg"
	"fmt"
	// "io"
	"image/color"
)


func main() {
	

	reader, err := os.Open("../media/swatch-ycbcr-8x8.jpg")
	if err != nil {
	    log.Fatal(err)
	}
	
	defer reader.Close()	

	// func Decode(r io.Reader) (image.Image, error)
	m, err := jpeg.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	processJPG(m)

}

func processPNG(m image.Image) {

	// img = image.NewYCbCr(m.Bounds(), image.YCbCrSubsampleRatio444)
	// fmt.Printf("The original image is a %T\n", m)
	// fmt.Printf("The new image is a %T\n", img)

	// fmt.Printf("Y: %d, Cb: %d, Cr: %d\n", len(img.Y), len(img.Cb), len(img.Cr))
	// fmt.Printf("YStride: %d\n", img.YStride)
	// fmt.Printf("CStride: %d\n", img.CStride)
	// fmt.Printf("SubsampleRatio: %v\n", img.SubsampleRatio)
	// fmt.Printf("Rect: %v\n",           img.Rect)


	b := m.Bounds()
	for x := b.Min.X; x < b.Max.X; x += 8 {
		for y := b.Min.Y; y < b.Max.Y; y += 8 {		
			fmt.Printf("(%d,%d): %v\n", x, y, m.At(x,y))
		}
	}

	return 
}

func processJPG(m image.Image) {
	writer, err := os.Create("../media/swatch-copy.png")
	if err != nil {
		log.Fatal(err)
	}

	defer writer.Close()
	var img *image.YCbCr = image.NewYCbCr(m.Bounds(), image.YCbCrSubsampleRatio444)
	b := m.Bounds()

	fmt.Printf("Y Len: %d\n", len(img.Y))
	fmt.Printf("width: %d\n", img.YStride)

	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {		
			var r,g,b,_ uint32 = m.At(x,y).RGBA()
			yy, cb, cr := color.RGBToYCbCr(uint8(r),uint8(g),uint8(b))
			// yy, _, _ := color.RGBToYCbCr(uint8(r),uint8(g),uint8(b))
			img.Y[x+y*img.YStride] = yy
			img.Cb[x+img.CStride*y] = cb
			img.Cr[x+img.CStride*y] = cr
		}
	}

	// options := jpeg.Options{Quality:100}

	// err = jpeg.Encode(writer, img, *options)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err = png.Encode(writer, img)
	if err != nil {
		log.Fatal(err)
	}

}