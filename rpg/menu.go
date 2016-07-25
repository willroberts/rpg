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
	"engo.io/engo/common"
)

type MenuScene struct {
	Width  int
	Height int
}

func (s *MenuScene) Preload() {
	log.Println("preloading menu fonts")
	var err error
	gameFonts, err = PreloadFonts()
	if err != nil {
		panic(err)
	}
	log.Println("preloading menu sprites")
	gameSprites, err = PreloadSprites()
	if err != nil {
		panic(err)
	}
}

func (s *MenuScene) Setup(w *ecs.World) {
	common.SetBackground(color.RGBA{40, 44, 40, 255})
	w.AddSystem(&common.RenderSystem{})

	// Create title label and "Portrait:" label.
	// FIXME: Determine width automatically - via t.SpaceComponent.Width?
	tl := NewLabel("Game Title", gameFonts.TitleFont, 200, 64)
	pl := NewLabel("Portrait:", gameFonts.MenuFont, 32, 160)

	// draw a character image selector
	csp := NewCharSelectPanel()

	// draw a character stats selector

	// draw a "Play" button

	// render everything
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&tl.BasicEntity, &tl.RenderComponent, &tl.SpaceComponent)
			s.Add(&pl.BasicEntity, &pl.RenderComponent, &pl.SpaceComponent)
			s.Add(&csp.BasicEntity, &csp.RenderComponent, &csp.SpaceComponent)
		}
	}
}

func (s *MenuScene) Type() string {
	return "MenuScene"
}
