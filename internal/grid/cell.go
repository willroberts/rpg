package grid

type Container interface{}

type Occupier interface{}

type Cell interface {
	IsTraversable() bool

	GetContents() Container
	SetContents(Container)
	RemoveContents()

	GetOccupant() Occupier
	SetOccupant(Occupier)
	RemoveOccupant()
}

type cell struct {
	traversable bool
	contents    Container
	occupant    Occupier
}

func (c *cell) IsTraversable() bool {
	return c.traversable
}

func (c *cell) GetContents() Container {
	return c.contents
}

func (c *cell) SetContents(cc Container) {
	c.contents = cc
}

func (c *cell) RemoveContents() {
	c.contents = nil
}

func (c *cell) GetOccupant() Occupier {
	return c.occupant
}

func (c *cell) SetOccupant(o Occupier) {
	c.occupant = o
}

func (c *cell) RemoveOccupant() {
	c.occupant = nil
}

func NewCell(traversable bool, contents Container, occupant Occupier) Cell {
	return &cell{
		traversable: traversable,
		contents:    contents,
		occupant:    occupant,
	}
}
