package main

import (
	"bufio"
	"log"
	"os"

	"./peexl"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	image := scanner.Text()

	imageDefinition := &peexl.ImageDefinition{
		StringData: image,
		Width:      25,
		Height:     6,
	}

	peexl.DecodeImageString(imageDefinition)
}
