package rpg

import (
	"image/color"
	"log"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	// Font used at the top of the screen during gameplay.
	hudFontPath string = "fonts/hud.ttf"
	// Font used in the ActivityLog window.
	logFontPath string = "fonts/combatlog.ttf"
	// Font used on the Title Screen.
	titleFontPath string = "fonts/title.ttf"
)

// A FontSet contains all named fonts we use in the game.
type FontSet struct {
	HUDFont   *common.Font
	LogFont   *common.Font
	TitleFont *common.Font
}

// LoadFont preloads a single font and returns it.
func LoadFont(path string, size float64) (*common.Font, error) {
	if err := engo.Files.Load(path); err != nil {
		return &common.Font{}, err
	}
	f := &common.Font{
		URL:  path,
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
	if gameScene.Fonts != nil {
		log.Println("fonts already loaded, reusing")
		return gameScene.Fonts, nil
	}
	fs := &FontSet{}

	hf, err := LoadFont(hudFontPath, 32)
	if err != nil {
		return &FontSet{}, err
	}
	fs.HUDFont = hf

	lf, err := LoadFont(logFontPath, 24)
	if err != nil {
		return &FontSet{}, err
	}
	fs.LogFont = lf

	tf, err := LoadFont(titleFontPath, 64)
	if err != nil {
		return &FontSet{}, err
	}
	fs.TitleFont = tf

	return fs, nil
}
