package main

import (
	"fmt"

	"./plib"
)

func main() {
	count := plib.CheckPasswordRange(231832, 767346)
	fmt.Println(count)
}
