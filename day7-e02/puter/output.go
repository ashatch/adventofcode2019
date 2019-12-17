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

// channel output

/*
ChannelOutputStrategy for recording output
*/
type ChannelOutputStrategy struct {
	Output chan int
}

/*
SendOutput for stored output strategy
*/
func (s *ChannelOutputStrategy) SendOutput(value int) {
	s.Output <- value
}

/*
NewChannelOutput creates a storing output
*/
func NewChannelOutput(c chan int) *ChannelOutputStrategy {
	return &ChannelOutputStrategy{
		Output: c,
	}
}

// printing channel output

/*
PrintingChannelOutputStrategy for recording output
*/
type PrintingChannelOutputStrategy struct {
	Output chan int
}

/*
PrintingChannelOutputStrategy for stored output strategy
*/
func (s *PrintingChannelOutputStrategy) SendOutput(value int) {
	fmt.Println("chan sending output", value)
	s.Output <- value
	fmt.Println("chan output sent")
}

/*
NewPrintingChannelOutput creates a printing channel output
*/
func NewPrintingChannelOutput(c chan int) *PrintingChannelOutputStrategy {
	return &PrintingChannelOutputStrategy{
		Output: c,
	}
}
