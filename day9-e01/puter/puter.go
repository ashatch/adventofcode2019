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
	OutputInstruction   = 4
	JumpIfTrue          = 5
	JumpIfFalse         = 6
	LessThan            = 7
	Equals              = 8
	RelativeBaseAdjust  = 9
	HaltInstruction     = 99
)

func writeValue(programArray []int, relativeBase int, programCounter int, argumentIndex int, value int) {
	instruction := programArray[programCounter]

	thousands := int(math.Floor(float64(instruction / 1000 % 1000)))
	hundreds := int(math.Floor(float64((instruction - thousands*1000) / 100 % 100)))

	relativeMode := hundreds == 2

	if relativeMode {
		resultIndex := programArray[programCounter+argumentIndex+1]
		offset := relativeBase + resultIndex
		programArray[offset] = value
	} else {
		resultIndex := programArray[programCounter+argumentIndex+1]
		programArray[resultIndex] = value
	}
}

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
		arg := programArray[programCounter+operandNumber+1]
		offset := relativeBase + arg
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
MyPuter does the thing
*/
func MyPuter(inputStrategy InputStrategy, outputStrategy OutputStrategy, program string) []int {
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
				writeValue(programArray, relativeBase, i, 2, value)
				i += 3
			}
		case MultiplyInstruction:
			{
				lhs := operand(programArray, relativeBase, i, 0)
				rhs := operand(programArray, relativeBase, i, 1)
				value := lhs * rhs
				writeValue(programArray, relativeBase, i, 2, value)
				i += 3
			}
		case InputInstruction:
			{
				value := inputStrategy.GetInput()
				writeValue(programArray, relativeBase, i, 0, value)
				i++
			}
		case OutputInstruction:
			{
				operand := operand(programArray, relativeBase, i, 0)
				outputStrategy.SendOutput(operand)
				i++
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

				if a < b {
					writeValue(programArray, relativeBase, i, 2, 1)
				} else {
					writeValue(programArray, relativeBase, i, 2, 0)
				}
				i += 3
			}
		case Equals:
			{
				a := operand(programArray, relativeBase, i, 0)
				b := operand(programArray, relativeBase, i, 1)

				if a == b {
					writeValue(programArray, relativeBase, i, 2, 1)
				} else {
					writeValue(programArray, relativeBase, i, 2, 0)
				}
				i += 3
			}
		case RelativeBaseAdjust:

			{
				parameter := operand(programArray, relativeBase, i, 0)
				relativeBase += parameter
				i++
			}
		case HaltInstruction:
			{
				inputStrategy.Close()
				return programArray
			}
		default:
			{
				fmt.Println("fault", instruction)
				return programArray
			}
		}
		// fmt.Println(programArray)
	}
	return programArray
}
