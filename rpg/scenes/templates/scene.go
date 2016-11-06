package rpg

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type ExampleScene struct{}

func (scene *ExampleScene) Preload() {
}

func (scene *ExampleScene) Setup(w *ecs.World) {
	common.SetBackground(color.White)
	w.AddSystem(&common.RenderSystem{})
}

func (scene *ExampleScene) Type() string {
	return "ExampleScene"
}
