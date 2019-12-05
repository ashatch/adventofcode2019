package wirelib

// Pos is a position
type Pos struct {
	x int
	y int
}

// WireStep is a placement of wire on the grid
type WireStep struct {
	from       Pos
	to         Pos
	length     int
	startSteps int
	endSteps   int
	index      int
	horizontal bool
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
}
