// character.go

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

// A Character is a generic entity which can occupy space on a tile, including
// player characters and non-player characters.
type Character interface {
	GetX() int
	GetY() int
	SetX(int)
	SetY(int)

	GetName() string
	GetHostility() string
	SetHostility(string)

	GetHitPoints() int
	ModifyHitPoints(int)
	GetDamage() int
	GetXPBonus() int

	Destroy()
}
