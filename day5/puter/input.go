package puter

import "fmt"

/*
How to get user input
*/
type InputStrategy interface {
	GetInput()
}

type StdinInputStrategy struct {
}

func (s *StdinInputStrategy) GetInput() {
	fmt.Printf("woop")
}

func NewStdinInput() InputStrategy {
	return &StdinInputStrategy{}
}
