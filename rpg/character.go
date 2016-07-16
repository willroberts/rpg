package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Terminology:
// Character: Anything which can move, interact, attack, etc.
// Player: The player-controlled character
// NPC: Non-hostile characters like shopkeepers
// Enemy: Hostile NPCs

// Variables and constants named "character" should be applicable to both the
// player and enemies/NPCs.

const (
	// Graphics
	characterSpritesheetPath   string = "spritesheets/characters-32x32.png"
	characterSpritesheetWidth  int    = 32
	characterSpritesheetHeight int    = 32

	// Graphics coordinates
	characterSizeX   float32 = 80
	characterSizeY   float32 = 80
	characterOffsetX float32 = 8
	characterOffsetY float32 = 4
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
	player               *Player
	playerEntityID       uint64
	characterSpritesheet *common.Spritesheet
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
	c := &Player{
		BasicEntity: ecs.NewBasic(),
		ControlComponent: ControlComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
		HitPoints: 10,
		X:         x,
		Y:         y,
	}
	playerEntityID = c.BasicEntity.ID()

	// Configure collision.
	c.CollisionComponent = common.CollisionComponent{
		Solid: true,
		Main:  true,
	}

	// Add graphics.
	playerTexture := characterSpritesheet.Cell(spriteIndex)
	c.RenderComponent = common.RenderComponent{
		Drawable: playerTexture,
		Scale:    engo.Point{2, 2},
	}
	c.RenderComponent.SetZIndex(1)
	c.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			(characterSizeX * float32(x)) + characterOffsetX,
			(characterSizeY * float32(y)) + characterOffsetY,
		},
		Width:  characterSizeX,
		Height: characterSizeY,
	}

	return c
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
