package main

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type GameWorld struct{}

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type Character struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type ControlSystem struct {
	entity *Character
}

func (c *ControlSystem) Add(char *Character) {
	c.entity = char
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	if c.entity != nil && basic.ID() == c.entity.ID() {
		c.entity = nil
	}
}

func (c *ControlSystem) Update(dt float32) {
	if engo.Input.Button("moveup").JustPressed() {
		c.entity.SpaceComponent.Position.Y -= 80
	}
	if engo.Input.Button("movedown").JustPressed() {
		c.entity.SpaceComponent.Position.Y += 80
	}
	if engo.Input.Button("moveleft").JustPressed() {
		c.entity.SpaceComponent.Position.X -= 80
	}
	if engo.Input.Button("moveright").JustPressed() {
		c.entity.SpaceComponent.Position.X += 80
	}
}

func (game *GameWorld) Preload() {
	log.Println("preloading resources")
	log.Println("loading maps")
	if err := engo.Files.Load("maps/stone_tall.tmx"); err != nil {
		panic(err)
	}
	log.Println("loading sprites")
	if err := engo.Files.Load("spritesheets/characters-32x32.png"); err != nil {
		panic(err)
	}
}

func (game *GameWorld) Setup(w *ecs.World) {
	log.Println("setting up game world")
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	resource, err := engo.Files.Resource("maps/stone_tall.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level
	tileComponents := make([]*Tile, 0)

	log.Println("processing tile layers")
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

	log.Println("processing image layers")
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

	log.Println("creating character")
	character := Character{BasicEntity: ecs.NewBasic()}
	spriteSheet := common.NewSpritesheetFromFile("spritesheets/characters-32x32.png", 32, 32)
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

	log.Println("configuring systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&character.BasicEntity, &character.RenderComponent,
				&character.SpaceComponent)
			for _, v := range tileComponents {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		}
	}
	w.AddSystem(&ControlSystem{&character})
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &character.SpaceComponent,
		TrackingBounds: levelData.Bounds(),
	})

	log.Println("binding controls")
	engo.Input.RegisterButton("moveup", engo.ArrowUp)
	engo.Input.RegisterButton("movedown", engo.ArrowDown)
	engo.Input.RegisterButton("moveleft", engo.ArrowLeft)
	engo.Input.RegisterButton("moveright", engo.ArrowRight)
	log.Println("controls bound")
	log.Println("use the arrow keys to move")
}

func (game *GameWorld) Type() string {
	return "GameWorld"
}

func main() {
	log.Println("configuring game")
	opts := engo.RunOptions{
		Title:  "RPG",
		Width:  960,
		Height: 720,
	}
	log.Println("starting game")
	engo.Run(opts, &GameWorld{})
}
