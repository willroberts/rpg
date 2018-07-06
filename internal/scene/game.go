package scene

import (
	"image/color"

	"github.com/willroberts/rpg/internal/grid"
	"github.com/willroberts/rpg/internal/tmx"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"go.uber.org/zap"
)

const (
	MapFile string = "maps/stone.tmx"
)

type GameScene struct {
	Logger *zap.Logger
	World  *ecs.World
	Map    *tmx.Map
	Grid   grid.Grid
}

func (scene *GameScene) Preload() {
	scene.Logger.Info("starting preload")

	// Map
	if err := tmx.PreloadMap(MapFile); err != nil {
		scene.Logger.Error("failed to preload map",
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
	m, err := tmx.LoadMap(MapFile)
	if err != nil {
		scene.Logger.Error("failed to load map",
			zap.String("err", err.Error()),
		)
		return
	}
	scene.Map = m

	// Grid
	scene.Grid = grid.NewGrid(m.Level.Width(), m.Level.Height())

	// RenderSystem
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			for _, t := range m.TileSet {
				s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
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
