package wirelib

// Pos is a position
type Pos struct {
	x int
	y int
}

// WireStep is a placement of wire on the grid
type WireStep struct {
	From       Pos
	To         Pos
	Length     int
	StartSteps int
	EndSteps   int
	Index      int
	Horizontal bool
	Direction  int
}

// WireRoute is contiguous sequence of WireStep
type WireRoute struct {
	route []WireStep
}

// Intersection is where two WireStep cross
type Intersection struct {
	Horizontal WireStep
	Vertical   WireStep
	X          int
	Y          int
	StepsX     int
	StepsY     int
}
