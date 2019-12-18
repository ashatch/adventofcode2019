package main

import (
	"bufio"
	"log"
	"os"

	puter "./puter"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stdinInput := puter.NewStdinInput()
	puter.MyPuter(stdinInput, scanner.Text())
}
