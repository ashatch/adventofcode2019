package puter

import (
	"fmt"
	"sync"
)

func generatePermutations(data []int) <-chan []int {
	c := make(chan []int)
	go func(c chan []int) {
		defer close(c)
		permutate(c, data)
	}(c)
	return c
}
func permutate(c chan []int, inputs []int) {
	output := make([]int, len(inputs))
	copy(output, inputs)
	c <- output

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int, len(inputs))
		copy(output, inputs)
		c <- output
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func FindMaxAmpSequence(program string) int {
	var maxOutput int
	inputValues := []int{5, 6, 7, 8, 9}

	permutations := generatePermutations(inputValues)

	for input := range permutations {
		output := AmpSequence(program, input)
		if output > maxOutput {
			maxOutput = output
		}
	}

	return maxOutput
}

func AmpSequence(program string, input []int) int {
	var wg sync.WaitGroup

	inputAmpA := NewChannelInput(make(chan int))
	inputAmpB := NewChannelInput(make(chan int))
	inputAmpC := NewChannelInput(make(chan int))
	inputAmpD := NewChannelInput(make(chan int))
	inputAmpE := NewChannelInput(make(chan int))

	ampOutputA := NewChannelOutput(inputAmpB)
	ampOutputB := NewChannelOutput(inputAmpC)
	ampOutputC := NewChannelOutput(inputAmpD)
	ampOutputD := NewChannelOutput(inputAmpE)
	ampOutputE := NewChannelOutput(inputAmpA)

	// A
	wg.Add(1)
	go MyPuter(&wg, "A", inputAmpA, ampOutputA, program)

	// B
	wg.Add(1)
	go MyPuter(&wg, "B", inputAmpB, ampOutputB, program)

	// C
	wg.Add(1)
	go MyPuter(&wg, "C", inputAmpC, ampOutputC, program)

	// D
	wg.Add(1)
	go MyPuter(&wg, "D", inputAmpD, ampOutputD, program)

	// E
	wg.Add(1)
	go MyPuter(&wg, "E", inputAmpE, ampOutputE, program)

	inputAmpA.Input <- input[0]
	inputAmpB.Input <- input[1]
	inputAmpC.Input <- input[2]
	inputAmpD.Input <- input[3]
	inputAmpE.Input <- input[4]

	inputAmpA.Input <- 0

	wg.Wait()

	fmt.Println("Thruster output:", ampOutputE.LastValue)
	return ampOutputE.LastValue
}
