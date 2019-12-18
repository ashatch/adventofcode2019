package wirelib

import (
	"fmt"
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

func TestStepsToIntersection(t *testing.T) {
	route1 := Parse("U2,R10")
	route2 := Parse("R1,U10")

	intersections := FindIntersections(route1, route2)
	steps := StepsToIntersection(intersections[0])

	if steps != 6 {
		t.Fail()
	}
}

func TestStepsToIntersectionWithRandD(t *testing.T) {
	route1 := Parse("R8,U5,L5")
	route2 := Parse("U7,R6,D4")

	intersections := FindIntersections(route1, route2)
	steps := StepsToIntersection(intersections[0])
	fmt.Println(steps)

	if steps != 30 {
		t.Fail()
	}
}
