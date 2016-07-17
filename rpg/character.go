// character.go
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
	// Graphics
	characterSpritesheetPath   string = "spritesheets/characters-32x32.png"
	characterSpritesheetWidth  int    = 32
	characterSpritesheetHeight int    = 32

	decorationSpritesheetPath   string = "spritesheets/decoration-20x20-40x40.png"
	decorationSpritesheetWidth  int    = 40
	decorationSpritesheetHeight int    = 40

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
)

var (
	characterSpritesheet  *common.Spritesheet
	decorationSpritesheet *common.Spritesheet
)

// To satisfy this interface, use these methods to return c.X and c.Y.
// Characters can be referenced in the grid.
type Character interface {
	GetX() int
	GetY() int
	SetX(int)
	SetY(int)

	GetType() string
	GetHostility() string
	SetHostility(string)

	GetHitPoints() int
	ModifyHitPoints(int)
	GetDamage() int

	Destroy()
}
