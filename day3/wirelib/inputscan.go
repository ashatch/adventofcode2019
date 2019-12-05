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
			from: Pos{
				x: currentPosition.x,
				y: currentPosition.y,
			},
			to: Pos{
				x: currentPosition.x,
				y: currentPosition.y,
			},
			length:     length,
			startSteps: currentSteps,
			endSteps:   currentSteps + length,
			index:      len(route.route),
			horizontal: false,
		}

		switch opCode {
		case "L":
			{
				step.horizontal = true
				step.to.x -= length
				route.route = append(route.route, step)

				currentPosition.x -= length
				currentSteps += length
			}
		case "R":
			{
				step.horizontal = true
				step.to.x += length
				route.route = append(route.route, step)

				currentPosition.x += length
				currentSteps += length
			}
		case "U":
			{
				step.horizontal = false
				step.to.y += length
				route.route = append(route.route, step)

				currentPosition.y += length
				currentSteps += length
			}
		case "D":
			{
				step.horizontal = false
				step.to.y -= length
				route.route = append(route.route, step)

				currentPosition.y -= length
				currentSteps += length
			}
		}
	}

	return route
}
