package rpg

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	tileScale float32 = 4.0
)

// A Tile is the basic map unit. We parse a Tiled map and create a Tile for each
// tile in the map.
type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// loadMap parses a Tiled map, returning a processed level and a set of Tiles to
// be rendered.
func loadMap(m string) (*common.Level, []*Tile, error) {
	var tiles []*Tile
	resource, err := engo.Files.Resource(m)
	if err != nil {
		return &common.Level{}, tiles, err
	}

	l := resource.(common.TMXResource).Level
	for _, tl := range l.TileLayers {
		for _, te := range tl.Tiles {
			if te.Image != nil {
				t := &Tile{BasicEntity: ecs.NewBasic()}
				t.RenderComponent = common.RenderComponent{
					Drawable: te,
					Scale:    engo.Point{tileScale, tileScale},
				}
				t.RenderComponent.SetZIndex(zWorld)
				t.SpaceComponent = common.SpaceComponent{
					Position: te.Point,
					Width:    0,
					Height:   0,
				}
				tiles = append(tiles, t)
			}
		}
	}

	return l, tiles, nil
}

// preloadMapAssets loads a Tiled map file at the given path.
func preloadMapAssets(m string) error {
	if err := engo.Files.Load(m); err != nil {
		return err
	}
	return nil
}
