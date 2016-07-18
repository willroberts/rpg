// activitylog.go

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
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	ActivityLogPosX    int = 16  // 16px from the left
	ActivityLogPosY    int = 580 // gameheight minus two tiles (160px) plus 20px buffer to center it
	ActivityLogSpacing int = 24
)

var (
	GameLog         *ActivityLog
	ActivityLogFont *common.Font
)

// ActivityLog stores a rolling set of five messages to be displayed on the screen.
type ActivityLog struct {
	Log [5]*ActivityMessage
}

// Update rotates the ActivityLog and then updates the last entry.
// FIXME: Condense this code.
func (l *ActivityLog) Update(m string) {
	r1 := l.Log[0]
	r2 := l.Log[1]
	r3 := l.Log[2]
	r4 := l.Log[3]
	r5 := l.Log[4]

	r1.Text = r2.Text
	r2.Text = r3.Text
	r3.Text = r4.Text
	r4.Text = r5.Text
	r5.Text = m

	r1.Draw()
	r2.Draw()
	r3.Draw()
	r4.Draw()
	r5.Draw()
}

// ActivityMessage is a log of an event from the game logic.
type ActivityMessage struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Text string
}

// Draw draws a log message on the screen
func (m *ActivityMessage) Draw() {
	m.RenderComponent.Drawable = common.Text{Font: ActivityLogFont, Text: m.Text}
	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&m.BasicEntity, &m.RenderComponent, &m.SpaceComponent)
		}
	}
}

// initializeActivityMessage creates one of the message rows in the ActivityLog.
func initializeActivityMessage(msg string, offset int) *ActivityMessage {
	m := &ActivityMessage{Text: msg}
	m.RenderComponent.Drawable = common.Text{Font: ActivityLogFont, Text: m.Text}
	m.RenderComponent.SetZIndex(zHUD)
	m.SetShader(common.HUDShader)
	m.SpaceComponent = common.SpaceComponent{Position: engo.Point{
		float32(ActivityLogPosX),
		float32(ActivityLogPosY + offset),
	}}
	m.Draw()
	return m
}

// initializeFonts creates the various sizes of fonts we need.
func initializeLogFont() error {
	if ActivityLogFont == nil {
		ActivityLogFont = &common.Font{
			URL:  "fonts/combatlog.ttf",
			BG:   color.Black,
			FG:   color.White,
			Size: 24,
		}
	}
	if err := ActivityLogFont.CreatePreloaded(); err != nil {
		return err
	}
	return nil
}

// newActivityLog creates a rotating log window on the screen.
func newActivityLog() *ActivityLog {
	l := &ActivityLog{}

	// Draw background rectangle.
	/*
		r := &struct{BasicEntity: ecs.NewBasic()}
		r.RenderComponent = common.RenderComponent{
			Drawable: "FOO", // FIXME
			Scale:    engo.Point{1, 1},
		}
		p.RenderComponent.SetZIndex(3)
		for _, sys := range GameWorld.Systems() {
			switch s := sys.(type) {
			case *common.RenderSystem:
				s.Add(&r.BasicEntity, &r.RenderComponent, &r.SpaceComponent)
			}
		}
	*/

	// Draw log messages.
	offset := 0
	for i := 0; i < 5; i++ {
		l.Log[i] = initializeActivityMessage("", offset)
		offset += ActivityLogSpacing
	}
	return l
}
