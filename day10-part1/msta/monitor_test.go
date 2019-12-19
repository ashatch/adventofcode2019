package msta

import (
	"testing"

	"github.com/atedja/go-vector"
	"github.com/stretchr/testify/assert"
)

func pointerTo(v vector.Vector) *vector.Vector {
	return &v
}

func TestParseRow(t *testing.T) {
	asteroidMapData := []string{
		".#..#",
		".....",
		"#####",
		"....#",
		"...##",
	}

	expectedFirstRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{0.0, 0.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{1.0, 0.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{2.0, 0.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{3.0, 0.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{4.0, 0.0}))},
	}

	expectedSecondRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{0.0, 1.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{1.0, 1.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{2.0, 1.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{3.0, 1.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{4.0, 1.0}))},
	}

	expectedThirdRow := []*MapCell{
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{0.0, 2.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{1.0, 2.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{2.0, 2.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{3.0, 2.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{4.0, 2.0}))},
	}

	expectedFourthRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{0.0, 3.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{1.0, 3.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{2.0, 3.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{3.0, 3.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{4.0, 3.0}))},
	}

	expectedFifthRow := []*MapCell{
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{0.0, 4.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{1.0, 4.0}))},
		&MapCell{PrintChar: ".", HasAsteroid: false, Position: pointerTo(vector.NewWithValues([]float64{2.0, 4.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{3.0, 4.0}))},
		&MapCell{PrintChar: "#", HasAsteroid: true, Position: pointerTo(vector.NewWithValues([]float64{4.0, 4.0}))},
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

	a := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{1.0, 0.0}))}
	b := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{4.0, 0.0}))}
	c := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{0.0, 2.0}))}
	d := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{1.0, 2.0}))}
	e := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{2.0, 2.0}))}
	f := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{3.0, 2.0}))}
	g := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{4.0, 2.0}))}
	h := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{4.0, 3.0}))}
	i := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{3.0, 4.0}))}
	j := &MapCell{Position: pointerTo(vector.NewWithValues([]float64{4.0, 4.0}))}

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

	expectedPosition := vector.NewWithValues([]float64{3.0, 4.0})

	site := FindDetectorSite(asteroidMap)
	assert.Equal(t, 8, site.CanSeeCount)
	assert.Equal(t, expectedPosition, *site.Position)
}

// func TestExampleMap(t *testing.T) {
// 	asteroidMap := []string{
// 		".#..#",
// 		".....",
// 		"#####",
// 		"....#",
// 		"...##",
// 	}

// 	expected := []string{
// 		".7..7",
// 		".....",
// 		"67775",
// 		"....7",
// 		"...87",
// 	}

// 	detectorSites := FindDetectorSites(asteroidMap)

// 	assert.Equal(t, expected, detectorSites)
// }
