package sprite

import (
	"engo.io/engo"
	"engo.io/engo/common"
)

func PreloadSpritesheet(filename string) error {
	return engo.Files.Load(filename)
}

func LoadSpritesheet(filename string, w int, h int) *common.Spritesheet {
	return common.NewSpritesheetFromFile(filename, w, h)
}
