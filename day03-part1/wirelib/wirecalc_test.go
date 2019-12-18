package wirelib

import (
	"testing"
)

func TestFindIntersectionsSucceeds(t *testing.T) {
	route1 := Parse("U1,R10")
	route2 := Parse("R1,U10")

	intersections := FindIntersections(route1, route2)

	if len(intersections) == 0 {
		t.Fail()
	}
}

func TestFindIntersectionsWithNoIntersections(t *testing.T) {
	route1 := Parse("R10")
	route2 := Parse("U10")

	intersections := FindIntersections(route1, route2)

	if len(intersections) != 0 {
		t.Fail()
	}
}

func TestExamples(t *testing.T) {
	route1 := Parse("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	route2 := Parse("U62,R66,U55,R34,D71,R55,D58,R83")

	dist := MinDistance(route1, route2)

	if dist != 159 {
		t.Fail()
	}
}
