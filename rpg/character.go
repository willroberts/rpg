package rpg

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	characterSpritesheetPath   string = "spritesheets/characters-32x32.png"
	characterSpritesheetWidth  int    = 32
	characterSpritesheetHeight int    = 32

	characterSizeX     float32 = 80
	characterSizeY     float32 = 80
	characterOffsetX   float32 = 6
	characterOffsetY   float32 = 4
	characterStartingX float32 = characterSizeX + characterOffsetX
	characterStartingY float32 = characterSizeY + characterOffsetY
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

type Character struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	ControlComponent
}

func NewCharacter(spriteIndex int) Character {
	c := Character{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
	}

	// Add graphics
	// FIXME: Move spritesheet logic to a place where it can be shared by enemies.
	spritesheet := common.NewSpritesheetFromFile(characterSpritesheetPath,
		characterSpritesheetWidth, characterSpritesheetHeight)
	characterTexture := spritesheet.Cell(spriteIndex)
	c.RenderComponent = common.RenderComponent{
		Drawable: characterTexture,
		Scale:    engo.Point{2, 2},
	}
	c.RenderComponent.SetZIndex(1)
	c.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{characterStartingX, characterStartingY},
		Width:    characterSizeX,
		Height:   characterSizeY,
	}

	return c
}
