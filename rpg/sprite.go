// sprite.go
// Handles the loading and storing of sprites.

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
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	// Indices in CharacterSpritesheet.
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

	// Indices in DecorationSpritesheet.
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

	CharacterSpritesheetPath  string = "spritesheets/characters-32x32.png"
	DecorationSpritesheetPath string = "spritesheets/decoration-20x20-40x40.png"

	// Spritesheets
	charSpritesheetWidth  int = 32
	charSpritesheetHeight int = 32
	decoSpritesheetWidth  int = 40
	decoSpritesheetHeight int = 40

	// SpaceComponent dimensions of characters.
	charScale float32 = 2.0 // 32 -> 64px
	charSizeX float32 = 80
	charSizeY float32 = 80

	// Character art is 64x64 when scaled, and needs to be slightly offset in order
	// to be centered in an 80x80 tile.
	charOffsetX float32 = 8
	charOffsetY float32 = 4
)

// A SpriteSet contains all sprites we use in the game.
type SpriteSet struct {
	Characters  *common.Spritesheet
	Decorations *common.Spritesheet
}

// LoadSprites preloads a single spritesheet and returns it.
func LoadSprites(path string, w, h int) (*common.Spritesheet, error) {
	if err := engo.Files.Load(path); err != nil {
		return &common.Spritesheet{}, err
	}
	ss := common.NewSpritesheetFromFile(path, w, h)
	return ss, nil
}

// preloadSprites reads sprites from files on disk.
func preloadSprites() error {
	var err error
	gameSpritesChar, err = LoadSprites(CharacterSpritesheetPath,
		charSpritesheetWidth, charSpritesheetHeight)
	if err != nil {
		return err
	}
	gameSpritesDeco, err = LoadSprites(DecorationSpritesheetPath,
		decoSpritesheetWidth, decoSpritesheetHeight)
	if err != nil {
		return err
	}
	return nil
}
