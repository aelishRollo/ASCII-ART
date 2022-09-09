package main

import (
    "fmt"
    "image"
    "image/png"
    "os"
    "io"
    //"net/http"
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

/*    fmt.Println(pixels)
    fmt.Println("Datura")
    fmt.Println("Datura")
    fmt.Println("Brugmansia")

    fmt.Println("We gonna print length of pixels now")

    fmt.Println(len(pixels))

    fmt.Println("We gonna print length of pixels[0] now baby!@!!!!")

    fmt.Println(len(pixels[0]))

    fmt.Println("Datura")
    fmt.Println("Datura")
    fmt.Println("Brugmansia")

    fmt.Println("We gonna print pixels now")
*/


    fmt.Println(getBrightnessArray(pixels))
    // now we want to get a new array which will recieve the pixel array as an argument,
    // and return an array of brightness values for each pixel
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

func getBrightnessArray([][]Pixel) [][]int {

    brightnessArray := [][]int {}
    tempArray := []int{}
    brightnessNumber := 0
    for i := 0; i < len(pixels); i++ {
            tempArray = nil
        for j := 0; j < len(pixels[0]); j++ { // each temp array is a row/column 
                
                 brightnessNumber = 0
                 brightnessNumber += pixels[i][j].R
                 brightnessNumber += pixels[i][j].G
                 brightnessNumber += pixels[i][j].B
                 brightnessNumber = brightnessNumber/3
                 tempArray = append(tempArray, brightnessNumber)
                
        }
        brightnessArray = append(brightnessArray,tempArray)     
    }

    fmt.Println("YOOOO")
    fmt.Println("YOOOO")
    fmt.Println("YOOOO")
    fmt.Println("YOOOO")
    fmt.Println("This is the length of brightnessArray")
    fmt.Println(len(brightnessArray))
    fmt.Println("This is length of brightnessArray[0]")
    fmt.Println(len(brightnessArray[0])) 
    fmt.Println("the array itself")
    fmt.Println(brightnessArray)
    
        return brightnessArray
}



// Pixel struct example
type Pixel struct {
    R int
    G int
    B int
    A int
}
