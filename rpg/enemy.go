// enemy.go
package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Enemy struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent

	X, Y int

	Hostility string
	HitPoints int
}

func NewEnemy(x, y, spriteIndex int) *Enemy {
	e := &Enemy{
		BasicEntity: ecs.NewBasic(),
		X:           x,
		Y:           y,
		Hostility:   "hostile",
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
			(characterSizeX * float32(x)) + characterOffsetX,
			(characterSizeY * float32(y)) + characterOffsetY,
		},
		Width:  characterSizeX,
		Height: characterSizeY,
	}

	// Add to grid.
	grid.AddCharacter(e, x, y)

	return e
}

// Satisfy the Character interface.
func (e *Enemy) GetX() int             { return e.X }
func (e *Enemy) GetY() int             { return e.Y }
func (e *Enemy) SetX(x int)            { e.X = x }
func (e *Enemy) SetY(y int)            { e.Y = y }
func (e *Enemy) GetHostility() string  { return e.Hostility }
func (e *Enemy) SetHostility(h string) { e.Hostility = h }
func (e *Enemy) GetHitPoints() int     { return e.HitPoints }

func (e *Enemy) ModifyHitPoints(amount int) {
	e.HitPoints += amount
	if e.HitPoints <= 0 {
		e.Destroy()
	}
}

func (e *Enemy) Destroy() {
	log.Println("Enemy was killed!")
}
