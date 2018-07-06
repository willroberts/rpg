package tilemap

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	TileScale  float32 = 4.0
	TileZIndex float32 = 0.0
)

type Map struct {
	Level   *common.Level
	TileSet []*Tile
}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func PreloadMap(filename string) error {
	return engo.Files.Load(filename)
}

func LoadMap(filename string) (*Map, error) {
	m := &Map{}
	m.Level = &common.Level{}
	m.TileSet = make([]*Tile, 0)

	res, err := engo.Files.Resource(filename)
	if err != nil {
		return m, err
	}

	m.Level = res.(common.TMXResource).Level
	for _, tl := range m.Level.TileLayers {
		for _, t := range tl.Tiles {
			if t.Image != nil {
				tile := &Tile{
					BasicEntity: ecs.NewBasic(),
					RenderComponent: common.RenderComponent{
						Drawable: t,
						Scale: engo.Point{
							X: TileScale,
							Y: TileScale,
						},
					},
					SpaceComponent: common.SpaceComponent{
						Position: t.Point,
						Width:    0,
						Height:   0,
					},
				}
				tile.RenderComponent.SetZIndex(TileZIndex)
				m.TileSet = append(m.TileSet, tile)
			}
		}
	}

	return m, nil
}
