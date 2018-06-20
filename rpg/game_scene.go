package rpg

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// GameScene is our first and only scene at the moment. It includes the first
// map, a static set of enemies, and only one room.
type GameScene struct{}

// Preload validates and loads assets.
func (scene *GameScene) Preload() {
	log.Println("preloading")
	preloadMapAssets("maps/stone.tmx")
	var err error
	gameSprites, err = PreloadSprites()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	gameFonts, err = PreloadFonts()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
}

// Setup initializes all systems necessary for the game to function.
func (scene *GameScene) Setup(u engo.Updater) {
	// In EngoEngine/engo#513, ecs.World ws replaced with engo.Updater.
	// This is how the authors updated their demos for compatibility.
	w, _ := u.(*ecs.World)

	log.Println("creating scene")
	gameWorld = w
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&CameraSystem{})

	log.Println("loading map")
	level, tiles, err := loadMap("maps/stone.tmx")
	if err != nil {
		log.Fatalln("error:", err.Error())
	}

	log.Println("creating level grid")
	gameGrid = newGrid(level.Width(), level.Height())

	log.Println("creating player")
	gamePlayer = newPlayer("Edmund", spriteWhiteZombie, 1, 1)
	if err := loadExperienceTable(); err != nil {
		log.Fatalln("error:", err.Error())
	}

	log.Println("creating enemies")
	if err = loadEnemyTypes(); err != nil {
		log.Fatalln("error:", err.Error())
	}
	enemies := []*Enemy{
		newEnemy("Skeleton", spriteSkeleton, 2, 7),
		newEnemy("Skeleton", spriteSkeleton, 8, 6),
		newEnemy("Skeleton", spriteSkeleton, 5, 5),
		newEnemy("Goblin", spriteGoblin, 4, 11),
		newEnemy("Goblin", spriteGoblin, 7, 12),
		newEnemy("Bear", spriteBear, 6, 17),
		newEnemy("Demon", spriteDemon, 10, 22),
	}

	log.Println("configuring systems")
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			log.Println("configuring render system")
			for _, t := range tiles {
				s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			s.Add(&gamePlayer.BasicEntity, &gamePlayer.RenderComponent, &gamePlayer.SpaceComponent)
			for _, e := range enemies {
				s.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}
		case *CameraSystem:
			log.Println("configuring camera system")
			s.Add(&gamePlayer.BasicEntity, &gamePlayer.CameraComponent,
				&gamePlayer.SpaceComponent)
		}
	}

	log.Println("configuring camera")
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &gamePlayer.SpaceComponent,
		TrackingBounds: level.Bounds(),
	})

	log.Println("creating hud")
	gameHUD, err = newHUD()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	gameLog = newActivityLog()
	gameLog.Update("Welcome to the game.")
	gameLog.Update("There are three skeletons near you.")
	gameLog.Update("Try moving into them to attack.")

	log.Println("binding controls")
	bindControls()
	log.Println("use the arrow keys to move")
}

// Type returns the name of the scene. This is used to satisfy engo's Scene
// interface.
func (scene *GameScene) Type() string {
	return "GameScene"
}
