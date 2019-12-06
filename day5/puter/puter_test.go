package puter

import (
	"testing"
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
		programArray := MyPuter(nil, inputs[i])
		if programArray[0] != expectedOutputAtIndexZero[i] {
			t.Fail()
		}
	}
}

func TestInputOpcode(t *testing.T) {
	inputData := []string{
		"42",
	}

	input := NewSuppliedInput(inputData)
	programArray := MyPuter(input, "3,3,99,0")
	if programArray[3] != 42 {
		t.Fail()
	}
}
