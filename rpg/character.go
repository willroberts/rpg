package rpg

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	characterSpritesheetPath string = "spritesheets/characters-32x32.png"

	startingX float32 = 86
	startingY float32 = 84
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

// FIXME: Make generic for other characters.
func NewCharacter() Character {
	c := Character{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
	}

	// Add graphics
	// FIXME: Move spritesheet logic to a place where it can be shared by enemies.
	spritesheet := common.NewSpritesheetFromFile(characterSpritesheetPath, 32, 32)
	characterTexture := spritesheet.Cell(spriteWhiteZombie)
	c.RenderComponent = common.RenderComponent{
		Drawable: characterTexture,
		Scale:    engo.Point{2, 2},
	}
	c.RenderComponent.SetZIndex(1)
	c.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{startingX, startingY},
		Width:    80,
		Height:   80,
	}

	return c
}
