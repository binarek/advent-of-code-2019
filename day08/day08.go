package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	imageWidth  = 25
	imageHeight = 6
	layerSize   = imageWidth * imageHeight
)

func main() {
	encodedImage, _ := readEncodedImage()
	fmt.Println(SolvePart1(encodedImage))
}

func readEncodedImage() ([]int8, error) {
	file, _ := os.Open(func() string {
		if len(os.Args) >= 2 {
			return os.Args[1]
		}
		return "./day08_input"
	}())
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	image := make([]int8, len(line))
	for i, char := range line {
		digit, _ := strconv.Atoi(string(char))
		image[i] = int8(digit)
	}
	if len(image)%layerSize != 0 {
		return image, errors.New("Invalid image data")
	}
	return image, nil
}

func SolvePart1(encodedImage []int8) int {
	imageLayers := splitLayers(encodedImage)

	var foundLayer []int8
	minZeroCount := math.MaxInt64
	for _, layer := range imageLayers {
		zeroCount := 0
		for _, digit := range layer {
			if digit == 0 {
				zeroCount++
			}
		}
		if zeroCount < minZeroCount {
			foundLayer = layer
			minZeroCount = zeroCount
		}
	}

	oneCount, twoCount := 0, 0
	for _, digit := range foundLayer {
		if digit == 1 {
			oneCount++
		} else if digit == 2 {
			twoCount++
		}
	}
	return oneCount * twoCount
}

func splitLayers(encodedImage []int8) [][]int8 {
	imageLen := len(encodedImage)
	layers := make([][]int8, 0, imageLen/layerSize)

	start, end := 0, layerSize
	for end <= imageLen {
		layers = append(layers, encodedImage[start:end])
		start = end
		end += layerSize
	}
	return layers
}
