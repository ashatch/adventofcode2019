package wirelib

import (
	"math"
)

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func maxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func doesIntersect(first WireStep, second WireStep) (result *Intersection, err error) {
	if first.Horizontal == second.Horizontal {
		return nil, err
	}

	var horizontal WireStep
	var vertical WireStep

	if first.Horizontal {
		horizontal = first
		vertical = second
	} else {
		horizontal = second
		vertical = first
	}

	hLeft := minOf(horizontal.From.x, horizontal.To.x)
	hRight := maxOf(horizontal.From.x, horizontal.To.x)

	vBottom := minOf(vertical.From.y, vertical.To.y)
	vTop := maxOf(vertical.From.y, vertical.To.y)

	intersects := hLeft < vertical.From.x && vertical.From.x < hRight && vBottom < horizontal.From.y && horizontal.From.y < vTop

	if !intersects {
		return nil, err
	}

	result = &Intersection{
		Horizontal: horizontal,
		Vertical:   vertical,
		X:          vertical.From.x,
		Y:          horizontal.From.y,
	}
	return result, nil
}

/*
FindIntersections finds intersections in WireRoute
*/
func FindIntersections(firstWire WireRoute, secondWire WireRoute) []Intersection {
	intersections := []Intersection{}

	for _, step1 := range firstWire.route {
		for _, step2 := range secondWire.route {
			testResult, _ := doesIntersect(step1, step2)
			if testResult != nil {
				intersections = append(intersections, *testResult)
			}
		}
	}

	return intersections
}

/*
DistanceToIntersection from origin
*/
func DistanceToIntersection(i Intersection) int {
	return int(math.Abs(float64(i.X)) + math.Abs(float64(i.Y)))
}

func MinDistance(route1 WireRoute, route2 WireRoute) int {
	intersections := FindIntersections(route1, route2)

	var minDist int = DistanceToIntersection(intersections[0])

	for _, i := range intersections {
		dist := DistanceToIntersection(i)
		if dist < minDist {
			minDist = dist
		}

	}

	return minDist
}
