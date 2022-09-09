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

    //fmt.Println(pixels)


    fmt.Println(getBrightnessArray(pixels))
    fmt.Println(mapBrightnessArrayToASCII(getBrightnessArray(pixels)))
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

func getBrightnessArray(pixels [][]Pixel) [][]int {

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

   
    
        return brightnessArray
}


func mapBrightnessArrayToASCII(brightnessArray [][]int) [][]string {

    ASCIIArray := [][]string {}
    tempArray := []string{}
    currentChar := ""
    for i := 0; i < len(brightnessArray); i++ {
            tempArray = nil
        for j := 0; j < len(brightnessArray[0]); j++ { // each temp array is a row/column 
                
                  
                 currentChar = ascify(brightnessArray[i][j])
                 tempArray = append(tempArray, currentChar)
                
        }
        ASCIIArray = append(ASCIIArray,tempArray)     
    }

   
    
        return ASCIIArray
}

func ascify(num int) string {
	result := ""
	stringOfChars := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	result = string(stringOfChars[num*64/255])
	return result
}





// Pixel struct example
type Pixel struct {
    R int
    G int
    B int
    A int
}
