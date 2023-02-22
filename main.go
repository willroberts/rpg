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

	"github.com/EngoEngine/engo"
	"github.com/willroberts/rpg/rpg"
)

const (
	width  int = 720
	height int = 720
)

// Main starts the game.
func main() {
	o := engo.RunOptions{
		Title:  "RPG",
		Width:  width,
		Height: height,
	}
	log.Printf("starting game at video resolution %dx%d\n", o.Width, o.Height)
	engo.Run(o, &rpg.GameScene{})
}
