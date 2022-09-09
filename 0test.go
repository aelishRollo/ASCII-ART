package main

import (
    "fmt"
    "image"
    "image/png"
    "os"
    "io"
    "net/http"
)

func main() {
    // You can register another format here
    image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

    file, err := os.Open("/home/aelish/repos/ascii-art/photos/noface.png")

    if err != nil {
        fmt.Println("Error: File could not be opened")
        os.Exit(1)
    }

    defer file.Close()

    pixels, err := getPixels(file)

    if err != nil {
        fmt.Println("Error: Image could not be decoded")
        os.Exit(1)
    }

    fmt.Println(pixels)
    fmt.Println(getBrightnessFromPixels(pixels))

    // now we want to get a new array which will recieve the pixel array as an argument,
    // and return an array of brightness values for each pixel
    // the logic is there, but the issue is a mismatch in type between the custom type Pixel, and
    // []int. So we need to get each value from the Pixel and assign that to be a new value in the 
    // array
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
    img, _, err := image.Decode(file)

    if err != nil {
        return nil, err
    }

    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    var pixels [][]Pixel
    for y := 0; y < height; y++ {
        var row []Pixel
        for x := 0; x < width; x++ {
            row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
        }
        pixels = append(pixels, row)
    }

    return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
    return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}


func getBrightnessFromPixels(inputArray [][]Pixel) []int {	//accepts a two dimensional array, and returns an integer array
	result := []int{}
	tempVar := 0
	for m := 0; m < len(inputArray); m++ {
		tempVar = 0
		for n := 0; n < len(inputArray[m]); n++ {
			tempVar += inputArray[m][n]
		}
		result = append(result, tempVar)
	}
	return result
}




// Pixel struct example
type Pixel struct {
    R int
    G int
    B int
    A int
}
