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
	labels := []*Label{
		NewLabel("Create Character", gameFonts.TitleFont, 128, 50),
		NewLabel("Name:", gameFonts.MenuFont, 40, 150),
		NewLabel("Portrait:", gameFonts.MenuFont, 40, 230),
		NewLabel("Attributes:", gameFonts.MenuFont, 40, 380),
		NewLabel("STR: XX", gameFonts.MenuFont, 180, 460),
		NewLabel("DEX: XX", gameFonts.MenuFont, 180, 540),
		NewLabel("INT: XX", gameFonts.MenuFont, 400, 460),
		NewLabel("VIT: XX", gameFonts.MenuFont, 400, 540),
		NewLabel("Start", gameFonts.TitleFont, 280, 620),
	}

	// draw a character image selector
	csp := NewCharSelectPanel()

	// draw a character stats selector

	// render everything
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			for _, l := range labels {
				s.Add(&l.BasicEntity, &l.RenderComponent, &l.SpaceComponent)
			}
			s.Add(&csp.BasicEntity, &csp.RenderComponent, &csp.SpaceComponent)
			for _, p := range csp.Portraits {
				s.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
			}
		}
	}
}

func (s *MenuScene) Type() string {
	return "MenuScene"
}
