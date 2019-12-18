package wirelib

import (
	"strconv"
	"strings"
)

// Parse a route
func Parse(instructions string) WireRoute {
	instructionList := strings.Split(instructions, ",")

	route := WireRoute{
		route: []WireStep{},
	}

	currentPosition := Pos{
		x: 0,
		y: 0,
	}

	currentSteps := 0

	for _, op := range instructionList {
		opCode := op[0:1]
		length, _ := strconv.Atoi(op[1:])

		step := WireStep{
			From: Pos{
				x: currentPosition.x,
				y: currentPosition.y,
			},
			To: Pos{
				x: currentPosition.x,
				y: currentPosition.y,
			},
			Length:     length,
			Direction:  1,
			StartSteps: currentSteps,
			EndSteps:   currentSteps + length,
			Index:      len(route.route),
			Horizontal: false,
		}

		switch opCode {
		case "L":
			{
				step.Horizontal = true
				step.To.x -= length
				step.Direction = -1
				route.route = append(route.route, step)

				currentPosition.x -= length
				currentSteps += length
			}
		case "R":
			{
				step.Horizontal = true
				step.To.x += length
				route.route = append(route.route, step)

				currentPosition.x += length
				currentSteps += length
			}
		case "U":
			{
				step.Horizontal = false
				step.To.y += length
				route.route = append(route.route, step)

				currentPosition.y += length
				currentSteps += length
			}
		case "D":
			{
				step.Horizontal = false
				step.To.y -= length
				step.Direction = -1
				route.route = append(route.route, step)

				currentPosition.y -= length
				currentSteps += length
			}
		}
	}

	return route
}
