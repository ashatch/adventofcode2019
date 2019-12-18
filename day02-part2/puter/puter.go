package puter

import (
	"fmt"
	"strconv"
	"strings"
)

/*
does the thing
*/
func MyPuter(program string, noun int, verb int) int {
	programArrayStrings := strings.Split(program, ",")
	var programArray = []int{}

	for _, i := range programArrayStrings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		programArray = append(programArray, j)
	}

	programArray[1] = noun
	programArray[2] = verb

	// fmt.Println(programArray)

	for i := 0; i < len(programArray); i++ {
		token := programArray[i]

		if token == 1 {
			// fmt.Println("adding")
			programArray[programArray[i+3]] = programArray[programArray[i+1]] + programArray[programArray[i+2]]
			// fmt.Println(programArray)
			i += 3
		} else if token == 2 {
			// fmt.Println("multiplying")
			programArray[programArray[i+3]] = programArray[programArray[i+1]] * programArray[programArray[i+2]]
			// fmt.Println(programArray)
			i += 3
		} else if token == 99 {
			// fmt.Println(programArray)
			return programArray[0]
		} else {
			fmt.Println("fault")
			return -1
		}
	}
	return -1

}
