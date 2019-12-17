package main

import (
	"bufio"
	"log"
	"os"

	puter "./puter"
)

func main() {
	file, err := os.Open("example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	program := scanner.Text()

	input := []int{9, 8, 7, 6, 5}

	puter.AmpSequence(program, input)
	// out := puter.FindMaxAmpSequence(program)
	// fmt.Println(out)
}
