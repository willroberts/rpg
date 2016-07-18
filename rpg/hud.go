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
	CombatLogPosX int = 16  // 16px from the left
	CombatLogPosY int = 548 // 16px from the bottom
)

var (
	GameHUD       *HUD
	HUDFont       *common.Font
	CombatLogFont *common.Font
)

// CombatLog stores a rolling set of five combat messages to be displayed on the
// screen.
type CombatLog struct {
	Log [5]*CombatMessage
}

// Update rotates the CombatLog and then updates the last entry.
// TODO: Redraw this on the screen.
func (l *CombatLog) Update(m string) {
	for i := 4; i > 0; i-- {
		l.Log[i-1] = l.Log[i]
	}
	l.Log[4] = newCombatMessage(m, 96)
}

// CombatMessage is a log of an event from the combat system.
type CombatMessage struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Text string
}

// The HUD contains all on-screen text, controls, and buttons.
type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// UpdateHealth redraws the HP text on the screen when PLayer HP changes.
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
func initializeFonts() error {
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

	if CombatLogFont == nil {
		CombatLogFont = &common.Font{
			URL:  "fonts/combatlog.ttf",
			BG:   color.Black,
			FG:   color.White,
			Size: 24,
		}
	}
	if err := CombatLogFont.CreatePreloaded(); err != nil {
		return err
	}

	return nil
}

// newCombatLog creates a rotating log window on the screen.
func newCombatLog() error {
	l := &CombatLog{}
	offset := 0
	for i := 0; i < 5; i++ {
		l.Log[i] = newCombatMessage(fmt.Sprintf("message %d", i), offset)
		offset += 32
	}
	return nil
}

// newCombatMessage creates one of the message rows in the CombatLog.
func newCombatMessage(msg string, offset int) *CombatMessage {
	m := &CombatMessage{Text: msg}
	m.RenderComponent.Drawable = common.Text{Font: CombatLogFont, Text: m.Text}
	m.SetShader(common.HUDShader)
	m.SpaceComponent = common.SpaceComponent{Position: engo.Point{
		float32(CombatLogPosX),
		float32(CombatLogPosY + offset),
	}}
	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&m.BasicEntity, &m.RenderComponent, &m.SpaceComponent)
		}
	}
	return m
}

// newHUD configures and returns a HUD system.
func newHUD() (*HUD, error) {
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: HUDFont,
		Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
	}
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
