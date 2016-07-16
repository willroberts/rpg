// grid.go
package rpg

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

	Character *Character
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
