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

var GameWorld *ecs.World

func (scene *DefaultScene) Preload() {
	log.Println("[assets] preloading resources")
	PreloadMapAssets("maps/stone.tmx")
	log.Println("[assets] loading sprites")
	err := engo.Files.Load("spritesheets/characters-32x32.png")
	if err != nil {
		panic(err)
	}
	err = engo.Files.Load("spritesheets/decoration-20x20-40x40.png")
	if err != nil {
		panic(err)
	}
	characterSpritesheet = common.NewSpritesheetFromFile(
		characterSpritesheetPath,
		characterSpritesheetWidth,
		characterSpritesheetHeight)
	decorationSpritesheet = common.NewSpritesheetFromFile(
		decorationSpritesheetPath,
		decorationSpritesheetWidth,
		decorationSpritesheetHeight)
}

func (scene *DefaultScene) Setup(w *ecs.World) {
	log.Println("[setup] setting up scene")
	GameWorld = w
	common.SetBackground(color.Black)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&ControlSystem{})

	log.Println("[setup] loading map")
	level, tiles, err := LoadMap("maps/stone.tmx")
	if err != nil {
		panic(err)
	}

	log.Println("[setup] processing grid")
	grid = NewGrid(level.Width(), level.Height())

	log.Println("[setup] creating player")
	player = NewPlayer(1, 1, spriteWhiteZombie)

	log.Println("[setup] creating enemies")
	err = LoadEnemyTypes()
	if err != nil {
		panic(err)
	}
	enemies := []*Enemy{
		NewEnemy("skeleton", spriteSkeleton, 2, 7),
		NewEnemy("skeleton", spriteSkeleton, 8, 6),
		NewEnemy("skeleton", spriteSkeleton, 5, 5),
		NewEnemy("goblin", spriteGoblin, 4, 11),
		NewEnemy("goblin", spriteGoblin, 7, 12),
		NewEnemy("bear", spriteBear, 6, 17),
		NewEnemy("demon", spriteDemon, 10, 22),
	}

	log.Println("[setup] configuring systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
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
