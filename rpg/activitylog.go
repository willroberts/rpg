// activitylog.go

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

const (
	logWindowX      float32 = 8
	logWindowY      float32 = 570
	logWindowWidth  float32 = 400
	logWindowHeight float32 = 138

	logMessageX        int     = 16
	logMessageY        int     = 580
	logMessageCount    int     = 5
	logMessageSpacing  int     = 25
	logMessageFontSize float64 = 24
)

// ActivityLog stores a rolling set of five messages to be displayed on the screen.
type ActivityLog struct {
	Log [logMessageCount]*ActivityMessage
}

// Update rotates the ActivityLog and then updates the last entry.
func (l *ActivityLog) Update(m string) {
	// Shift all messages up one line.
	for i := 0; i < logMessageCount-1; i++ {
		l.Log[i].Text = l.Log[i+1].Text
	}

	// Place the new message at the bottom.
	l.Log[logMessageCount-1].Text = m

	// Redraw all messages in the log.
	for i := 0; i < logMessageCount; i++ {
		l.Log[i].Draw()
	}
}

// ActivityLogWindow is the HUD element under the messages.
type ActivityLogWindow struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	PosX, PosY    float32
	Width, Height float32
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
	m.RenderComponent.Drawable = common.Text{Font: gameFontLog, Text: m.Text}
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&m.BasicEntity, &m.RenderComponent, &m.SpaceComponent)
		}
	}
}

// initializeActivityMessage creates one of the message rows in the ActivityLog.
func initializeActivityMessage(msg string, offset int) *ActivityMessage {
	m := &ActivityMessage{Text: msg}
	m.RenderComponent.Drawable = common.Text{Font: gameFontLog, Text: m.Text}
	m.RenderComponent.SetZIndex(zText)
	m.SetShader(common.HUDShader)
	m.SpaceComponent = common.SpaceComponent{Position: engo.Point{
		float32(logMessageX),
		float32(logMessageY + offset),
	}}
	m.Draw()
	return m
}

// initializeLogFont prepares the font used in the activity log.
func initializeLogFont() error {
	if gameFontLog == nil {
		gameFontLog = &common.Font{
			URL:  "fonts/combatlog.ttf",
			BG:   color.Black,
			FG:   color.White,
			Size: logMessageFontSize,
		}
	}
	if err := gameFontLog.CreatePreloaded(); err != nil {
		return err
	}
	return nil
}

// newActivityLog creates a rotating log window on the screen.
// FIXME: Player can still walk under the activity log.
func newActivityLog() *ActivityLog {
	// Draw background rectangle.
	w := &ActivityLogWindow{
		BasicEntity: ecs.NewBasic(),
		PosX:        logWindowX,
		PosY:        logWindowY,
		Width:       logWindowWidth,
		Height:      logWindowHeight,
	}
	w.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.Black,
	}
	w.RenderComponent.SetZIndex(zHUD)
	w.SetShader(common.HUDShader)
	w.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{w.PosX, w.PosY},
		Width:    w.Width,
		Height:   w.Height,
	}
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&w.BasicEntity, &w.RenderComponent, &w.SpaceComponent)
		}
	}

	// Draw log messages.
	l := &ActivityLog{}
	var offset int
	for i := 0; i < logMessageCount; i++ {
		l.Log[i] = initializeActivityMessage("", offset)
		offset += logMessageSpacing
	}
	return l
}
