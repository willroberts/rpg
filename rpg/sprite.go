package rpg

import (
	"log"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
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

	characterSpritesheetPath  string = "spritesheets/characters-32x32.png"
	decorationSpritesheetPath string = "spritesheets/decoration-20x20-40x40.png"
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

// PreloadSprites preloads all known sprites and returns them in a SpriteSet.
func PreloadSprites() (*SpriteSet, error) {
	if gameScene.Sprites != nil {
		log.Println("sprites already loaded, reusing")
		return gameScene.Sprites, nil
	}
	ss := &SpriteSet{}

	cs, err := LoadSprites(characterSpritesheetPath, 32, 32)
	if err != nil {
		return &SpriteSet{}, err
	}
	ss.Characters = cs

	ds, err := LoadSprites(decorationSpritesheetPath, 40, 40)
	if err != nil {
		return &SpriteSet{}, err
	}
	ss.Decorations = ds

	return ss, nil
}
