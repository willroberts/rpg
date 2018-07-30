package scene

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/willroberts/rpg/internal/camera"
	"github.com/willroberts/rpg/internal/char"
	"github.com/willroberts/rpg/internal/grid"
	"github.com/willroberts/rpg/internal/input"
	"github.com/willroberts/rpg/internal/sprite"
	"github.com/willroberts/rpg/internal/text"
	"github.com/willroberts/rpg/internal/tilemap"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"go.uber.org/zap"
)

const (
	// Indices in character spritesheet.
	spriteHuman = iota
	spriteGreenZombie
	spriteOliveZombie
	spriteGoblin
	spriteBear
	spriteWhiteZombie
	spriteMummy
	spriteSkeleton
	spriteDemon
	spriteGargoyle

	// Indices in decoration spritesheet.
	spriteBones           int = 3
	spriteStairsDownLeft  int = 4
	spriteAnimalBones     int = 5
	spriteTinySignpost    int = 8
	spriteGravestone      int = 9
	spriteSignpost        int = 10
	spriteStairsUpLeft    int = 12
	spriteStairsDownRight int = 13
	spriteStairsUpRight   int = 14
	spriteStairsUpLeft2   int = 15 // darker?
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	MapFile        string = "maps/stone.tmx"
	CharSpriteFile string = "spritesheets/characters-32x32.png"
	DecoSpriteFile string = "spritesheets/decoration-20x20-40x40.png"
	EnemyDataFile  string = "assets/data/enemies.json"
	TitleFontFile  string = "fonts/title.ttf"
	HUDFontFile    string = "fonts/hud.ttf"
	LogFontFile    string = "fonts/combatlog.ttf"
)

const (
	LevelEnemyCount int = 10
)

type GameScene struct {
	Logger          *zap.Logger
	World           *ecs.World
	Map             *tilemap.Map
	Grid            grid.Grid
	CharSprites     *common.Spritesheet
	DecoSprites     *common.Spritesheet
	Fonts           text.FontSet
	Player          char.Character
	Enemies         []char.Character
	EnemyAttributes map[string]*char.EnemyAttributes
}

func (scene *GameScene) Preload() {
	scene.Logger.Info("starting preload")

	// Map
	if err := tilemap.PreloadMap(MapFile); err != nil {
		scene.Logger.Error("failed to preload map",
			zap.String("err", err.Error()),
		)
		return
	}

	// Sprites
	if err := sprite.PreloadSpritesheet(CharSpriteFile); err != nil {
		scene.Logger.Error("failed to preload character sprites",
			zap.String("err", err.Error()),
		)
		return
	}
	if err := sprite.PreloadSpritesheet(DecoSpriteFile); err != nil {
		scene.Logger.Error("failed to preload decoration sprites",
			zap.String("err", err.Error()),
		)
		return
	}

	// Fonts
	tf, err := text.PreloadFont(TitleFontFile, 64)
	if err != nil {
		scene.Logger.Error("failed to preload title font",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.Fonts.TitleFont = tf

	hf, err := text.PreloadFont(HUDFontFile, 32)
	if err != nil {
		scene.Logger.Error("failed to preload HUD font",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.Fonts.HUDFont = hf

	lf, err := text.PreloadFont(LogFontFile, 24)
	if err != nil {
		scene.Logger.Error("failed to preload log font",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.Fonts.LogFont = lf

	scene.Logger.Info("preload complete")
}

func (scene *GameScene) Setup(u engo.Updater) {
	scene.Logger.Info("starting setup")

	// World
	w, ok := u.(*ecs.World)
	if !ok {
		scene.Logger.Error("world is not engo.Updater")
		return
	}
	scene.World = w

	// Map
	m, err := tilemap.LoadMap(MapFile)
	if err != nil {
		scene.Logger.Error("failed to load map",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.Map = m

	// Grid
	scene.Grid = grid.NewGrid(scene.Map.Level.Width(),
		scene.Map.Level.Height())

	// Sprites
	scene.CharSprites = sprite.LoadSpritesheet(CharSpriteFile, 32, 32)
	scene.DecoSprites = sprite.LoadSpritesheet(DecoSpriteFile, 40, 40)

	// Player
	scene.Player = char.NewPlayer("Edmund", 1, 1,
		scene.CharSprites.Cell(spriteWhiteZombie),
		scene.CharSprites.Cell(spriteGravestone))
	scene.Grid.GetCell(1, 1).SetOccupant(scene.Player)

	// Enemy Attributes
	e, err := char.LoadEnemies(EnemyDataFile)
	if err != nil {
		scene.Logger.Error("failed to load enemy data",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.EnemyAttributes = e

	// Enemies
	enemyCount := 0
	scene.Enemies = make([]char.Character, 0)
	for enemyCount < LevelEnemyCount {
		randomX := rand.Intn(scene.Map.Level.Width())
		randomY := rand.Intn(scene.Map.Level.Height())
		if scene.Grid.GetCell(randomX, randomY).GetOccupant() != nil {
			continue
		}

		attrs, ok := e["Skeleton"]
		if !ok {
			scene.Logger.Error("failed to load specific enemy",
				zap.String("enemy type", "Skeleton"),
			)
			return
		}

		spr := scene.CharSprites.Cell(spriteSkeleton)

		scene.Logger.Info("adding an enemy")
		enemyCount++
		testEnemy := char.NewEnemy("Skeleton", randomX, randomY, attrs, spr)
		scene.Grid.GetCell(randomX, randomY).SetOccupant(testEnemy)
		scene.Enemies = append(scene.Enemies, testEnemy)
	}

	// Camera
	scene.World.AddSystem(&camera.CameraSystem{})
	scene.World.AddSystem(&common.EntityScroller{
		SpaceComponent: scene.Player.GetSpaceComponent(),
		TrackingBounds: scene.Map.Level.Bounds(),
	})

	// RenderSystem
	common.SetBackground(color.Black)
	scene.World.AddSystem(&common.RenderSystem{})

	// Input
	scene.World.AddSystem(&input.InputSystem{
		Player: scene.Player,
		Grid:   scene.Grid,
	})
	input.BindControls()

	// Process all systems
	for _, sys := range scene.World.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			// Add the level tiles to the RenderSystem.
			for _, t := range m.TileSet {
				s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			// Add the Player to the RenderSystem.
			s.Add(scene.Player.GetEntity(),
				scene.Player.GetRenderComponent(),
				scene.Player.GetSpaceComponent())
			// Add Enemies to the RenderSystem.
			for _, e := range scene.Enemies {
				s.Add(e.GetEntity(), e.GetRenderComponent(), e.GetSpaceComponent())
			}
		case *camera.CameraSystem:
			// Add the Player to the CameraSystem.
			s.Add(scene.Player.GetEntity(),
				scene.Player.GetSpaceComponent(),
				scene.Player.GetCameraComponent())
		}
	}

	scene.Logger.Info("setup complete")
}

func (scene *GameScene) Type() string {
	return "GameScene"
}

func (scene *GameScene) InitLogger() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()
	scene.Logger = logger
	return nil
}

func (scene *GameScene) GameOver() {
	scene.Logger.Info("game over")
}
