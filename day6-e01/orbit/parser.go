package orbit

import (
	"strings"
)

/*
ParseDeclaration splits a string
*/
func ParseDeclaration(line string) *Declaration {
	items := strings.Split(line, ")")
	return &Declaration{
		body:      items[0],
		satellite: items[1],
	}
}

func AddDeclaration(declaration *Declaration, system *OrbitalSystem) {
	if system.bodies[declaration.body] == nil {
		system.bodies[declaration.body] = newBody(declaration.body)
	}

	if system.bodies[declaration.satellite] == nil {
		system.bodies[declaration.satellite] = newBody(declaration.body)
	}

	body := system.bodies[declaration.body]
	satellite := system.bodies[declaration.satellite]

	satellite.Orbiting = body
	addSatellite(body, satellite)
}
