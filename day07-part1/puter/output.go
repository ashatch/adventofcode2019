package puter

import (
	"fmt"
)

/*
OutputStrategy is how to output stuff from the puter
*/
type OutputStrategy interface {
	SendOutput(value int)
}

// stdin

/*
StdoutOutputStrategy strategy for stdout
*/
type StdoutOutputStrategy struct {
}

/*
SendOutput method
*/
func (s *StdoutOutputStrategy) SendOutput(value int) {
	fmt.Println(value)
}

/*
NewStdoutOutput outputs to stdout
*/
func NewStdoutOutput() *StdoutOutputStrategy {
	return &StdoutOutputStrategy{}
}

// storedoutput

/*
StoredOutputStrategy for recording output
*/
type StoredOutputStrategy struct {
	Output []int
}

/*
SendOutput for stored output strategy
*/
func (s *StoredOutputStrategy) SendOutput(value int) {
	s.Output = append(s.Output, value)
}

/*
NewStoredOutput creates a storing output
*/
func NewStoredOutput() *StoredOutputStrategy {
	return &StoredOutputStrategy{
		Output: []int{},
	}
}
