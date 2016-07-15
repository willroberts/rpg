package rpg

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	levelWidth  float32
	levelHeight float32
)

type DefaultScene struct{}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (scene *DefaultScene) Preload() {
	log.Println("[assets] preloading resources")
	log.Println("[assets] loading maps")
	if err := engo.Files.Load("maps/stone.tmx"); err != nil {
		panic(err)
	}
	log.Println("[assets] loading sprites")
	if err := engo.Files.Load("spritesheets/characters-32x32.png"); err != nil {
		panic(err)
	}
}

func (scene *DefaultScene) Setup(w *ecs.World) {
	log.Println("[setup] setting up scene")

	common.SetBackground(color.Black)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&ControlSystem{})

	resource, err := engo.Files.Resource("maps/stone.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level
	levelWidth = levelData.Bounds().Max.X
	levelHeight = levelData.Bounds().Max.Y

	log.Println("[setup] processing tile layers")
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

	log.Println("[setup] processing image layers")
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

	log.Println("[setup] creating character")
	character := Character{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
	}
	spriteSheet := common.NewSpritesheetFromFile(
		"spritesheets/characters-32x32.png", 32, 32)
	skeletonTexture := spriteSheet.Cell(7)
	character.RenderComponent = common.RenderComponent{
		Drawable: skeletonTexture,
		Scale:    engo.Point{2, 2},
	}
	character.RenderComponent.SetZIndex(1)
	character.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{86, 84},
		Width:    80,
		Height:   80,
	}

	log.Println("[setup] configuring systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&character.BasicEntity, &character.RenderComponent,
				&character.SpaceComponent)
			for _, v := range tileComponents {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		case *ControlSystem:
			sys.Add(&character.BasicEntity, &character.ControlComponent,
				&character.SpaceComponent)
		}
	}
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &character.SpaceComponent,
		TrackingBounds: levelData.Bounds(),
	})

	log.Println("[input] binding controls")
	engo.Input.RegisterButton("moveup", engo.ArrowUp)
	engo.Input.RegisterButton("movedown", engo.ArrowDown)
	engo.Input.RegisterButton("moveleft", engo.ArrowLeft)
	engo.Input.RegisterButton("moveright", engo.ArrowRight)
	log.Println("[input] controls bound")
	log.Println("[input] use the arrow keys to move")
}

func (scene *DefaultScene) Type() string {
	return "DefaultScene"
}
