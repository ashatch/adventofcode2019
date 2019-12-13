package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"./lib"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	totalFuel := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineInteger, _ := strconv.Atoi(line)
		totalFuel += lib.FuelForModule(lineInteger)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalFuel)
}
