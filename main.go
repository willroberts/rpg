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
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA

// Package main starts the game.
package main

import (
	"log"

	"engo.io/engo"
	"github.com/willroberts/rpg/rpg"
)

const (
	width  int = 960
	height int = 880
)

// Main starts the game.
func main() {
	o := engo.RunOptions{
		Title:  "RPG",
		Width:  width,
		Height: height,
	}
	log.Printf("starting game. video resolution: %dx%d\n", width, height)
	engo.Run(o, &rpg.GameScene{})
}
