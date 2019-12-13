package main

import (
	"bufio"
	"fmt"
	"os"

	"./puter"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	programString := scanner.Text()
	result := puter.MyPuter(programString, true, 12, 2)

	if scanner.Err() != nil {
		fmt.Println("faily")
	}

	fmt.Println("result at index 0", result[0])
}
