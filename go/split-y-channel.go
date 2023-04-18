package main

import (
	"image"
	"log"
	// "io"
	"os"
	"image/png"
	"image/color"
	"image/jpeg"
	"fmt"
)


func main() {
	CopyYChannel()
}

func CopyYChannel() {
	defer fmt.Println("I'm outta here!")

	reader, err := os.Open("media/chairs-smol.jpg")
	if err != nil {
	    log.Fatal(err)
	}
	writer, err := os.Create("media/chairs-block-channel.png")
	if err != nil {
		log.Fatal(err)
	}

	m, err := jpeg.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}


	defer reader.Close()	
	defer writer.Close()

	img := image.NewYCbCr(m.Bounds(), image.YCbCrSubsampleRatio411)
	bounds := m.Bounds()
	fmt.Printf("Image is a %T\n", m)

	for x := bounds.Min.X; x < bounds.Max.X; x+=8 {
		for y := bounds.Min.Y; y < bounds.Max.Y; y+= 8 {
			// xOff, yOff := x / 8, y / 8
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					var r,g,b, _ uint32 = m.At(x+i,y+j).RGBA()
					var yy, _, _ uint8 = color.RGBToYCbCr(uint8(r),uint8(g),uint8(b))
					img.Y[x+i*8+y+j] = yy
					
				}	
			}
		}		
	}

	if err := png.Encode(writer, img); err != nil {
		log.Fatal(err)
	}

}