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
	if route.route[0].startSteps != 0 {
		t.Fail()
	}
	if route.route[0].endSteps != 10 {
		t.Fail()
	}
	if route.route[0].length != 10 {
		t.Fail()
	}
	if route.route[1].startSteps != 10 {
		t.Fail()
	}
	if route.route[1].endSteps != 110 {
		t.Fail()
	}
	if route.route[1].length != 100 {
		t.Fail()
	}
}

func TestPositions(t *testing.T) {
	route := Parse("R10,U10,L5,D1")

	if route.route[3].to.x != 5 {
		t.Fail()
	}

	if route.route[3].to.y != 9 {
		t.Fail()
	}
}

func TestIndices(t *testing.T) {
	route := Parse("L10,U100,D10,U5")

	for i := 0; i < 4; i++ {
		if route.route[i].index != i {
			t.Fail()
		}
	}
}

func TestHorizontalVertical(t *testing.T) {
	route := Parse("L10,U100,R10,D5")

	if !route.route[0].horizontal {
		t.Fail()
	}

	if route.route[1].horizontal {
		t.Fail()
	}

	if !route.route[2].horizontal {
		t.Fail()
	}

	if route.route[3].horizontal {
		t.Fail()
	}
}
