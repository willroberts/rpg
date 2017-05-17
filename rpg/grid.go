package rpg

import "errors"

// A Grid contains an array of arrays of cells corresponding to a Tiled map. It
// is used for level and movement data for increased accuracy over pixel-based
// approaches.
type Grid struct {
	Rows []*GridRow

	MinX int
	MaxX int
	MinY int
	MaxY int
}

// AddCharacter adds a character to the grid. Raises an error on failure, as this
// should not happen during gameplay.
func (g *Grid) AddCharacter(c Character, atX, atY int) error {
	dst := g.GetCell(atX, atY)
	if dst.Character != nil {
		return errors.New("cannot add character to occupied grid cell")
	}
	dst.Character = c
	return nil
}

// GetCell returns the GridCell at the given X and Y coordinates.
func (g *Grid) GetCell(x, y int) *GridCell {
	return g.Rows[y].Cells[x]
}

// MoveCharacter moves an existing character to a new location if nothing is
// already there. If a hostile entity is there, combat is started.
func (g *Grid) MoveCharacter(c Character, toX, toY int) {
	dst := g.GetCell(toX, toY)
	if dst.Character != nil {
		if dst.Character.GetHostility() == "hostile" {
			initiateCombat(c, dst.Character)
		}
		return
	}
	src := g.GetCell(c.GetX(), c.GetY())
	src.Character = nil
	dst.Character = c
	c.SetX(toX)
	c.SetY(toY)
}

// RemoveCharacter clears the Character entity from the cell at the given X and Y
// coordinates.
func (g *Grid) RemoveCharacter(fromX, fromY int) {
	g.GetCell(fromX, fromY).Character = nil
}

// A GridCell has X and Y coordinates, can contain one Character entity, and can
// contain multiple Items.
type GridCell struct {
	X, Y int

	Character Character
}

// A GridRow contains an array of GridCells.
type GridRow struct {
	Cells []*GridCell
}

// NewGrid returns an empty, pre-initialized Grid of the given dimensions.
func NewGrid(x, y int) *Grid {
	r := make([]*GridRow, y)
	for i := 0; i < y; i++ {
		c := make([]*GridCell, x)
		for j := 0; j < x; j++ {
			c[j] = &GridCell{X: j, Y: i}
		}
		r[i] = &GridRow{Cells: c}
	}
	return &Grid{Rows: r, MinX: 0, MaxX: x - 1, MinY: 0, MaxY: y - 1}
}
