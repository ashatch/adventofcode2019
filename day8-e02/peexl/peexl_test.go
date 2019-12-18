package peexl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstLayer(t *testing.T) {
	def := &ImageDefinition{
		StringData: "01012222",
		Width:      2,
		Height:     2,
	}
	zeroCount := FreqCountValueAtLayer(def, 0, 0)
	assert.Equal(t, 2, zeroCount)
}

func TestSecondLayer(t *testing.T) {
	def := &ImageDefinition{
		StringData: "01012222",
		Width:      2,
		Height:     2,
	}
	zeroCount := FreqCountValueAtLayer(def, 1, 0)
	assert.Equal(t, 0, zeroCount)
}

func TestBoundary(t *testing.T) {
	def := &ImageDefinition{
		StringData: "11110222",
		Width:      2,
		Height:     2,
	}

	assert.Equal(t, 0, FreqCountValueAtLayer(def, 0, 0))
	assert.Equal(t, 1, FreqCountValueAtLayer(def, 1, 0))
}

func TestMixed(t *testing.T) {
	def := &ImageDefinition{
		StringData: "1111222233334444",
		Width:      2,
		Height:     2,
	}
	assert.Equal(t, 4, FreqCountValueAtLayer(def, 3, 4))
}

func TestLayerInfo(t *testing.T) {
	def := &ImageDefinition{
		StringData: "0222112222120000",
		Width:      2,
		Height:     2,
	}

	layerInfo := ExtractImageLayerInfo(def)

	assert.Equal(t, 4, layerInfo.LayerCount)
	assert.Equal(t, 4, layerInfo.LayerPixelCount)
	assert.Equal(t, 16, layerInfo.PixelCount)
}

func TestRenderPixel(t *testing.T) {
	def := &ImageDefinition{
		StringData: "0222112222120000",
		Width:      2,
		Height:     2,
	}

	assert.Equal(t, Black, RenderPixel(def, 0, 0))
	assert.Equal(t, White, RenderPixel(def, 1, 0))
	assert.Equal(t, White, RenderPixel(def, 0, 1))
	assert.Equal(t, Black, RenderPixel(def, 1, 1))
}

func TestOffset(t *testing.T) {
	def := &ImageDefinition{
		StringData: "012012012012012112112112112112212212212212212",
		Width:      3,
		Height:     5,
	}
	offset := Offset(def, 1, 1, 1)

	assert.Equal(t, 19, offset)
}
