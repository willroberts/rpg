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
	characterOffsetX   float32 = 8
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

var (
	characterSpritesheet *common.Spritesheet
	characterEntityID    uint64
)

type Character struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
	ControlComponent

	HitPoints int
}

func NewCharacter(spriteIndex int) Character {
	c := Character{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
		HitPoints: 10,
	}
	characterEntityID = c.BasicEntity.ID()

	// Configure collision.
	c.CollisionComponent = common.CollisionComponent{
		Solid: true,
		Main:  true,
	}

	// Add graphics.
	characterTexture := characterSpritesheet.Cell(spriteIndex)
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

func moveCharacter(e controlEntity) {
	// Have the arrow keys move the character one tile at a time.
	if engo.Input.Button("moveleft").JustPressed() {
		e.SpaceComponent.Position.X -= characterSizeX
	}
	if engo.Input.Button("moveright").JustPressed() {
		e.SpaceComponent.Position.X += characterSizeX
	}
	if engo.Input.Button("moveup").JustPressed() {
		e.SpaceComponent.Position.Y -= characterSizeY
	}
	if engo.Input.Button("movedown").JustPressed() {
		e.SpaceComponent.Position.Y += characterSizeY
	}

	// Prevent the character from leaving the map.
	if e.SpaceComponent.Position.X < 0 {
		e.SpaceComponent.Position.X = characterOffsetX
	} else if e.SpaceComponent.Position.X > widthLimit {
		e.SpaceComponent.Position.X = widthLimit + characterOffsetX
	}
	if e.SpaceComponent.Position.Y < 0 {
		e.SpaceComponent.Position.Y = characterOffsetY
	} else if e.SpaceComponent.Position.Y > heightLimit {
		e.SpaceComponent.Position.Y = heightLimit + characterOffsetY
	}
}
