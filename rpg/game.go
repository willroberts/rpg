package rpg

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// gameScene is a pointer to the GameScene type in the global
// namespace. This is currently stored this way so subsystems can
// access shared resources like the combat log.
var gameScene *GameScene

// GameScene is our first and only scene at the moment. It includes the first
// map, a static set of enemies, and only one room.
type GameScene struct {
	World           *ecs.World
	Fonts           *FontSet
	Sprites         *SpriteSet
	Grid            *Grid
	HUD             *HUD
	Log             *ActivityLog
	Player          *Player
	EnemyTypes      map[string]EnemyAttributes
	ExperienceTable map[string]int
}

// Preload validates and loads assets.
func (scene *GameScene) Preload() {
	log.Println("preloading")

	// Store a pointer to our GameScene in the global namespace.
	gameScene = scene

	preloadMapAssets("maps/stone.tmx")
	var err error
	scene.Sprites, err = PreloadSprites()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	scene.Fonts, err = PreloadFonts()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
}

// Setup initializes all systems necessary for the game to function.
// TODO: This function is too large -- break it up.
func (scene *GameScene) Setup(u engo.Updater) {
	// In EngoEngine/engo#513, ecs.World ws replaced with engo.Updater.
	// This is how the authors updated their demos for compatibility.
	w, _ := u.(*ecs.World)

	log.Println("creating scene")
	scene.World = w
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&CameraSystem{})

	log.Println("loading map")
	level, tiles, err := loadMap("maps/stone.tmx")
	if err != nil {
		log.Fatalln("error:", err.Error())
	}

	log.Println("creating level grid")
	scene.Grid = newGrid(level.Width(), level.Height())

	log.Println("creating player")
	gameScene.Player = newPlayer("Edmund", spriteWhiteZombie, 1, 1)
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
			s.Add(&gameScene.Player.BasicEntity, &gameScene.Player.RenderComponent, &gameScene.Player.SpaceComponent)
			for _, e := range enemies {
				s.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}
		case *CameraSystem:
			log.Println("configuring camera system")
			s.Add(&gameScene.Player.BasicEntity, &gameScene.Player.CameraComponent,
				&gameScene.Player.SpaceComponent)
		}
	}

	log.Println("configuring camera")
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &gameScene.Player.SpaceComponent,
		TrackingBounds: level.Bounds(),
	})

	log.Println("creating hud")
	gameScene.HUD, err = newHUD()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	scene.Log = newActivityLog()
	scene.Log.Update("Welcome to the game.")
	scene.Log.Update("There are three skeletons near you.")
	scene.Log.Update("Try moving into them to attack.")

	log.Println("binding controls")
	bindControls()
	log.Println("use the arrow keys to move")
}

// Type returns the name of the scene. This is used to satisfy engo's Scene
// interface.
func (scene *GameScene) Type() string {
	return "GameScene"
}
