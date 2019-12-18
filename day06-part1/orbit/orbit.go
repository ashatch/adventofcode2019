package orbit

func NewSystem(rootLabel string) *OrbitalSystem {
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
		Label:      label,
		Orbiting:   nil,
		Satellites: []*Body{},
	}
	return body
}

func addSatellite(body *Body, satellite *Body) {
	body.Satellites = append(body.Satellites, satellite)
	satellite.Orbiting = body
}

func breadthFirst(body *Body, depth int, stats *OrbitStats) {
	for _, b := range body.Satellites {
		stats.DirectOrbitCount++
		stats.IndirectOrbitCount += depth
		breadthFirst(b, depth+1, stats)
	}
}

/*
Count does things
*/
func Count(system *OrbitalSystem) *OrbitStats {
	stats := &OrbitStats{
		DirectOrbitCount:   0,
		IndirectOrbitCount: 0,
	}

	breadthFirst(system.root, 0, stats)

	return stats
}
