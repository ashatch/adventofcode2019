package puter

import "testing"

func TestInputStrategy(t *testing.T) {
	arr := []int{
		1,
		2,
		3,
	}

	input := NewSuppliedInput(arr)

	for i := 0; i < len(arr); i++ {
		result := input.GetInput()
		if result != arr[i] {
			t.Fail()
		}
	}
}
