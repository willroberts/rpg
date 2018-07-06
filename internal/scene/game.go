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
}

func (scene *GameScene) Setup(u engo.Updater) {
	w, ok := u.(*ecs.World)
	if !ok {
		// Write to engine log.
	}
	scene.World = w

	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})

	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			// Draw things on screen here.
			_ = s
		}
	}
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
