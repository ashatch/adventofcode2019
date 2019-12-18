package lib

import (
	"math"
)

/*
FuelForModule calculates fuel for a module
*/
func FuelForModule(mass int) int {
	if mass <= 0 {
		return 0
	} else {
		fuel := int(math.Max((math.Floor(float64(mass)/3.0))-2, 0))
		fuelMass := FuelForModule(fuel)
		return fuel + fuelMass
	}
}
