package wirelib

import (
	"testing"
)

func TestLength(t *testing.T) {
	route := Parse("L100")
	if len(route.route) != 1 {
		t.Error("wrong length")
	}

	route = Parse("L100,U100")
	if len(route.route) != 2 {
		t.Error("wrong length")
	}

	route = Parse("L100,U100,R1")
	if len(route.route) != 3 {
		t.Error("wrong length")
	}
}

func TestStepsAndLength(t *testing.T) {
	route := Parse("L10,U100")
	if route.route[0].StartSteps != 0 {
		t.Fail()
	}
	if route.route[0].EndSteps != 10 {
		t.Fail()
	}
	if route.route[0].Length != 10 {
		t.Fail()
	}
	if route.route[1].StartSteps != 10 {
		t.Fail()
	}
	if route.route[1].EndSteps != 110 {
		t.Fail()
	}
	if route.route[1].Length != 100 {
		t.Fail()
	}
}

func TestPositions(t *testing.T) {
	route := Parse("R10,U10,L5,D1")

	if route.route[3].To.x != 5 {
		t.Fail()
	}

	if route.route[3].To.y != 9 {
		t.Fail()
	}
}

func TestIndices(t *testing.T) {
	route := Parse("L10,U100,D10,U5")

	for i := 0; i < 4; i++ {
		if route.route[i].Index != i {
			t.Fail()
		}
	}
}

func TestHorizontalVertical(t *testing.T) {
	route := Parse("L10,U100,R10,D5")

	if !route.route[0].Horizontal {
		t.Fail()
	}

	if route.route[1].Horizontal {
		t.Fail()
	}

	if !route.route[2].Horizontal {
		t.Fail()
	}

	if route.route[3].Horizontal {
		t.Fail()
	}
}
