package wirelib

import "math"

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

type intersectionTestResult struct {
	horizontal WireStep
	vertical   WireStep
	x          int
	y          int
}

func doesIntersect(first WireStep, second WireStep) (result *intersectionTestResult, err error) {
	if first.horizontal == second.horizontal {
		return nil, err
	}

	var horizontal WireStep
	var vertical WireStep

	if first.horizontal {
		horizontal = first
		vertical = second
	} else {
		horizontal = second
		vertical = first
	}

	hLeft := minOf(horizontal.from.x, horizontal.to.x)
	hRight := maxOf(horizontal.from.x, horizontal.to.x)

	vBottom := minOf(vertical.from.y, vertical.to.y)
	vTop := maxOf(vertical.from.y, vertical.to.y)

	intersects := hLeft < vertical.from.x && vertical.from.x < hRight && vBottom < horizontal.from.y && horizontal.from.y < vTop

	if !intersects {
		return nil, err
	}

	result = &intersectionTestResult{
		horizontal: horizontal,
		vertical:   vertical,
		x:          vertical.from.x,
		y:          horizontal.from.y,
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
				i := Intersection{
					Horizontal: testResult.horizontal,
					Vertical:   testResult.vertical,
					X:          testResult.x,
					Y:          testResult.y,
				}
				intersections = append(intersections, i)
			}
		}
	}

	return intersections
}

/*
DistanceToIntersection from origin
*/
func DistanceToIntersection(i Intersection) int64 {
	return int64(math.Abs(float64(i.X)) + math.Abs(float64(i.Y)))
}
