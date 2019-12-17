package puter

import "fmt"

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
	fmt.Println(program)
	return 0
	// var maxOutput int
	// inputValues := []int{0, 1, 2, 3, 4}

	// permutations := generatePermutations(inputValues)

	// for input := range permutations {
	// 	output := AmpSequence(program, input)
	// 	if output > maxOutput {
	// 		maxOutput = output
	// 	}
	// }

	// return maxOutput
}

func AmpSequence(program string, input []int) {
	ampOutputA := NewChannelOutput(make(chan int))
	ampOutputB := NewChannelOutput(make(chan int))
	ampOutputC := NewChannelOutput(make(chan int))
	ampOutputD := NewChannelOutput(make(chan int))
	ampOutputE := NewPrintingChannelOutput(make(chan int))

	inputAmpA := NewChannelInput(ampOutputE.Output)
	inputAmpB := NewChannelInput(ampOutputA.Output)
	inputAmpC := NewChannelInput(ampOutputB.Output)
	inputAmpD := NewChannelInput(ampOutputC.Output)
	inputAmpE := NewChannelInput(ampOutputD.Output)

	// A
	go MyPuter("A", inputAmpA, ampOutputA, program)

	// B
	go MyPuter("B", inputAmpB, ampOutputB, program)

	// C
	go MyPuter("C", inputAmpC, ampOutputC, program)

	// D
	go MyPuter("D", inputAmpD, ampOutputD, program)

	// E
	go MyPuter("E", inputAmpE, ampOutputE, program)

	inputAmpA.Input <- 9
	inputAmpB.Input <- 8
	inputAmpC.Input <- 7
	inputAmpD.Input <- 6
	inputAmpE.Input <- 5

	inputAmpA.Input <- 0
}
