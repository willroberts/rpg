// main.go

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

// Package main starts the game.
package main

import (
	"log"

	"engo.io/engo"
	"github.com/willroberts/rpg/scenes"
)

// basic scene: 480x480
// game scene: 720x720
const (
	width  int = 480
	height int = 480
)

// Main starts the game.
func main() {
	o := engo.RunOptions{
		Title:  "RPG",
		Width:  width,
		Height: height,
	}
	log.Printf("starting game at video resolution %dx%d\n", width, height)
	engo.Run(o, &scenes.BasicScene{})
}
