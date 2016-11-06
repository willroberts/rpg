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
