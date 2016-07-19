// menu.go

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
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA

package rpg

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// MenuScene displays the main menu.
type MenuScene struct{}

// Preload validates and loads assets. In can cause panics, since the game cannot
// run without its assets.
func (scene *MenuScene) Preload() {
	log.Println("preloading fonts")
	if err := engo.Files.Load("fonts/hud.ttf"); err != nil {
		panic(err)
	}
	if err := engo.Files.Load("fonts/combatlog.ttf"); err != nil {
		panic(err)
	}
}

// Setup initializes all systems necessary for the game to function.
func (scene *MenuScene) Setup(w *ecs.World) {
	log.Println("creating scene")
	gameWorld = w
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&ControlSystem{})

	log.Println("configuring systems")
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			log.Println("configuring render system")
		case *ControlSystem:
			log.Println("configuring control system")
		}
	}

	log.Println("configuring camera")

	log.Println("creating hud")
	initializeHUDFont()
	initializeLogFont()

	log.Println("binding controls")
	bindControls()
	log.Println("controls bound")
}

// Type returns the name of the scene. This is used to satisfy engo's Scene
// interface.
func (scene *MenuScene) Type() string {
	return "MenuScene"
}
