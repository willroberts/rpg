// menu.go
// The game's main menu, including character creation.

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
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	menuFontSize  float64 = 32
	titleFontSize float64 = 64
)

type MenuScene struct {
	Width     int
	Height    int
	TitleFont *common.Font
	MenuFont  *common.Font
}

func (s *MenuScene) Preload() {
	log.Println("preloading fonts")
	if err := engo.Files.Load("fonts/title.ttf"); err != nil {
		panic(err)
	}
	s.TitleFont = &common.Font{
		URL:  "fonts/title.ttf",
		FG:   color.White,
		Size: titleFontSize,
	}
	if err := s.TitleFont.CreatePreloaded(); err != nil {
		panic(err)
	}
	if err := engo.Files.Load("fonts/menu.ttf"); err != nil {
		panic(err)
	}
	s.MenuFont = &common.Font{
		URL:  "fonts/menu.ttf",
		FG:   color.White,
		Size: menuFontSize,
	}
	if err := s.MenuFont.CreatePreloaded(); err != nil {
		panic(err)
	}
}

func (s *MenuScene) Setup(w *ecs.World) {
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})

	// draw some kind of non-black background

	// draw the game title at the top
	t := &Label{BasicEntity: ecs.NewBasic()}
	t.RenderComponent.Drawable = common.Text{
		Font: s.TitleFont,
		Text: "Game Title",
	}
	t.SetShader(common.HUDShader)
	// FIXME: Determine how wide the rendered text is (t.SpaceComponent.Width?)
	t.SpaceComponent = common.SpaceComponent{Position: engo.Point{200, 64}}

	// draw "Character Creation"
	l := &Label{BasicEntity: ecs.NewBasic()}
	l.RenderComponent.Drawable = common.Text{
		Font: s.MenuFont,
		Text: "Character Creation",
	}
	l.SetShader(common.HUDShader)
	l.SpaceComponent = common.SpaceComponent{Position: engo.Point{64, 160}}

	// draw a character image selector

	// draw a character stats selector

	// draw a "Play" button

	// render everything
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			s.Add(&l.BasicEntity, &l.RenderComponent, &l.SpaceComponent)
		}
	}
}

func (s *MenuScene) Type() string {
	return "MenuScene"
}

type Label struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
