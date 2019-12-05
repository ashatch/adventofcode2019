package wirelib

import (
	"testing"
)

func TestFindIntersectionsA(t *testing.T) {
	route1 := Parse("U1,R10")
	route2 := Parse("R1,U10")

	intersections := FindIntersections(route1, route2)

	if len(intersections) == 0 {
		t.Fail()
	}
}

func TestFindIntersectionsB(t *testing.T) {
	route1 := Parse("R10")
	route2 := Parse("U10")

	intersections := FindIntersections(route1, route2)

	if len(intersections) != 0 {
		t.Fail()
	}
}
