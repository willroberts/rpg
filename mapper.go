package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (game *GameWorld) Preload() {
	if err := engo.Files.Load("maps/stone.tmx"); err != nil {
		panic(err)
	}
}

func (game *GameWorld) Setup(w *ecs.World) {
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	resource, err := engo.Files.Resource("maps/stone.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level
	tileComponents := make([]*Tile, 0)

	for _, tileLayer := range levelData.TileLayers {
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
				tileComponents = append(tileComponents, tile)
			}
		}
	}

	for _, imageLayer := range levelData.ImageLayers {
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
				tileComponents = append(tileComponents, tile)
			}
		}
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tileComponents {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}

	/*
		for _, objectLayer := range levelData.ObjectLayers {
			for _, object := range objectLayer.Objects {}
			for _, polylineObject := range objectLayer.PolyObjects {}
		}
	*/
}
