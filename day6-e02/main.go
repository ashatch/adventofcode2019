package main

import (
	"bufio"
	"fmt"
	"os"

	"./orbit"
)

func main() {
	system := orbit.NewSystem("COM")

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		orbit.AddDeclaration(orbit.ParseDeclaration(line), system)
	}

	count := orbit.TransferCount(system, "YOU", "SAN")

	fmt.Println(count)
}
