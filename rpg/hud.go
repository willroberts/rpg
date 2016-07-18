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

const (
	zWorld = iota
	zCharacter
	zHUD
	zText
)

var (
	GameHUD *HUD
	HUDFont *common.Font
)

// The HUD contains all on-screen text, controls, and buttons.
type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// UpdateHealth redraws the HP text on the screen when Player HP changes.
func (h *HUD) UpdateHealth() {
	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(h.BasicEntity)
			h.RenderComponent.Drawable = common.Text{
				Font: HUDFont,
				Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
			}
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
}

// initializeFonts creates the various sizes of fonts we need.
func initializeHUDFont() error {
	if HUDFont == nil {
		HUDFont = &common.Font{
			URL:  "fonts/hud.ttf",
			BG:   color.Black,
			FG:   color.White,
			Size: 48,
		}
	}
	if err := HUDFont.CreatePreloaded(); err != nil {
		return err
	}
	return nil
}

// newHUD configures and returns a HUD system.
func newHUD() (*HUD, error) {
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: HUDFont,
		Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
	}
	h.RenderComponent.SetZIndex(zHUD)
	h.SetShader(common.HUDShader)
	h.SpaceComponent = common.SpaceComponent{Position: engo.Point{16, 16}}
	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
	return h, nil
}
