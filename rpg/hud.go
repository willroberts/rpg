// hud.go

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

	hudFormatter string  = "Level: %d\nXP: %d\nHP: %d"
	hudFontSize  float64 = 32
)

// The HUD contains all on-screen text, controls, and buttons.
type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Update redraws the HUD to use current values.
func (h *HUD) Update() {
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(h.BasicEntity)
			h.RenderComponent.Drawable = common.Text{
				Font: gameFontHUD,
				Text: fmt.Sprintf(hudFormatter, gamePlayer.GetLevel(),
					gamePlayer.GetExperience(), gamePlayer.GetHitPoints()),
			}
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
}

// initializeHUDFont prepares the font used in the HUD.
func initializeHUDFont() error {
	if gameFontHUD == nil {
		gameFontHUD = &common.Font{
			URL:  "fonts/hud.ttf",
			BG:   color.Black,
			FG:   color.White,
			Size: hudFontSize,
		}
	}
	if err := gameFontHUD.CreatePreloaded(); err != nil {
		return err
	}
	return nil
}

// newHUD configures and returns a HUD system.
func newHUD() (*HUD, error) {
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: gameFontHUD,
		Text: fmt.Sprintf(hudFormatter, gamePlayer.GetLevel(),
			gamePlayer.GetExperience(), gamePlayer.GetHitPoints()),
	}
	h.RenderComponent.SetZIndex(zHUD)
	h.SetShader(common.HUDShader)
	h.SpaceComponent = common.SpaceComponent{Position: engo.Point{0, 0}}
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
	return h, nil
}
