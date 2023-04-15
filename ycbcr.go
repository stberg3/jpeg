package main

import (
	// "encoding/base64"
	"fmt"
	// "image"
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
	_ "image/jpeg"
)


func main() {
	// Decode the PNG data. If reading from file, create a reader with
	//
	reader, err := os.Open("media/prism.png")
	if err != nil {
	    log.Fatal(err)
	}
	defer reader.Close()
	m, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	x := m.Bounds().Max.X/2;
	y := m.Bounds().Max.Y/2;
	fmt.Printf("Pixel [0,0]: %v \n", m.At(0,0));
	fmt.Printf("Bounds [%d,%d]: %v \n", x, y, m.At(x,y));
	fmt.Printf("Bounds [0,0]: %v \n", m.Bounds());

	WriteStuff("Hello! Can you hear me??")
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