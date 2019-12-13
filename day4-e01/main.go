package main

import (
	"fmt"

	"./plib"
)

func main() {
	count := plib.CheckPasswordRange(146810, 612564)
	fmt.Println(count)
}
