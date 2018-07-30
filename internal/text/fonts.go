package text

import (
	"image/color"

	"engo.io/engo"
	"engo.io/engo/common"
)

type FontSet struct {
	TitleFont *common.Font
	HUDFont   *common.Font
	LogFont   *common.Font
}

func PreloadFont(path string, size float64) (*common.Font, error) {
	if err := engo.Files.Load(path); err != nil {
		return nil, err
	}

	f := &common.Font{
		URL:  path,
		FG:   color.White,
		Size: size,
	}

	if err := f.CreatePreloaded(); err != nil {
		return nil, err
	}

	return f, nil
}
