// character.go
// Terminology:
// 	Character: Anything which can move, interact, attack, etc.
// 	Player: The player-controlled character
// 	NPC: Non-hostile characters like shopkeepers
// 	Enemy: Hostile NPCs
// Variables and constants named "character" should be applicable to both the
// player and enemies/NPCs.
package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	// Graphics
	characterSpritesheetPath   string = "spritesheets/characters-32x32.png"
	characterSpritesheetWidth  int    = 32
	characterSpritesheetHeight int    = 32

	// Graphics coordinates
	characterSizeX   float32 = 80
	characterSizeY   float32 = 80
	characterOffsetX float32 = 8
	characterOffsetY float32 = 4
)

// Sprite indices in the spritesheet.
const (
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
)

var (
	characterSpritesheet *common.Spritesheet
)
