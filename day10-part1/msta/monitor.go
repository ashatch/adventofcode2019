package msta

import (
	"fmt"
	"math"

	"github.com/atedja/go-vector"
)

type MapCell struct {
	PrintChar   string
	HasAsteroid bool
	Position    *vector.Vector
	CanSeeCount int
}

type MapRow struct {
	Cells []*MapCell
}

type AsteroidMap struct {
	Rows []*MapRow
}

func MapCellsWithAsteroids(asteroidMap *AsteroidMap) []*MapCell {
	cells := []*MapCell{}
	for _, row := range asteroidMap.Rows {
		for _, cell := range row.Cells {
			if cell.HasAsteroid {
				cells = append(cells, cell)
			}
		}
	}
	return cells
}

func ParseRow(row string, rowIndex int) *MapRow {
	cells := []*MapCell{}

	for col := 0; col < len(row); col++ {
		printChar := string(row[col])
		vec := vector.NewWithValues([]float64{float64(col), float64(rowIndex)})

		var hasAsteroid bool = false
		if printChar == "#" {
			hasAsteroid = true
		}

		cell := &MapCell{
			PrintChar:   printChar,
			Position:    &vec,
			HasAsteroid: hasAsteroid,
			CanSeeCount: 0,
		}
		cells = append(cells, cell)

	}
	return &MapRow{
		Cells: cells,
	}
}

func ParseAsteroidMap(asteroidMap []string) *AsteroidMap {
	rows := []*MapRow{}

	for row := 0; row < len(asteroidMap); row++ {
		rows = append(rows, ParseRow(asteroidMap[row], row))
	}

	return &AsteroidMap{
		Rows: rows,
	}
}

func Equal(a, b vector.Vector) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if math.Abs(v-b[i]) > 0.00000000001 {
			return false
		}
	}
	return true
}

func CanSee(from, to *MapCell, asteroidMap *AsteroidMap) bool {

	cellsWithAsteroids := MapCellsWithAsteroids(asteroidMap)

	dir := vector.Subtract(*to.Position, *from.Position)
	unitDir := vector.Unit(dir)
	len := dir.Magnitude()

	for _, otherCell := range cellsWithAsteroids {
		if !(Equal(*otherCell.Position, *from.Position) || Equal(*otherCell.Position, *to.Position)) {
			dirOther := vector.Subtract(*otherCell.Position, *from.Position)
			unitDirOther := vector.Unit(dirOther)
			magOther := dirOther.Magnitude()
			if Equal(unitDir, unitDirOther) && magOther < len {
				return false
			}
		}
	}

	return true
}

func updateDetectorSiteCount(cellsWithAsteroids []*MapCell, asteroidMap *AsteroidMap) {
	for _, from := range cellsWithAsteroids {
		for _, to := range cellsWithAsteroids {
			if !Equal(*from.Position, *to.Position) {
				canSee := CanSee(from, to, asteroidMap)
				if canSee {
					to.CanSeeCount++
				}
			}
		}
	}
}

func ShowDetectorSite(asteroidMap *AsteroidMap) {
	cellsWithAsteroids := MapCellsWithAsteroids(asteroidMap)
	updateDetectorSiteCount(cellsWithAsteroids, asteroidMap)

	for _, row := range asteroidMap.Rows {
		for _, cell := range row.Cells {
			if cell.HasAsteroid {
				fmt.Print(cell.CanSeeCount)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func FindDetectorSite(asteroidMap *AsteroidMap) *MapCell {
	cellsWithAsteroids := MapCellsWithAsteroids(asteroidMap)
	updateDetectorSiteCount(cellsWithAsteroids, asteroidMap)
	var maxCanSee int
	var bestCell *MapCell = cellsWithAsteroids[0]
	for _, cell := range cellsWithAsteroids {
		if cell.CanSeeCount > maxCanSee {
			maxCanSee = cell.CanSeeCount
			bestCell = cell
		}
	}
	return bestCell
}
