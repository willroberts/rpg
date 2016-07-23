// input.go

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

import "engo.io/engo"

// bindControls assigns the arrow keys to Player movement.
func bindControls() {
	engo.Input.RegisterButton("moveup", engo.ArrowUp)
	engo.Input.RegisterButton("movedown", engo.ArrowDown)
	engo.Input.RegisterButton("moveleft", engo.ArrowLeft)
	engo.Input.RegisterButton("moveright", engo.ArrowRight)
}
