package main

import (
	"image"
	"log"
	"os"
	"image/jpeg"
	"image/color"
)


func main() {	
	var in_file_name string =  "../media/chairs-smol.jpg"
	var out_file_name string =  "../media/chairs-smol-demo.jpg"
	reader, err := os.Open(in_file_name)
	if err != nil {
	    log.Fatal(err)
	}
	
	writer, err := os.Create(out_file_name)
	if err != nil {  log.Fatal(err) }

	defer writer.Close()
	defer reader.Close()	

	m, err := jpeg.Decode(reader)
	if err != nil { log.Fatal(err) }

 	var bounds image.Rectangle = m.Bounds()
	var img *image.YCbCr = image.NewYCbCr(bounds, image.YCbCrSubsampleRatio444)
	for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
		for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
			var r, g, b, _ uint32 = m.At(x,y).RGBA()
			yy, cb, cr := color.RGBToYCbCr(uint8(r),uint8(g),uint8(b))
			img.Y[img.YOffset(x,y)] = yy
			img.Cb[img.COffset(x,y)] = cb
			img.Cr[img.COffset(x,y)] = cr			
		}
	}
	options := jpeg.Options{Quality:100}
	err = jpeg.Encode(writer, img, &options)
	if err != nil { log.Fatal(err) }
	
}
