package puter

import (
	"testing"
)

func TestPuter(t *testing.T) {
	programs := []string{
		"1,0,0,0,99",
		"2,3,0,3,99",
		"2,4,4,5,99,0",
	}

	results := [][]int{
		{2, 0, 0, 0, 99},
		{2, 3, 0, 6, 99},
		{2, 4, 4, 5, 99, 9801},
	}

	for i := 0; i < len(programs); i++ {
		result := MyPuter(programs[i], false, 0, 0)
		for j := 0; j < len(results[i]); j++ {
			if result[j] != results[i][j] {
				t.Fail()
			}
		}

	}
}
