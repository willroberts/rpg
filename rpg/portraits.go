// charselect.go
// Creates a grid of Character portraits for character creation.

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
	"engo.io/engo"
	"engo.io/engo/common"
)

type CharSelectPanel struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewCharSelectPanel() *CharSelectPanel {
	p := &CharSelectPanel{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Drawable: common.Rectangle{},
			Color:    color.RGBA{20, 20, 20, 255},
		},
	}
	p.SetShader(common.HUDShader)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{32, 200},
		Width:    320,
		Height:   160,
	}
	return p
}
