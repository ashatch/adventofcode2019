package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"./wirelib"
)

type wire struct {
	x         int
	y         int
	length    int
	steps     int
	direction int
}

type intersection struct {
	x     int
	y     int
	steps int
}

func main() {
	file, err := os.Open("steps.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	wire1 := scanner.Text()
	scanner.Scan()
	wire2 := scanner.Text()

	route1 := wirelib.Parse(wire1)
	route2 := wirelib.Parse(wire2)

	intersections := wirelib.FindIntersections(route1, route2)

	for _, i := range intersections {
		dist := wirelib.DistanceToIntersection(i)
		fmt.Println(dist)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
