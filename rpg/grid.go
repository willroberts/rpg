// grid.go
package rpg

import (
	"errors"
	"log"
)

var (
	grid *Grid
)

// Grid is an array of arrays of cells. Cells may hold one entity pointer, and
// an array of item pointers.
type Grid struct {
	Rows []*GridRow

	MinX int
	MaxX int
	MinY int
	MaxY int
}

type GridRow struct {
	Cells []*GridCell
}

type GridCell struct {
	X, Y int

	Character Character
	Items     []*Item
}

func NewGrid(x, y int) *Grid {
	rows := make([]*GridRow, x)
	for i := 0; i < x; i++ {
		cells := make([]*GridCell, y)
		for j := 0; j < y; j++ {
			cells[j] = &GridCell{X: j, Y: i}
		}
		rows[i] = &GridRow{Cells: cells}
	}
	return &Grid{Rows: rows, MinX: 0, MaxX: x - 1, MinY: 0, MaxY: y - 1}
}

func (g *Grid) GetCell(x, y int) *GridCell {
	return g.Rows[y].Cells[x]
}

// Add a character to the grid. Raises an error on failure, as this should not
// happen during gameplay.
func (g *Grid) AddCharacter(c Character, atX, atY int) error {
	targetCell := g.GetCell(atX, atY)
	if targetCell.Character != nil {
		return errors.New("cannot add character to occupied grid cell")
	}
	targetCell.Character = c
	return nil
}

// Move an existing character to a new location.
func (g *Grid) MoveCharacter(c Character, toX, toY int) {
	// Check to see if anything is already there.
	targetCell := g.GetCell(toX, toY)
	if targetCell.Character != nil {
		log.Println("Something is already there!")
		if targetCell.Character.GetHostility() == "hostile" {
			log.Println("And it's hostile! Oh no!")
			// Create a combat event
			InitiateCombat(c, targetCell.Character)
		}
		return
	}

	// Clear the Character pointer from the original position
	startingCell := g.GetCell(c.GetX(), c.GetY())
	startingCell.Character = nil

	// Write the Character pointer to the new position
	targetCell.Character = c

	// Update the character's coordinates
	c.SetX(toX)
	c.SetY(toY)
}
