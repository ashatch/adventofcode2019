package orbit

/*
Declaration identifies a body with its satellite
*/
type Declaration struct {
	body      string
	satellite string
}

/*
Body floats in space
*/
type Body struct {
	Label      string
	Orbiting   *Body
	Satellites []*Body
}

type OrbitalSystem struct {
	root   *Body
	bodies map[string]*Body
}

type OrbitStats struct {
	DirectOrbitCount   int
	IndirectOrbitCount int
}
