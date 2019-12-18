package puter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhaseSequenceExample(t *testing.T) {
	// 4, 3, 2, 1, 0
	program := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	input := []int{4, 3, 2, 1, 0}
	result := AmpSequence(program, input)
	assert.Equal(t, 43210, result)
}

func TestFindMaxAmpSequence(t *testing.T) {
	// 4, 3, 2, 1, 0
	program := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	result := FindMaxAmpSequence(program)
	assert.Equal(t, 65210, result)
}
