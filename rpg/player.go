// player.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts

// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var player *Player

// Player is the player-controlled entity in the game.
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

// When the player dies, we replace the character texture and stop processing
// incoming movement commands. Removes the Player from the ControlSystem and the
// RenderSystem, and then re-add the Player to the RenderSystem as a gravestone.
func (p *Player) Destroy() {
	for _, sys := range GameWorld.Systems() {
		switch s := sys.(type) {
		case *ControlSystem:
			s.Remove(p.BasicEntity)
		case *common.RenderSystem:
			s.Remove(p.BasicEntity)
			p.RenderComponent = common.RenderComponent{
				Drawable: decorationSpritesheet.Cell(spriteGravestone),
				Scale:    engo.Point{2, 2},
			}
			p.RenderComponent.SetZIndex(1)
			s.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		}
	}
}

// GetDamage returns the damage dealt by the Player.
func (p *Player) GetDamage() int { return 1 }

// GetHitPoints returns the current HP for the Player.
func (p *Player) GetHitPoints() int { return p.HitPoints }

// GetHostility returns the demeanor of an enemy or NPC. It's included here just
// to statisfy the Character interface.
// FIXME: See if we can remove this from the Character interface.
func (p *Player) GetHostility() string { return p.Hostility }

// GetType returns the name of the Player type. It's included here just to
// satisfy the Character interface. It is also used temporarily to provide a name
// for the combat log. Eventually we'll want to use the Player's chosen name in
// the combat log.
// FIXME: See if we can remove this from the Character interface.
func (p *Player) GetType() string { return p.Type }

// GetX returns the Player's X coordinate.
func (p *Player) GetX() int { return p.X }

// GetY returns the Player's Y coordinate.
func (p *Player) GetY() int { return p.Y }

// ModifyHitPoints adds a number to the Player's hit points. To deal damage to
// the Player, provide a negative number.
func (p *Player) ModifyHitPoints(amount int) {
	p.HitPoints += amount
}

// SetHostility changes the Player's demeanor. It's included here just to satisfy
// the Character interface.
// FIXME: See if we can remove this from the Character interface.
func (p *Player) SetHostility(h string) { p.Hostility = h }

// SetX updates the Player's X coordinate.
func (p *Player) SetX(x int) { p.X = x }

// SetY updates the Player's Y coordinate.
func (p *Player) SetY(y int) { p.Y = y }

// newPlayer creates and returns a Player.
func newPlayer(x, y, spriteIndex int) *Player {
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
		HitPoints: 10,
	}
	p.RenderComponent = common.RenderComponent{
		Drawable: charSpritesheet.Cell(spriteIndex),
		Scale:    engo.Point{2, 2},
	}
	p.RenderComponent.SetZIndex(1)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			(charSizeX * float32(x)) + charOffsetX,
			(charSizeY * float32(y)) + charOffsetY,
		},
		Width:  charSizeX,
		Height: charSizeY,
	}
	GameGrid.AddCharacter(p, x, y)
	return p
}

// movePlayer reads keyboard input, checks level bounds, and processes Player
// movement.
// FIXME: Is the ControlEntity parameter needed? We can refer to player globally.
// FIXME: Only call movePlayer from scene when a key has recently been pressed.
func movePlayer(e ControlEntity) {
	var d string
	if engo.Input.Button("moveleft").JustPressed() {
		d = "left"
	} else if engo.Input.Button("moveright").JustPressed() {
		d = "right"
	} else if engo.Input.Button("moveup").JustPressed() {
		d = "up"
	} else if engo.Input.Button("movedown").JustPressed() {
		d = "down"
	}
	if d == "" {
		// Don't process empty keypresses.
		return
	}
	switch d {
	case "left":
		if player.GetX() == GameGrid.MinX {
			log.Println("You can't go that way!")
			return
		} else {
			GameGrid.MoveCharacter(player, player.GetX()-1, player.GetY())
		}
	case "right":
		if player.GetX() == GameGrid.MaxX {
			log.Println("You can't go that way!")
			return
		} else {
			GameGrid.MoveCharacter(player, player.GetX()+1, player.GetY())
		}
	case "up":
		if player.GetY() == GameGrid.MinY {
			log.Println("You can't go that way!")
			return
		} else {
			GameGrid.MoveCharacter(player, player.GetX(), player.GetY()-1)
		}
	case "down":
		if player.GetY() == GameGrid.MaxY {
			log.Println("You can't go that way!")
			return
		} else {
			GameGrid.MoveCharacter(player, player.GetX(), player.GetY()+1)
		}
	}
	posX := float32(player.GetX()) * charSizeX
	posY := float32(player.GetY()) * charSizeY
	if player.GetHitPoints() > 0 {
		// The gravestone is a differently-sized asset which doesn't need an offset.
		posX += charOffsetX
		posY += charOffsetY
	}
	e.SpaceComponent.Position.X = posX
	e.SpaceComponent.Position.Y = posY
}
