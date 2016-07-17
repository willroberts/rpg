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

	X, Y int

	Hostility string
	HitPoints int
}

func NewPlayer(x, y, spriteIndex int) *Player {
	p := &Player{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
		X:         x,
		Y:         y,
		Hostility: "neutral",
		HitPoints: 10, // New characters start with 10 HP for now.
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

	// Add to grid.
	grid.AddCharacter(p, x, y)

	return p
}

// TODO: Prevent movement when adjacent grid contains an enemy
func movePlayer(e controlEntity) {
	// Handle keypresses.
	var moveDirection string
	if engo.Input.Button("moveleft").JustPressed() {
		moveDirection = "left"
	} else if engo.Input.Button("moveright").JustPressed() {
		moveDirection = "right"
	} else if engo.Input.Button("moveup").JustPressed() {
		moveDirection = "up"
	} else if engo.Input.Button("movedown").JustPressed() {
		moveDirection = "down"
	}

	// Don't process empty keypresses.
	// FIXME: Move keypress detection back to the scene, before this is called.
	if moveDirection == "" {
		return
	}

	switch moveDirection {
	case "left":
		if player.GetX() == grid.MinX {
			log.Println("You can't go that way!")
			return
		} else {
			grid.MoveCharacter(player, player.GetX()-1, player.GetY())
		}
	case "right":
		if player.GetX() == grid.MaxX {
			log.Println("You can't go that way!")
			return
		} else {
			grid.MoveCharacter(player, player.GetX()+1, player.GetY())
		}
	case "up":
		if player.GetY() == grid.MinY {
			log.Println("You can't go that way!")
			return
		} else {
			grid.MoveCharacter(player, player.GetX(), player.GetY()-1)
		}
	case "down":
		if player.GetY() == grid.MaxY {
			log.Println("You can't go that way!")
			return
		} else {
			grid.MoveCharacter(player, player.GetX(), player.GetY()+1)
		}
	}

	// Update the player's space component for redrawing if necessary.
	e.SpaceComponent.Position.X = (float32(player.GetX()) * characterSizeX) + characterOffsetX
	e.SpaceComponent.Position.Y = (float32(player.GetY()) * characterSizeY) + characterOffsetY
}

// Satisfy the Character interface.
func (p *Player) GetX() int             { return p.X }
func (p *Player) GetY() int             { return p.Y }
func (p *Player) SetX(x int)            { p.X = x }
func (p *Player) SetY(y int)            { p.Y = y }
func (p *Player) GetHostility() string  { return p.Hostility }
func (p *Player) SetHostility(h string) { p.Hostility = h }
func (p *Player) GetHitPoints() int     { return p.HitPoints }

func (p *Player) ModifyHitPoints(amount int) {
	p.HitPoints += amount
	if p.HitPoints <= 0 {
		p.Destroy()
	}
}

func (p *Player) Destroy() {
	log.Println("You died!")
}
