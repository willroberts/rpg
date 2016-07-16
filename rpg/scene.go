// scene.go
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

	log.Println("[setup] processing grid")
	grid = NewGrid(level.Width(), level.Height())

	log.Println("[setup] creating player")
	player = NewPlayer(4, 2, spriteWhiteZombie)

	log.Println("[setup] creating enemies")
	enemies := []*Enemy{
		NewEnemy(5, 6, spriteSkeleton),
	}

	log.Println("[setup] configuring systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.CollisionSystem:
			log.Println("[setup] configuring collision system")
			sys.Add(&player.BasicEntity, &player.CollisionComponent,
				&player.SpaceComponent)
			for _, e := range enemies {
				sys.Add(&e.BasicEntity, &e.CollisionComponent, &e.SpaceComponent)
			}

		case *common.RenderSystem:
			log.Println("[setup] configuring render system")
			for _, t := range tiles {
				sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			sys.Add(&player.BasicEntity, &player.RenderComponent,
				&player.SpaceComponent)
			for _, e := range enemies {
				sys.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}

		case *ControlSystem:
			log.Println("[setup] configuring control system")
			sys.Add(&player.BasicEntity, &player.ControlComponent,
				&player.SpaceComponent)
		}
	}
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &player.SpaceComponent,
		TrackingBounds: level.Bounds(),
	})

	BindControls()
}

func (scene *DefaultScene) Type() string {
	return "DefaultScene"
}
