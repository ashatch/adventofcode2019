package puter

import (
	"fmt"
	"strconv"
	"strings"
)

/*
does the thing
*/
func MyPuter(inputStrategy InputStrategy, program string) []int {
	programArrayStrings := strings.Split(program, ",")
	var programArray = []int{}

	for _, i := range programArrayStrings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		programArray = append(programArray, j)
	}

	for i := 0; i < len(programArray); i++ {
		token := programArray[i]

		if token == 1 {
			programArray[programArray[i+3]] = programArray[programArray[i+1]] + programArray[programArray[i+2]]
			i += 3
		} else if token == 2 {
			programArray[programArray[i+3]] = programArray[programArray[i+1]] * programArray[programArray[i+2]]
			i += 3
		} else if token == 3 {
			input, err := strconv.Atoi(inputStrategy.GetInput())
			if err == nil {
				argument := programArray[i+1]
				programArray[argument] = input
			} else {
				fmt.Println("fault - supplied non-integer data")
			}

			i++
		} else if token == 4 {
			argument := programArray[i+1]
			fmt.Println(programArray[argument])
			i++
		} else if token == 99 {
			return programArray
		} else {
			fmt.Println("fault")
			return programArray
		}
	}
	return programArray

}
