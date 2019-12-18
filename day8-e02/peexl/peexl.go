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

type LayerInfo struct {
	LayerIndex int
	ZeroCount  int
}

func FreqCountValueAtLayer(imageDefinition *ImageDefinition, layerIndex int, layerPixelCount int, valueToFind int) int {
	var count int = 0

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
	pixelCount := len(imageDefinition.StringData)
	layerPixelCount := imageDefinition.Width * imageDefinition.Height
	layerCount := pixelCount / layerPixelCount
	fmt.Println("pixelCount:", pixelCount)
	fmt.Println("layerPixelCount:", layerPixelCount)
	fmt.Println("layerCount", layerCount)

	layerInfo := []LayerInfo{}

	for layerIndex := 0; layerIndex < layerCount; layerIndex++ {

		zeroCount := FreqCountValueAtLayer(imageDefinition, layerIndex, layerPixelCount, 0)

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

	oneDigits := FreqCountValueAtLayer(imageDefinition, minZeroLayer.LayerIndex, layerPixelCount, 1)
	twoDigits := FreqCountValueAtLayer(imageDefinition, minZeroLayer.LayerIndex, layerPixelCount, 2)

	fmt.Println(oneDigits * twoDigits)
}
