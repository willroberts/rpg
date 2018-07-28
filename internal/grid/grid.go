package grid

import "github.com/willroberts/rpg/internal/char"

type Grid interface {
	GetCell(int, int) Cell
	CanMoveTo(int, int) bool
	MoveChar(char.Character, int, int)
}

type grid struct {
	width  int
	height int
	cells  [][]Cell
}

func (g *grid) GetCell(x, y int) Cell {
	return g.cells[y][x]
}

func (g *grid) CanMoveTo(x, y int) bool {
	if x >= g.width || x < 0 {
		return false
	}

	if y >= g.height || y < 0 {
		return false
	}

	target := g.GetCell(x, y)

	if !target.IsTraversable() {
		return false
	}

	if target.GetOccupant() != nil {
		// TODO: Implement combat detection here with GetHostility().
		enemy := target.GetOccupant().(char.Character)
		g.HandleCombat(enemy)
		return false
	}

	return true
}

func (g *grid) MoveChar(c char.Character, toX int, toY int) {
	// Update the grid contents.
	g.GetCell(c.GetX(), c.GetY()).SetOccupant(nil)
	g.GetCell(toX, toY).SetOccupant(c)

	// Update the character position.
	c.SetX(toX)
	c.SetY(toY)
}

func NewGrid(w, h int) Grid {
	g := &grid{width: w, height: h}

	cells := make([][]Cell, h)
	for y := 0; y < h; y++ {
		row := make([]Cell, w)
		for x := 0; x < w; x++ {
			row[x] = NewCell(true, nil, nil)
		}
		cells[y] = row
	}
	g.cells = cells

	return g
}
