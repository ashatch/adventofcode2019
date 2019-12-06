package puter

import (
	"bufio"
	"fmt"
	"os"
)

/*
InputStrategy is how to get user input
*/
type InputStrategy interface {
	GetInput() string
}

// stdin

/*
StdinInputStrategy strategy for stdin
*/
type StdinInputStrategy struct {
}

/*
GetInput for StdinInputStrategy
*/
func (s *StdinInputStrategy) GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("input> ")
	scanner.Scan()
	return scanner.Text()
}

/*
NewStdinInput reads from stdin
*/
func NewStdinInput() InputStrategy {
	return &StdinInputStrategy{}
}

// presupplied

/*
PreSuppliedInputStrategy encapsulates input
*/
type PreSuppliedInputStrategy struct {
	input []string
	index int
}

/*
GetInput for pre supplied
*/
func (s *PreSuppliedInputStrategy) GetInput() string {
	input := s.input[s.index]
	s.index++
	return input
}

/*
NewSuppliedInput creates a supplied input
*/
func NewSuppliedInput(input []string) InputStrategy {
	return &PreSuppliedInputStrategy{
		input: input,
		index: 0,
	}
}
