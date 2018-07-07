package scene

import (
	"image/color"

	"github.com/willroberts/rpg/internal/camera"
	"github.com/willroberts/rpg/internal/char"
	"github.com/willroberts/rpg/internal/grid"
	"github.com/willroberts/rpg/internal/input"
	"github.com/willroberts/rpg/internal/sprite"
	"github.com/willroberts/rpg/internal/tilemap"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"go.uber.org/zap"
)

const (
	MapFile        string = "maps/stone.tmx"
	CharSpriteFile string = "spritesheets/characters-32x32.png"
	DecoSpriteFile string = "spritesheets/decoration-20x20-40x40.png"
)

type GameScene struct {
	Logger      *zap.Logger
	World       *ecs.World
	Map         *tilemap.Map
	Grid        grid.Grid
	CharSprites *common.Spritesheet
	DecoSprites *common.Spritesheet
	Player      char.Character
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

	scene.Logger.Info("preload complete")
}

func (scene *GameScene) Setup(u engo.Updater) {
	scene.Logger.Info("starting setup")

	// World
	w, ok := u.(*ecs.World)
	if !ok {
		scene.Logger.Error("world is not engo.Updater")
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
	cs, err := sprite.LoadSpritesheet(CharSpriteFile, 32, 32)
	if err != nil {
		scene.Logger.Error("failed to load character sprites",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.CharSprites = cs

	ds, err := sprite.LoadSpritesheet(DecoSpriteFile, 40, 40)
	if err != nil {
		scene.Logger.Error("failed to load decoration sprites",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.DecoSprites = ds

	// Player
	scene.Player = char.NewPlayer("Edmund", 1, 1, scene.CharSprites.Cell(spriteWhiteZombie))
	scene.Grid.GetCell(1, 1).SetOccupant(scene.Player)

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
