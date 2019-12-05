package lib

import "testing"

func TestMassCalc(t *testing.T) {
	type test struct {
		input    int
		expected int
	}

	tests := []test{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range tests {
		fuelForModule := FuelForModule(tc.input)
		if fuelForModule != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, fuelForModule)
		}
	}
}
