// label.go
// Contains types and logic for drawable text labels.

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
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Label represents a drawable text label on the screen.
type Label struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Creates and returns a new Label.
func NewLabel(text string, font *common.Font, x, y float32) *Label {
	l := &Label{BasicEntity: ecs.NewBasic()}
	l.RenderComponent.Drawable = common.Text{Font: font, Text: text}
	l.SetShader(common.HUDShader)
	l.SpaceComponent = common.SpaceComponent{Position: engo.Point{x, y}}
	return l
}
