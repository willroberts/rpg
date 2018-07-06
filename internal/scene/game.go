package scene

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"go.uber.org/zap"
)

type GameScene struct {
	World  *ecs.World
	Logger *zap.Logger
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
