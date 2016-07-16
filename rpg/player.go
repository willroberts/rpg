// player.go
package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	player         *Player
	playerEntityID uint64
)

type Player struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
	ControlComponent

	HitPoints int
	X, Y      int
}

func NewPlayer(x, y, spriteIndex int) *Player {
	p := &Player{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
		HitPoints: 10,
		X:         x,
		Y:         y,
	}
	playerEntityID = p.BasicEntity.ID()

	// Configure collision.
	p.CollisionComponent = common.CollisionComponent{
		Solid: true,
		Main:  true,
	}

	// Add graphics.
	playerTexture := characterSpritesheet.Cell(spriteIndex)
	p.RenderComponent = common.RenderComponent{
		Drawable: playerTexture,
		Scale:    engo.Point{2, 2},
	}
	p.RenderComponent.SetZIndex(1)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			(characterSizeX * float32(x)) + characterOffsetX,
			(characterSizeY * float32(y)) + characterOffsetY,
		},
		Width:  characterSizeX,
		Height: characterSizeY,
	}

	return p
}

// TODO: Prevent movement when adjacent grid contains an enemy
func movePlayer(e controlEntity) {
	// Handle keypresses.
	var moveDirection string
	if engo.Input.Button("moveleft").JustPressed() {
		moveDirection = "left"
		if player.X == grid.MinX {
			log.Println("You can't go that way!")
			return
		} else {
			player.X -= 1
		}
	} else if engo.Input.Button("moveright").JustPressed() {
		moveDirection = "right"
		if player.X == grid.MaxX {
			log.Println("You can't go that way!")
			return
		} else {
			player.X += 1
		}
	} else if engo.Input.Button("moveup").JustPressed() {
		moveDirection = "up"
		if player.Y == grid.MinY {
			log.Println("You can't go that way!")
			return
		} else {
			player.Y -= 1
		}
	} else if engo.Input.Button("movedown").JustPressed() {
		moveDirection = "down"
		if player.Y == grid.MaxY {
			log.Println("You can't go that way!")
			return
		} else {
			player.Y += 1
		}
	}

	// Don't process empty keypresses.
	// FIXME: Move keypress detection back to the scene, before this is called.
	if moveDirection == "" {
		return
	}

	// Update the player's space component for redrawing if necessary.
	e.SpaceComponent.Position.X = (float32(player.X) * characterSizeX) + characterOffsetX
	e.SpaceComponent.Position.Y = (float32(player.Y) * characterSizeY) + characterOffsetY
}

// Satisfy the Character interface.
func (p *Player) GetX() int { return p.X }
func (p *Player) GetY() int { return p.Y }
