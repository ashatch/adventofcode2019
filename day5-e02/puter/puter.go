package puter

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	AddInstruction      = 1
	MultiplyInstruction = 2
	InputInstruction    = 3
	PrintInstruction    = 4
	HaltInstruction     = 99
)

func operand(programArray []int, programCounter int, operandNumber int) int {
	instruction := programArray[programCounter]

	thousands := int(math.Floor(float64(instruction / 1000 % 1000)))
	hundreds := int(math.Floor(float64((instruction - thousands*1000) / 100 % 100)))

	// fmt.Println("    (operand pc:", programCounter, "t", thousands, "h", hundreds, "on:", operandNumber, ")")
	// fmt.Println("         value: ", programArray[programCounter+operandNumber+1])

	immediateMode := ((operandNumber == 0) && (hundreds == 1)) || ((operandNumber == 1) && (thousands == 1))

	if immediateMode {
		return programArray[programCounter+operandNumber+1]
	}

	offset := programArray[programCounter+operandNumber+1]
	return programArray[offset]
}

func parseProgram(program string) []int {
	programArrayStrings := strings.Split(program, ",")
	var programArray = []int{}

	for _, i := range programArrayStrings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		programArray = append(programArray, j)
	}

	return programArray
}

/*
does the thing
*/
func MyPuter(inputStrategy InputStrategy, program string) []int {
	var programArray = parseProgram(program)

	for i := 0; i < len(programArray); i++ {
		baseInstruction := programArray[i]
		thousands := int(math.Floor(float64(baseInstruction / 1000 % 1000)))
		hundreds := int(math.Floor(float64((baseInstruction - thousands*1000) / 100 % 100)))
		instruction := baseInstruction - (thousands * 1000) - (hundreds * 100)

		// fmt.Println("base", baseInstruction, "instruction: ", instruction)

		switch instruction {
		case AddInstruction:
			{
				lhs := operand(programArray, i, 0)
				rhs := operand(programArray, i, 1)
				value := lhs + rhs
				resultIndex := programArray[i+3]
				// fmt.Println("+", lhs, rhs, "=", value, "=>", resultIndex)
				programArray[resultIndex] = value
				i += 3
			}
		case MultiplyInstruction:
			{
				lhs := operand(programArray, i, 0)
				rhs := operand(programArray, i, 1)
				value := lhs * rhs
				resultIndex := programArray[i+3]
				// fmt.Println("*", lhs, rhs, "=", value, "=>", resultIndex)
				programArray[resultIndex] = value
				i += 3
			}
		case InputInstruction:
			{
				input, err := strconv.Atoi(inputStrategy.GetInput())
				if err == nil {
					argument := programArray[i+1]
					programArray[argument] = input
				} else {
					fmt.Println("fault - supplied non-integer data")
				}
				i++
			}
		case PrintInstruction:
			{
				operand := operand(programArray, i, 0)
				fmt.Println(operand)
				i++
			}
		case HaltInstruction:
			{
				return programArray
			}
		default:
			{
				fmt.Println("fault", instruction)
				return programArray
			}
		}
	}
	return programArray
}
