// mapping.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts

// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	levelWidth  float32
	levelHeight float32
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
	tiles := make([]*Tile, 0)
	resource, err := engo.Files.Resource(m)
	if err != nil {
		return &common.Level{}, tiles, err
	}
	l := resource.(common.TMXResource).Level
	levelWidth = l.Bounds().Max.X
	levelHeight = l.Bounds().Max.Y
	log.Println("[setup] processing tile layers")
	for _, tl := range l.TileLayers {
		for _, te := range tl.Tiles {
			if te.Image != nil {
				t := &Tile{BasicEntity: ecs.NewBasic()}
				t.RenderComponent = common.RenderComponent{
					Drawable: te,
					Scale:    engo.Point{4, 4},
				}
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
	log.Println("[assets] preloading map")
	if err := engo.Files.Load(m); err != nil {
		return err
	}
	return nil
}
