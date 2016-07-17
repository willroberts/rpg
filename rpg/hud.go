// hud.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts

// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
package rpg

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	GameHUD  *HUD
	GameFont *common.Font
)

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (h *HUD) UpdateHealth() {
	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(h.BasicEntity)
			h.RenderComponent.Drawable = common.Text{
				Font: GameFont,
				Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
			}
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
}

func newHUD() (*HUD, error) {
	GameFont = &common.Font{
		URL:  "fonts/Roboto-Regular.ttf",
		BG:   color.Black,
		FG:   color.White,
		Size: 48,
	}
	err := GameFont.CreatePreloaded()
	if err != nil {
		return nil, err
	}

	// Create HP display
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: GameFont,
		Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
	}
	h.SetShader(common.HUDShader)
	h.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{16, 16},
	}

	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}

	return h, nil
}
