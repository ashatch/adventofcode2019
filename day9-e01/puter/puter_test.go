package puter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	inputData := []int{}

	input := NewSuppliedInput(inputData)
	programArray := MyPuter(input, nil, "1002,4,3,4,33")
	if programArray[4] != 99 {
		t.Fail()
	}
}

func TestParameterModeNegative(t *testing.T) {
	inputData := []int{}

	input := NewSuppliedInput(inputData)
	programArray := MyPuter(input, nil, "1101,100,-1,4,0")
	if programArray[4] != 99 {
		t.Fail()
	}
}

func TestParameterModePrint(t *testing.T) {
	inputData := []int{}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	MyPuter(input, output, "4,2,4,3,99")
}

func TestParameterModePrintImmediate(t *testing.T) {
	inputData := []int{}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	MyPuter(input, output, "104,2,4,3,99")
	fmt.Println(output.Output)
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

func TestLargeNumberOutput(t *testing.T) {
	inputData := []int{
		110,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	program := "104,1125899906842624,99"
	MyPuter(input, output, program)
	assert.Equal(t, 1125899906842624, output.Output[0])
}

/*
For example, if the relative base is 2000, then after the instruction 109,19, the relative base would be 2019. If the next instruction were 204,-34, then the value at address 1985 would be output.
*/

func TestRelativeModeProgram(t *testing.T) {
	inputData := []int{}

	input := NewSuppliedInput(inputData)

	programs := []string{
		"2101,3,1,7,4,7,99,0",
		"2102,3,1,7,4,7,99,0",

		"109,2,2101,3,1,9,4,9,99,0",
		"109,2,2102,3,1,9,4,9,99,0",

		"109,2,1201,2,3,9,4,9,99,0",
		"109,2,1202,2,3,9,4,9,99,0",

		"109,2,1201,2,3,9,204,7,99,0",
		"109,2,1202,2,3,9,204,7,99,0",
	}

	outputs := []int{
		6,
		9,
		6,
		9,
		6,
		9,
		6,
		9,
	}

	for i := 0; i < len(programs); i++ {
		output := NewStoredOutput()
		MyPuter(input, output, programs[i])
		assert.Equal(t, outputs[i], output.Output[0])
	}
}

func TestRelativeModeInput(t *testing.T) {
	inputData := []int{
		42,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()

	program := "109,6,203,1,4,7,99,0"

	MyPuter(input, output, program)

	assert.Equal(t, 42, output.Output[0])
}

func TestRelativeAndAbsoluteWithAddition(t *testing.T) {
	inputData := []int{}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	program := "109,8,21102,100,100,1,4,9,99,0"
	MyPuter(input, output, program)

	assert.Equal(t, 10000, output.Output[0])
}

func TestRelativeModeQuine(t *testing.T) {
	inputData := []int{
		110,
	}

	input := NewSuppliedInput(inputData)
	output := NewStoredOutput()
	program := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	MyPuter(input, output, program)
	programAsInts := []int{
		109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
	}
	assert.Equal(t, programAsInts, output.Output)
}
