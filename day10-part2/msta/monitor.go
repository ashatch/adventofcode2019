package msta

import (
	"fmt"
	"math"
	"sort"

	"github.com/ashatch/vector"
)

type MapCell struct {
	PrintChar   string
	HasAsteroid bool
	Position    *vector.Vector2f
	CanSeeCount int
	Destroyed   bool
}

type MapRow struct {
	Cells []*MapCell
}

type AsteroidMap struct {
	Rows []*MapRow
}

type MapCellTarget struct {
	Cell    *MapCell
	Degrees float64
}

func CellsOrderedByRotationFromPos(asteroidMap *AsteroidMap, position *MapCell) []*MapCellTarget {
	cells := MapCellsWithAsteroids(asteroidMap)

	cellsWithAngles := []*MapCellTarget{}

	for _, cell := range cells {
		if !position.Position.Equals(cell.Position) {
			asteroidRelativeToLaser := vector.Subtract2f(cell.Position, position.Position)
			var angle = -1 * vector.AngleBetween(vector.DownUnit2f, asteroidRelativeToLaser)
			if angle < 0 {
				angle = 2*math.Pi + angle
			}
			cellsWithAngles = append(cellsWithAngles, &MapCellTarget{Cell: cell, Degrees: angle})
		}
	}

	sort.Slice(cellsWithAngles, func(i, j int) bool {
		return cellsWithAngles[i].Degrees < cellsWithAngles[j].Degrees
	})

	return cellsWithAngles
}

func Zap(from *MapCell, asteroidMap *AsteroidMap, asteroids []*MapCellTarget) {
	targetZapCount := len(asteroids)
	var searchIndex = 0
	for targetZapCount != searchIndex {
		if asteroids[searchIndex].Cell.Destroyed && CanSee(from, asteroids[searchIndex].Cell, asteroidMap) {
			fmt.Println("Zap", asteroids[searchIndex].Cell.Position)
			asteroids[searchIndex].Cell.Destroyed = true
		}
		searchIndex++
		searchIndex = searchIndex % len(asteroids)
	}

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
		vec := vector.NewVector2f(float64(col), float64(rowIndex))

		var hasAsteroid bool = false
		if printChar == "#" {
			hasAsteroid = true
		}

		cell := &MapCell{
			PrintChar:   printChar,
			Position:    vec,
			HasAsteroid: hasAsteroid,
			CanSeeCount: 0,
			Destroyed:   false,
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

func CanSee(from, to *MapCell, asteroidMap *AsteroidMap) bool {

	cellsWithAsteroids := MapCellsWithAsteroids(asteroidMap)

	dir := vector.Subtract2f(to.Position, from.Position)
	unitDir := vector.Normalize2f(dir)
	len := dir.Magnitude()

	for _, otherCell := range cellsWithAsteroids {
		if !otherCell.Position.Equals(from.Position) || otherCell.Position.Equals(to.Position) {
			dirOther := vector.Subtract2f(otherCell.Position, from.Position)
			unitDirOther := vector.Normalize2f(dirOther)
			magOther := dirOther.Magnitude()
			if unitDir.EqualTo(unitDirOther, 0.000001) && magOther < len && !otherCell.Destroyed {
				return false
			}
		}
	}

	return true
}

func updateDetectorSiteCount(cellsWithAsteroids []*MapCell, asteroidMap *AsteroidMap) {
	for _, from := range cellsWithAsteroids {
		for _, to := range cellsWithAsteroids {
			if !from.Position.Equals(to.Position) {
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
