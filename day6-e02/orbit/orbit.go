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

func Path(system *OrbitalSystem, a string) []*Body {
	path := []*Body{}
	var current *Body
	current = system.bodies[a]

	for current.Orbiting != nil {
		path = append(path, current)
		current = current.Orbiting
	}

	path = append(path, current)

	return path
}

func PathBetween(system *OrbitalSystem, leaf string, parent string) []*Body {
	path := []*Body{}
	var current *Body
	current = system.bodies[leaf]

	for current.Label != parent {
		path = append(path, current)
		current = current.Orbiting
	}

	path = append(path, current)

	return path
}

func CommonBody(system *OrbitalSystem, a string, b string) *Body {
	pathA := Path(system, a)
	pathB := Path(system, b)

	for _, i := range pathA {
		for _, j := range pathB {
			if i == j {
				return i
			}
		}
	}

	return nil
}

func TransferCount(system *OrbitalSystem, a string, b string) int {
	commonBody := CommonBody(system, a, b)
	pathA := PathBetween(system, a, commonBody.Label)
	pathB := PathBetween(system, b, commonBody.Label)

	return (len(pathA) - 2) + (len(pathB) - 2)
}
