package peexl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstLayer(t *testing.T) {
	zeroCount := FreqCountValueAtLayer("01012222", 0, 4, 0)
	assert.Equal(t, 2, zeroCount)
}

func TestSecondLayer(t *testing.T) {
	zeroCount := FreqCountValueAtLayer("01012222", 1, 4, 0)
	assert.Equal(t, 0, zeroCount)
}

func TestBoundary(t *testing.T) {
	assert.Equal(t, 0, FreqCountValueAtLayer("11110222", 0, 4, 0))
	assert.Equal(t, 1, FreqCountValueAtLayer("11110222", 1, 4, 0))
}

func TestMixed(t *testing.T) {
	assert.Equal(t, 4, FreqCountValueAtLayer("1111222233334444", 3, 4, 4))
}
