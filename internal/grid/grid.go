package rpg

type Grid interface {
}

type grid struct {
	width  int
	height int
	cells  [][]Cell
}

func (g *grid) GetCell(x, y int) Cell {
	return g.cells[y][x]
}

func NewGrid(w, h int) Grid {
	g := grid{width: w, height: h}

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
