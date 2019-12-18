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
	JumpIfTrue          = 5
	JumpIfFalse         = 6
	LessThan            = 7
	Equals              = 8
	RelativeBaseAdjust  = 9
	HaltInstruction     = 99
)

func operand(programArray []int, relativeBase int, programCounter int, operandNumber int) int {
	instruction := programArray[programCounter]

	thousands := int(math.Floor(float64(instruction / 1000 % 1000)))
	hundreds := int(math.Floor(float64((instruction - thousands*1000) / 100 % 100)))

	immediateMode := ((operandNumber == 0) && (hundreds == 1)) || ((operandNumber == 1) && (thousands == 1))

	if immediateMode {
		return programArray[programCounter+operandNumber+1]
	}

	relativeMode := ((operandNumber == 0) && (hundreds == 2)) || ((operandNumber == 1) && (thousands == 2))

	if relativeMode {
		offset := relativeBase + programArray[programCounter+operandNumber+1]
		return programArray[offset]
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

	for i := 0; i < 10000; i++ {
		programArray = append(programArray, 0)
	}

	return programArray
}

/*
does the thing
*/
func MyPuter(inputStrategy InputStrategy, program string) []int {
	var programArray = parseProgram(program)
	var relativeBase int

	for i := 0; i < len(programArray); i++ {
		baseInstruction := programArray[i]
		thousands := int(math.Floor(float64(baseInstruction / 1000 % 1000)))
		hundreds := int(math.Floor(float64((baseInstruction - thousands*1000) / 100 % 100)))
		instruction := baseInstruction - (thousands * 1000) - (hundreds * 100)

		switch instruction {
		case AddInstruction:
			{
				lhs := operand(programArray, relativeBase, i, 0)
				rhs := operand(programArray, relativeBase, i, 1)
				value := lhs + rhs
				resultIndex := programArray[i+3]
				programArray[resultIndex] = value
				i += 3
			}
		case MultiplyInstruction:
			{
				lhs := operand(programArray, relativeBase, i, 0)
				rhs := operand(programArray, relativeBase, i, 1)
				value := lhs * rhs
				resultIndex := programArray[i+3]
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
				operand := operand(programArray, relativeBase, i, 0)
				fmt.Println(operand)
				i++
			}
		case HaltInstruction:
			{
				return programArray
			}
		case JumpIfTrue:
			{
				conditional := operand(programArray, relativeBase, i, 0)
				location := operand(programArray, relativeBase, i, 1)
				if conditional != 0 {
					i = location - 1 // -1 as we increment in the loop
				} else {
					i += 2
				}
			}
		case JumpIfFalse:
			{
				conditional := operand(programArray, relativeBase, i, 0)
				location := operand(programArray, relativeBase, i, 1)
				if conditional == 0 {
					i = location - 1 // -1 as we increment in the loop
				} else {
					i += 2
				}
			}
		case LessThan:
			{
				a := operand(programArray, relativeBase, i, 0)
				b := operand(programArray, relativeBase, i, 1)
				location := programArray[i+3]

				if a < b {
					programArray[location] = 1
				} else {
					programArray[location] = 0
				}
				i += 3
			}
		case Equals:
			{
				a := operand(programArray, relativeBase, i, 0)
				b := operand(programArray, relativeBase, i, 1)
				location := programArray[i+3]
				if a == b {
					programArray[location] = 1
				} else {
					programArray[location] = 0
				}
				i += 3
			}
		case RelativeBaseAdjust:
			{
				parameter := operand(programArray, relativeBase, i, 0)
				relativeBase = parameter
				i++
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
