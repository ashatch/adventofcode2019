package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var mass int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if i1, err := strconv.ParseFloat(line, 64); err == nil {
			x := math.Floor(i1 / 3.0)
			mass += (int(x) - 2)
		}
	}

	fmt.Println(mass)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
