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

	stats := orbit.Count(system)
	total := stats.DirectOrbitCount + stats.IndirectOrbitCount

	fmt.Println(total)
}
