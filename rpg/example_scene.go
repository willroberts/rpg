// example_scene.go
// Contains a bare minimum scene for example purposes.

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

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
