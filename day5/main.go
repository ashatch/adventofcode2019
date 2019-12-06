package main

import (
	puter "./puter"
)

func main() {
	stdinInput := puter.NewStdinInput()
	puter.MyPuter(stdinInput, "3,0,4,0,99")
}
