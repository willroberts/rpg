package sprite

import (
	"engo.io/engo"
	"engo.io/engo/common"
)

type SpriteSet struct {
	Characters  *common.Spritesheet
	Decorations *common.Spritesheet
}

func PreloadSpritesheet(filename string) error {
	return engo.Files.Load(filename)
}

func LoadSpritesheet(filename string, w int, h int) (*common.Spritesheet, error) {
	return common.NewSpritesheetFromFile(filename, w, h), nil
}
