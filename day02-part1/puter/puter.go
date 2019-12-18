package puter

import (
	"fmt"
	"strconv"
	"strings"
)

/*
MyPuter does the thing
*/
func MyPuter(program string, modify bool, noun int, verb int) []int {
	programArrayStrings := strings.Split(program, ",")
	var programArray = []int{}

	for _, i := range programArrayStrings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		programArray = append(programArray, j)
	}

	if modify {
		programArray[1] = noun
		programArray[2] = verb
	}

	for i := 0; i < len(programArray); i++ {
		token := programArray[i]

		if token == 1 {
			programArray[programArray[i+3]] = programArray[programArray[i+1]] + programArray[programArray[i+2]]
			i += 3
		} else if token == 2 {
			programArray[programArray[i+3]] = programArray[programArray[i+1]] * programArray[programArray[i+2]]
			i += 3
		} else if token == 99 {
			return programArray
		} else {
			fmt.Println("fault")
			return programArray
		}
	}
	return programArray

}
