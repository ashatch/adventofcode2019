package lib

import "math"

/*
calculates fuel for a module
*/
func FuelForModule(mass int64) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}
