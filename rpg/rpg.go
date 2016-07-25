// rpg.go

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

// Package rpg contains all systems needed to run our 2D RPG.
package rpg

import "engo.io/ecs"

// Create some shared variables in the package scope.
// All shared variables should start with "game".
// FIXME: Do all of these need package scope?
// TODO: Use message passing between systems instead of sharing data.
var (
	gameWorld           *ecs.World
	gameFonts           *FontSet
	gameSprites         *SpriteSet
	gameGrid            *Grid
	gameHUD             *HUD
	gameLog             *ActivityLog
	gamePlayer          *Player
	gameEnemyTypes      = make(map[string]EnemyAttributes)
	gameExperienceTable = make(map[string]int)
)
