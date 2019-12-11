package orbit

import "fmt"

func newSystem(rootLabel string) *OrbitalSystem {
	rootBody := newBody(rootLabel)

	system := &OrbitalSystem{
		root:   rootBody,
		bodies: make(map[string]*Body),
	}

	system.bodies[rootLabel] = rootBody

	return system
}

func newBody(label string) *Body {
	body := &Body{
		Label:      "COM",
		Orbiting:   nil,
		Satellites: []*Body{},
	}
	return body
}

func addSatellite(body *Body, satellite *Body) {
	body.Satellites = append(body.Satellites, satellite)
	satellite.Orbiting = body
}

/*
Count does things
*/
func Count() {
	fmt.Println("test")
	com := newBody("COM")

	planetB := newBody("planet-B")

	addSatellite(com, newBody("planet-A"))
	addSatellite(com, planetB)
	addSatellite(planetB, newBody("moon-1"))

	fmt.Println(com)
}
