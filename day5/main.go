package main

import (
	"bufio"
	"fmt"
	"os"

	"./puter"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	programString := scanner.Text()

	puter.MyPuter(programString)

	if scanner.Err() != nil {
		fmt.Println("faily")
	}
}
