package rpg

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type DefaultScene struct{}

func (scene *DefaultScene) Preload() {
	log.Println("[assets] preloading resources")
	PreloadMapAssets("maps/stone.tmx")
	log.Println("[assets] loading sprites")
	if err := engo.Files.Load("spritesheets/characters-32x32.png"); err != nil {
		panic(err)
	}
	characterSpritesheet = common.NewSpritesheetFromFile(characterSpritesheetPath,
		characterSpritesheetWidth, characterSpritesheetHeight)
}

func (scene *DefaultScene) Setup(w *ecs.World) {
	log.Println("[setup] setting up scene")
	common.SetBackground(color.Black)

	w.AddSystem(&common.CollisionSystem{})
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&CombatSystem{})
	w.AddSystem(&ControlSystem{})

	log.Println("[setup] loading map")
	level, tiles, err := LoadMap("maps/stone.tmx")
	if err != nil {
		panic(err)
	}

	log.Println("[setup] creating character")
	character := NewCharacter(spriteWhiteZombie)

	log.Println("[setup] creating enemies")
	enemies := []Enemy{
		NewEnemy(spriteSkeleton, 240, 240),
	}

	log.Println("[setup] configuring systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.CollisionSystem:
			log.Println("[setup] configuring collision system")
			sys.Add(&character.BasicEntity, &character.CollisionComponent,
				&character.SpaceComponent)
			for _, e := range enemies {
				sys.Add(&e.BasicEntity, &e.CollisionComponent, &e.SpaceComponent)
			}

		case *common.RenderSystem:
			log.Println("[setup] configuring render system")
			for _, t := range tiles {
				sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			sys.Add(&character.BasicEntity, &character.RenderComponent,
				&character.SpaceComponent)
			for _, e := range enemies {
				sys.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}

		case *ControlSystem:
			log.Println("[setup] configuring control system")
			sys.Add(&character.BasicEntity, &character.ControlComponent,
				&character.SpaceComponent)
		}
	}
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &character.SpaceComponent,
		TrackingBounds: level.Bounds(),
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
