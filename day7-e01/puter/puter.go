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
	HaltInstruction     = 99
)

func operand(programArray []int, programCounter int, operandNumber int) int {
	instruction := programArray[programCounter]

	thousands := int(math.Floor(float64(instruction / 1000 % 1000)))
	hundreds := int(math.Floor(float64((instruction - thousands*1000) / 100 % 100)))

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
MyPuter does the thing
*/
func MyPuter(inputStrategy InputStrategy, outputStrategy OutputStrategy, program string) []int {
	var programArray = parseProgram(program)

	for i := 0; i < len(programArray); i++ {
		baseInstruction := programArray[i]
		thousands := int(math.Floor(float64(baseInstruction / 1000 % 1000)))
		hundreds := int(math.Floor(float64((baseInstruction - thousands*1000) / 100 % 100)))
		instruction := baseInstruction - (thousands * 1000) - (hundreds * 100)

		switch instruction {
		case AddInstruction:
			{
				lhs := operand(programArray, i, 0)
				rhs := operand(programArray, i, 1)
				value := lhs + rhs
				resultIndex := programArray[i+3]
				programArray[resultIndex] = value
				i += 3
			}
		case MultiplyInstruction:
			{
				lhs := operand(programArray, i, 0)
				rhs := operand(programArray, i, 1)
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
		case OutputInstruction:
			{
				operand := operand(programArray, i, 0)				
				outputStrategy.SendOutput(operand)
				i++
			}
		case HaltInstruction:
			{
				return programArray
			}
		case JumpIfTrue:
			{
				conditional := operand(programArray, i, 0)
				location := operand(programArray, i, 1)
				if conditional != 0 {
					i = location - 1 // -1 as we increment in the loop
				} else {
					i += 2
				}
			}
		case JumpIfFalse:
			{
				conditional := operand(programArray, i, 0)
				location := operand(programArray, i, 1)
				if conditional == 0 {
					i = location - 1 // -1 as we increment in the loop
				} else {
					i += 2
				}
			}
		case LessThan:
			{
				a := operand(programArray, i, 0)
				b := operand(programArray, i, 1)
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
				a := operand(programArray, i, 0)
				b := operand(programArray, i, 1)
				location := programArray[i+3]
				if a == b {
					programArray[location] = 1
				} else {
					programArray[location] = 0
				}
				i += 3
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
