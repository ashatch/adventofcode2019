package puter

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
	inputValues := []int{0, 1, 2, 3, 4}

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

	ampInputA := []int{
		input[0],
		0,
	}

	ampOutputA := NewStoredOutput()
	ampOutputB := NewStoredOutput()
	ampOutputC := NewStoredOutput()
	ampOutputD := NewStoredOutput()
	ampOutputE := NewStoredOutput()

	// A
	MyPuter(NewSuppliedInput(ampInputA), ampOutputA, program)

	// B
	ampBData := []int{input[1], ampOutputA.Output[0]}
	inputAmpB := NewSuppliedInput(ampBData)
	MyPuter(inputAmpB, ampOutputB, program)

	// C
	ampCData := []int{input[2], ampOutputB.Output[0]}
	inputAmpC := NewSuppliedInput(ampCData)
	MyPuter(inputAmpC, ampOutputC, program)

	// D
	ampDData := []int{input[3], ampOutputC.Output[0]}
	inputAmpD := NewSuppliedInput(ampDData)
	MyPuter(inputAmpD, ampOutputD, program)

	// E
	ampEData := []int{input[4], ampOutputD.Output[0]}
	inputAmpE := NewSuppliedInput(ampEData)
	MyPuter(inputAmpE, ampOutputE, program)

	return ampOutputE.Output[0]
}

func AmpLoopSequence(program string, input []int) int {

	ampInputA := []int{
		input[0],
		0,
	}

	ampOutputA := NewStoredOutput()
	ampOutputB := NewStoredOutput()
	ampOutputC := NewStoredOutput()
	ampOutputD := NewStoredOutput()
	ampOutputE := NewStoredOutput()

	// A
	MyPuter(NewSuppliedInput(ampInputA), ampOutputA, program)

	// B
	ampBData := []int{input[1], ampOutputA.Output[0]}
	inputAmpB := NewSuppliedInput(ampBData)
	MyPuter(inputAmpB, ampOutputB, program)

	// C
	ampCData := []int{input[2], ampOutputB.Output[0]}
	inputAmpC := NewSuppliedInput(ampCData)
	MyPuter(inputAmpC, ampOutputC, program)

	// D
	ampDData := []int{input[3], ampOutputC.Output[0]}
	inputAmpD := NewSuppliedInput(ampDData)
	MyPuter(inputAmpD, ampOutputD, program)

	// E
	ampEData := []int{input[4], ampOutputD.Output[0]}
	inputAmpE := NewSuppliedInput(ampEData)
	MyPuter(inputAmpE, ampOutputE, program)

	return ampOutputE.Output[0]
}