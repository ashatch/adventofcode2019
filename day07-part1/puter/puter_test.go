package puter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuterWithGivenExamples(t *testing.T) {
	inputs := []string{
		"1,9,10,3,2,3,11,0,99,30,40,50",
		"1,0,0,0,99",
		"1,1,1,4,99,5,6,0,99",
	}

	expectedOutputAtIndexZero := []int{
		3500,
		2,
		30,
	}

	for i := 0; i < len(inputs); i++ {
		programArray := MyPuter(nil, nil, inputs[i])
		if programArray[0] != expectedOutputAtIndexZero[i] {
			t.Fail()
		}
	}
}

func TestInputOpcode(t *testing.T) {
	inputData := []int{
		42,
	}

	input := NewSuppliedInput(inputData)
	programArray := MyPuter(input, nil, "3,3,99,0")
	if programArray[3] != 42 {
		t.Fail()
	}
}

func TestParameterMode(t *testing.T) {
	programArray := MyPuter(nil, nil, "1002,4,3,4,33")
	if programArray[4] != 99 {
		t.Fail()
	}
}

func TestParameterModeNegative(t *testing.T) {
	programArray := MyPuter(nil, nil, "1101,100,-1,4,0")
	if programArray[4] != 99 {
		t.Fail()
	}
}

func TestParameterModePrint(t *testing.T) {
	output := NewStoredOutput()
	MyPuter(nil, output, "4,2,4,3,99")
}

func TestParameterModePrintImmediate(t *testing.T) {
	output := NewStoredOutput()
	MyPuter(nil, output, "104,2,4,3,99")
}

func TestEqual(t *testing.T) {
	inputData := []int{
		42,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()

	MyPuter(input, output, "3,9,8,9,10,9,4,9,99,-1,8")

	assert.Equal(t, 0, output.Output[0])
}

func TestLessThan(t *testing.T) {
	inputData := []int{
		42,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()

	MyPuter(input, output, "3,9,7,9,10,9,4,9,99,-1,8")

	assert.Equal(t, 0, output.Output[0])
}
func TestInputPosition(t *testing.T) {
	inputData := []int{
		40,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	MyPuter(input, output, "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")

	assert.Equal(t, 1, output.Output[0])
}

func TestBigTestyTesty(t *testing.T) {
	inputData := []int{
		110,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	MyPuter(input, output, "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")

	assert.Equal(t, 1001, output.Output[0])
}
