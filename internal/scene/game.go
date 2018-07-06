package scene

import (
	"image/color"

	"github.com/willroberts/rpg/internal/grid"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"go.uber.org/zap"
)

type GameScene struct {
	Logger *zap.Logger
	World  *ecs.World
	Grid   grid.Grid
}

func (scene *GameScene) Preload() {
	scene.Logger.Info("starting preload")
	scene.Logger.Info("preload complete")
}

func (scene *GameScene) Setup(u engo.Updater) {
	scene.Logger.Info("starting setup")

	w, ok := u.(*ecs.World)
	if !ok {
		scene.Logger.Error("world is not engo.Updater")
	}
	scene.World = w

	// Replace these values with level.Width()/.Height().
	scene.Grid = grid.NewGrid(10, 10)

	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})

	//for _, sys := range w.Systems() {
	//	switch s := sys.(type) {
	//	case *common.RenderSystem:
	//	}
	//}

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
