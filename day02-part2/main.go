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
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if puter.MyPuter(programString, noun, verb) == 19690720 {
				fmt.Println((noun * 100) + verb)
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Println("faily")
	}
}
