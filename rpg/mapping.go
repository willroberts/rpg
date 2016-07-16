package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	grid        *Grid
	levelWidth  float32
	levelHeight float32
)

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

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func PreloadMapAssets(m string) error {
	log.Println("[assets] preloading map")
	if err := engo.Files.Load(m); err != nil {
		return err
	}
	return nil
}

func LoadMap(m string) (*common.Level, []*Tile, error) {
	tiles := make([]*Tile, 0)

	resource, err := engo.Files.Resource(m)
	if err != nil {
		return &common.Level{}, tiles, err
	}

	tmxResource := resource.(common.TMXResource)
	level := tmxResource.Level
	levelWidth = level.Bounds().Max.X
	levelHeight = level.Bounds().Max.Y

	log.Println("[setup] processing tile layers")
	for _, tileLayer := range level.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement,
					Scale:    engo.Point{4, 4},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: tileElement.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}

	log.Println("[setup] processing image layers")
	for _, imageLayer := range level.ImageLayers {
		for _, imageElement := range imageLayer.Images {
			if imageElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: imageElement,
					Scale:    engo.Point{4, 4},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: imageElement.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return level, tiles, nil
}
