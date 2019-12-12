package puter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
InputStrategy is how to get user input
*/
type InputStrategy interface {
	GetInput() int
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
func (s *StdinInputStrategy) GetInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("input> ")
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	return value
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
	input []int
	index int
}

/*
GetInput for pre supplied
*/
func (s *PreSuppliedInputStrategy) GetInput() int {
	input := s.input[s.index]
	s.index++
	return input
}

/*
NewSuppliedInput creates a supplied input
*/
func NewSuppliedInput(input []int) InputStrategy {
	return &PreSuppliedInputStrategy{
		input: input,
		index: 0,
	}
}
