package peexl

import (
	"fmt"
	"strconv"
)

type ImageDefinition struct {
	StringData string
	Width      int
	Height     int
}

type ImageLayersInfo struct {
	PixelCount      int
	LayerPixelCount int
	LayerCount      int
}

type LayerInfo struct {
	LayerIndex int
	ZeroCount  int
}

const (
	Black       = 0
	White       = 1
	Transparent = 2
)

func Offset(image *ImageDefinition, layer int, x int, y int) int {
	layerInfo := ExtractImageLayerInfo(image)
	return (layerInfo.LayerPixelCount * layer) + (y * image.Width) + x
}

func PixelValue(image *ImageDefinition, offset int) int {
	pixelChar := string(image.StringData[offset])
	pixelValue, _ := strconv.Atoi(pixelChar)
	return pixelValue
}

func RenderImage(image *ImageDefinition) {
	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			value := RenderPixel(image, x, y)
			if value == Black {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func RenderPixel(image *ImageDefinition, x int, y int) int {
	layerInfo := ExtractImageLayerInfo(image)

	for layer := 0; layer < layerInfo.LayerCount; layer++ {
		offset := Offset(image, layer, x, y)
		pixel := PixelValue(image, offset)
		switch pixel {
		case White:
			{
				return White
			}
		case Black:
			{
				return Black
			}
		default:
			{
				// continue - it's transparent
			}
		}
	}
	return Transparent
}

func ExtractImageLayerInfo(image *ImageDefinition) *ImageLayersInfo {
	pixelCount := len(image.StringData)
	layerPixelCount := image.Width * image.Height
	layerCount := pixelCount / layerPixelCount

	return &ImageLayersInfo{
		PixelCount:      pixelCount,
		LayerPixelCount: layerPixelCount,
		LayerCount:      layerCount,
	}
}

func FreqCountValueAtLayer(imageDefinition *ImageDefinition, layerIndex int, valueToFind int) int {
	var count int = 0

	layerPixelCount := imageDefinition.Width * imageDefinition.Height
	startPixelIndex := layerIndex * layerPixelCount
	endPixelIndex := startPixelIndex + layerPixelCount

	for pixelIndex := startPixelIndex; pixelIndex < endPixelIndex; pixelIndex++ {
		pixelChar := string(imageDefinition.StringData[pixelIndex])
		pixelValue, _ := strconv.Atoi(pixelChar)
		if pixelValue == valueToFind {
			count++
		}
	}

	return count
}

func DecodeImageString(imageDefinition *ImageDefinition) {
	imageLayerInfo := ExtractImageLayerInfo(imageDefinition)
	fmt.Println("pixelCount:", imageLayerInfo.PixelCount)
	fmt.Println("layerPixelCount:", imageLayerInfo.LayerPixelCount)
	fmt.Println("layerCount", imageLayerInfo.LayerCount)

	layerInfo := []LayerInfo{}

	for layerIndex := 0; layerIndex < imageLayerInfo.LayerCount; layerIndex++ {
		zeroCount := FreqCountValueAtLayer(imageDefinition, layerIndex, 0)

		item := LayerInfo{
			LayerIndex: layerIndex,
			ZeroCount:  zeroCount,
		}
		layerInfo = append(layerInfo, item)
	}

	var minZeroLayer = layerInfo[0]
	for _, layer := range layerInfo {
		if layer.ZeroCount < minZeroLayer.ZeroCount {
			minZeroLayer = layer
		}
	}

	oneDigits := FreqCountValueAtLayer(imageDefinition, minZeroLayer.LayerIndex, 1)
	twoDigits := FreqCountValueAtLayer(imageDefinition, minZeroLayer.LayerIndex, 2)

	fmt.Println(oneDigits * twoDigits)
}
