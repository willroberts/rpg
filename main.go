package main

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

type Character struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	ControlComponent
}

type ControlComponent struct {
	SchemeVert  string
	SchemeHoriz string
}

type ControlSystem struct {
	//entity *Character // FIXME
	entities []controlEntity
}

type controlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
	*common.SpaceComponent
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, control *ControlComponent, space *common.SpaceComponent) {
	c.entities = append(c.entities, controlEntity{basic, control, space})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	del := -1
	// Determine if the requested entity is in our entities slice.
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			del = index
			break
		}
	}
	// If we found the entity, delete it.
	if del >= 0 {
		c.entities = append(c.entities[:del], c.entities[del+1:]...)
	}
}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		// Move the character.
		// This currently works because the character is the only entity.
		// As soon as we add another entity, the movement needs to be changed
		// to a speed model (based on dt) instead of a fixed rate model.
		if engo.Input.Button("moveup").JustPressed() {
			e.SpaceComponent.Position.Y -= 80
		}
		if engo.Input.Button("movedown").JustPressed() {
			e.SpaceComponent.Position.Y += 80
		}
		if engo.Input.Button("moveleft").JustPressed() {
			e.SpaceComponent.Position.X -= 80
		}
		if engo.Input.Button("moveright").JustPressed() {
			e.SpaceComponent.Position.X += 80
		}

		// Detect when the player attempts to move outside the game window.
		var heightLimit float32 = levelHeight - e.SpaceComponent.Height
		var widthLimit float32 = levelWidth - e.SpaceComponent.Width

		if e.SpaceComponent.Position.Y < 0 {
			// Dirty hack: retain the character's Y offset
			e.SpaceComponent.Position.Y = 4
		} else if e.SpaceComponent.Position.Y > heightLimit {
			// Dirty hack: retain the character's Y offset
			e.SpaceComponent.Position.Y = heightLimit + 4
		}

		if e.SpaceComponent.Position.X < 0 {
			// Dirty hack: retain the character's X offset
			e.SpaceComponent.Position.X = 6
		} else if e.SpaceComponent.Position.X > widthLimit {
			// Dirty hack: retain the character's X offset
			e.SpaceComponent.Position.X = widthLimit + 6
		}
	}
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

func main() {
	log.Println("[engine] configuring engine")
	opts := engo.RunOptions{
		Title:  "RPG",
		Width:  960,
		Height: 720,
	}
	log.Println("[engine] starting game")
	engo.Run(opts, &DefaultScene{})
}
