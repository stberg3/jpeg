package main

import (
	"image"
	"fmt"
	"image/png"
	"os"
	"log"
)

func main() {
	reader, err := os.Open("media/prism.png")
	if err != nil {
	    log.Fatal(err)
	}
	defer reader.Close()	

	m, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The image is of type %T\n", m)

}


func readMetadata(m *image.Image) {

}
