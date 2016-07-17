// player.go
package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	player *Player
)

type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	ControlComponent

	Type      string
	Hostility string
	HitPoints int

	X, Y int
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
		Type:      "player",
		Hostility: "neutral",
		HitPoints: 10, // New characters start with 10 HP for now.
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
	positionX := float32(player.GetX()) * characterSizeX
	positionY := float32(player.GetY()) * characterSizeY
	// The gravestone is a differently-sized asset which doesn't need an offset.
	if player.HitPoints > 0 {
		positionX += characterOffsetX
		positionY += characterOffsetY
	}
	e.SpaceComponent.Position.X = positionX
	e.SpaceComponent.Position.Y = positionY
}

func (p *Player) GetX() int             { return p.X }
func (p *Player) GetY() int             { return p.Y }
func (p *Player) SetX(x int)            { p.X = x }
func (p *Player) SetY(y int)            { p.Y = y }
func (p *Player) GetType() string       { return p.Type }
func (p *Player) GetHostility() string  { return p.Hostility }
func (p *Player) SetHostility(h string) { p.Hostility = h }
func (p *Player) GetHitPoints() int     { return p.HitPoints }
func (p *Player) GetDamage() int        { return 1 }

func (p *Player) ModifyHitPoints(amount int) {
	p.HitPoints += amount
}

// When the player dies, we replace the character texture and stop processing
// incoming movement commands.
func (p *Player) Destroy() {
	for _, system := range GameWorld.Systems() {
		switch sys := system.(type) {
		case *ControlSystem:
			// 1. Remove the Player from the ControlSystem
			sys.Remove(p.BasicEntity)
		case *common.RenderSystem:
			// 2. Remove the Player from the RenderSystem
			sys.Remove(p.BasicEntity)
			// 3. Add the Gravestone to the RenderSystem
			playerTexture := decorationSpritesheet.Cell(spriteGravestone)
			p.RenderComponent = common.RenderComponent{
				Drawable: playerTexture,
				Scale:    engo.Point{2, 2},
			}
			p.RenderComponent.SetZIndex(1)
			sys.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		}
	}
}
