package msta

import (
	"fmt"
	"testing"

	"github.com/ashatch/vector"
	"github.com/stretchr/testify/assert"
)

func TestParseRow(t *testing.T) {
	asteroidMapData := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}

	expectedFirstRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(0.0, 0.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(1.0, 0.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(2.0, 0.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(3.0, 0.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(4.0, 0.0)},
	}

	expectedSecondRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(0.0, 1.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(1.0, 1.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(2.0, 1.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(3.0, 1.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(4.0, 1.0)},
	}

	expectedThirdRow := []*MapCell{
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(0.0, 2.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(1.0, 2.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(2.0, 2.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(3.0, 2.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(4.0, 2.0)},
	}

	expectedFourthRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(0.0, 3.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(1.0, 3.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(2.0, 3.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(3.0, 3.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(4.0, 3.0)},
	}

	expectedFifthRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(0.0, 4.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(1.0, 4.0)},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: vector.NewVector2f(2.0, 4.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(3.0, 4.0)},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: vector.NewVector2f(4.0, 4.0)},
	}

	asteroidMap := ParseAsteroidMap(asteroidMapData)

	assert.Equal(t, expectedFirstRow, asteroidMap.Rows[0].Cells)
	assert.Equal(t, expectedSecondRow, asteroidMap.Rows[1].Cells)
	assert.Equal(t, expectedThirdRow, asteroidMap.Rows[2].Cells)
	assert.Equal(t, expectedFourthRow, asteroidMap.Rows[3].Cells)
	assert.Equal(t, expectedFifthRow, asteroidMap.Rows[4].Cells)
}

func TestMapCellsWithAsteroids(t *testing.T) {
	asteroidMapData := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}

	asteroidMap := ParseAsteroidMap(asteroidMapData)

	cellsWithAsteroids := MapCellsWithAsteroids(asteroidMap)
	assert.Equal(t, 10, len(cellsWithAsteroids))
}

func TestCanSee(t *testing.T) {
	asteroidMapData := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
	asteroidMap := ParseAsteroidMap(asteroidMapData)

	/*
	 .a..b
	 .....
	 cdefg
	 ....h
	 ...ij
	*/

	a := &MapCell{Position: vector.NewVector2f(1.0, 0.0)}
	b := &MapCell{Position: vector.NewVector2f(4.0, 0.0)}
	c := &MapCell{Position: vector.NewVector2f(0.0, 2.0)}
	d := &MapCell{Position: vector.NewVector2f(1.0, 2.0)}
	e := &MapCell{Position: vector.NewVector2f(2.0, 2.0)}
	f := &MapCell{Position: vector.NewVector2f(3.0, 2.0)}
	g := &MapCell{Position: vector.NewVector2f(4.0, 2.0)}
	h := &MapCell{Position: vector.NewVector2f(4.0, 3.0)}
	i := &MapCell{Position: vector.NewVector2f(3.0, 4.0)}
	j := &MapCell{Position: vector.NewVector2f(4.0, 4.0)}

	assert.True(t, CanSee(a, b, asteroidMap))
	assert.True(t, CanSee(b, a, asteroidMap))
	assert.True(t, CanSee(b, g, asteroidMap))
	assert.False(t, CanSee(b, h, asteroidMap))
	assert.False(t, CanSee(g, j, asteroidMap))

	assert.False(t, CanSee(i, a, asteroidMap))
	iCanSee := []*MapCell{b, c, d, e, f, g, h, j}
	for _, shouldSee := range iCanSee {
		assert.True(t, CanSee(i, shouldSee, asteroidMap))
	}
}

func TestFindDetectorSite(t *testing.T) {
	asteroidMapData := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}
	asteroidMap := ParseAsteroidMap(asteroidMapData)

	expectedPosition := vector.NewVector2f(3.0, 4.0)

	site := FindDetectorSite(asteroidMap)
	assert.Equal(t, 8, site.CanSeeCount)
	assert.Equal(t, expectedPosition, site.Position)
}

func TestCellsOrderedByRotationFromPos(t *testing.T) {
	asteroidMapData := []string{
		".#....#####...#..",
		"##...##.#####..##",
		"##...#...#.#####.",
		"..#.....#...###..",
		"..#.#.....#....##",
	}

	asteroidMap := ParseAsteroidMap(asteroidMapData)
	laserPosition := &MapCell{
		Position: vector.NewVector2f(8.0, 3.0),
	}

	cells := CellsOrderedByRotationFromPos(asteroidMap, laserPosition)

	fmt.Println("zapping")
	Zap(laserPosition, asteroidMap, cells)

}
