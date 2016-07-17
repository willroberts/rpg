// character.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts

// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA

// Terminology:
// 	Character: Anything which can move, interact, attack, etc.
// 	Player: The player-controlled character
// 	NPC: Non-hostile characters like shopkeepers
// 	Enemy: Hostile NPCs
// Variables and constants named "character" should be applicable to both the
// player and enemies/NPCs.
package rpg

import "engo.io/engo/common"

const (
	// Sprite constants are named indices in the spritesheets.
	spriteHuman = iota
	spriteGreenZombie
	spriteOliveZombie
	spriteGoblin
	spriteBear
	spriteWhiteZombie
	spriteMummy
	spriteSkeleton
	spriteDemon
	spriteGargoyle
	spriteBones           int = 3
	spriteStairsDownLeft  int = 4
	spriteAnimalBones     int = 5
	spriteTinySignpost    int = 8
	spriteGravestone      int = 9
	spriteSignpost        int = 10
	spriteStairsUpLeft    int = 12
	spriteStairsDownRight int = 13
	spriteStairsUpRight   int = 14
	spriteStairsUpLeft2   int = 15 // darker?

	// Spritesheets
	charSpritesheetPath         string = "spritesheets/characters-32x32.png"
	charSpritesheetWidth        int    = 32
	charSpritesheetHeight       int    = 32
	decorationSpritesheetPath   string = "spritesheets/decoration-20x20-40x40.png"
	decorationSpritesheetWidth  int    = 40
	decorationSpritesheetHeight int    = 40

	// SpaceComponent dimensions of characters.
	charSizeX float32 = 80
	charSizeY float32 = 80

	// Character art is 64x64 when scaled, and needs to be slightly offset in order
	// to be centered in an 80x80 tile.
	charOffsetX float32 = 8
	charOffsetY float32 = 4
)

var (
	charSpritesheet       *common.Spritesheet
	decorationSpritesheet *common.Spritesheet
)

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

	Destroy()
}
