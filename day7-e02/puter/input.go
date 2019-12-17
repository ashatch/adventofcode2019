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
	Close()
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

func (s *StdinInputStrategy) Close() {
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

func (s *PreSuppliedInputStrategy) Close() {
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

// channel

/*
ChannelInputStrategy encapsulates input
*/
type ChannelInputStrategy struct {
	Input  chan int
	Closed bool
	Output *ChannelOutputStrategy
}

/*
GetInput for ChannelInputStrategy
*/
func (s *ChannelInputStrategy) GetInput() int {
	input := <-s.Input
	return input
}

func (s *ChannelInputStrategy) Close() {
	s.Closed = true
	close(s.Input)
}

/*
NewChannelInput creates channel based input
*/
func NewChannelInput(c chan int) *ChannelInputStrategy {
	return &ChannelInputStrategy{
		Input:  c,
		Closed: false,
	}
}
