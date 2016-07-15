package rpg

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Enemy struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent

	HitPoints int
}

func NewEnemy(spriteIndex int, startingX, startingY float32) Enemy {
	e := Enemy{
		BasicEntity: ecs.NewBasic(),
		HitPoints:   5,
	}

	// Configure collision.
	e.CollisionComponent = common.CollisionComponent{
		Solid: true,
	}

	// Add graphics.
	enemyTexture := characterSpritesheet.Cell(spriteIndex)
	e.RenderComponent = common.RenderComponent{
		Drawable: enemyTexture,
		Scale:    engo.Point{2, 2},
	}
	e.RenderComponent.SetZIndex(1)
	e.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			startingX + characterOffsetX,
			startingY + characterOffsetY,
		},
		Width:  characterSizeX,
		Height: characterSizeY,
	}

	return e
}
