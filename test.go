package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {

	imgPath := "/home/aelish/repos/ascii-art/photos/noface.png"  //image path here

	fmt.Println("How to get image dimensions - image.DecodeConfig")
	fmt.Println()
	fmt.Println(imgPath)
	file, err := os.Open(imgPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Width:", image.Width, "Height:", image.Height)
}
