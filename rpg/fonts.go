// fonts.go
// Handles the loading and storing of fonts.

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

	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	// Font used at the top of the screen during gameplay.
	HUDFontPath string = "fonts/hud.ttf"
	// Font used in the ActivityLog window.
	LogFontPath string = "fonts/combatlog.ttf"
	// Font used in the Main Menu and Character Select.
	MenuFontPath string = "fonts/menu.ttf"
	// Font used on the Title Screen.
	TitleFontPath string = "fonts/title.ttf"
)

// A FontSet contains all named fonts we use in the game.
type FontSet struct {
	HUDFont   *common.Font
	LogFont   *common.Font
	MenuFont  *common.Font
	TitleFont *common.Font
}

// LoadFont preloads a single font and returns it.
func LoadFont(font string, size float64) (*common.Font, error) {
	if err := engo.Files.Load(font); err != nil {
		return &common.Font{}, err
	}
	f := &common.Font{
		URL:  font,
		FG:   color.White,
		Size: size,
	}
	if err := f.CreatePreloaded(); err != nil {
		return &common.Font{}, err
	}
	return f, nil
}

// PreloadFonts preloads all known fonts and returns them in a FontSet.
func PreloadFonts() (*FontSet, error) {
	if gameFonts != nil {
		log.Println("fonts already loaded, reusing")
		return gameFonts, nil
	}

	log.Println("preloading fonts")
	fs := &FontSet{}

	hf, err := LoadFont(HUDFontPath, 32)
	if err != nil {
		return &FontSet{}, err
	}
	fs.HUDFont = hf
	// FIXME: Reusing fonts here since they're identical.
	fs.MenuFont = hf

	lf, err := LoadFont(LogFontPath, 24)
	if err != nil {
		return &FontSet{}, err
	}
	fs.LogFont = lf

	tf, err := LoadFont(TitleFontPath, 64)
	if err != nil {
		return &FontSet{}, err
	}
	fs.TitleFont = tf

	return fs, nil
}
