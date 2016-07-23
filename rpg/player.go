// player.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

package rpg

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const failedMovementMsg string = "You can't go that way!"

// Player is the player-controlled entity in the game.
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	CameraComponent

	Name         string
	Hostility    string
	HitPoints    int
	MaxHitPoints int
	Level        int
	Experience   int

	X, Y int
}

// Destroy removes a Player entity from the game. When the player dies, we
// replace the character texture and stop processing incoming movement commands.
// Removes the Player from the CameraSystem and the RenderSystem, and then re-add
// the Player to the RenderSystem as a gravestone.
func (p *Player) Destroy() {
	gameLog.Update("You died!")
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *CameraSystem:
			s.Remove(p.BasicEntity)
		case *common.RenderSystem:
			s.Remove(p.BasicEntity)
			p.RenderComponent = common.RenderComponent{
				Drawable: gameSpritesDeco.Cell(spriteGravestone),
				Scale:    engo.Point{charScale, charScale},
			}
			p.RenderComponent.SetZIndex(zCharacter)
			s.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		}
	}
}

// GetDamage returns the damage dealt by the Player.
func (p *Player) GetDamage() int { return 1 }

// GetHitPoints returns the current HP for the Player.
func (p *Player) GetHitPoints() int { return p.HitPoints }

// GetMaxHitPoints returns the maximum HP for the Player.
func (p *Player) GetMaxHitPoints() int { return p.MaxHitPoints }

// GetLevel returns the current Level of the Player.
func (p *Player) GetLevel() int { return p.Level }

// GetExperience returns the current XP for the Player.
func (p *Player) GetExperience() int { return p.Experience }

// GetHostility returns the demeanor of an enemy or NPC. It's included here just
// to statisfy the Character interface.
// FIXME: See if we can remove this from the Character interface.
func (p *Player) GetHostility() string { return p.Hostility }

// GetName returns the name of the Player.
func (p *Player) GetName() string { return p.Name }

// GetXPBonus is stubbed for Player entities. Only needed for Enemy entities, but
// we include it here to satisfy the Character interface.
func (p *Player) GetXPBonus() int { return 0 }

// GetX returns the Player's X coordinate.
func (p *Player) GetX() int { return p.X }

// GetY returns the Player's Y coordinate.
func (p *Player) GetY() int { return p.Y }

// ModifyHitPoints adds a number to the Player's hit points. To deal damage to
// the Player, provide a negative number.
func (p *Player) ModifyHitPoints(amount int) {
	p.HitPoints += amount
	if p.HitPoints < 0 {
		p.HitPoints = 0
	}
}

// ModifyExperience adds a number to the Player's experience points.
func (p *Player) ModifyExperience(amount int) {
	p.Experience += amount
	nextLevel := strconv.Itoa(p.Level + 1)
	if p.Experience >= gameExperienceTable[nextLevel] {
		p.LevelUp()
	}
}

// LevelUp increases the Player's level, increases MaxHitPoints, and restores HP.
func (p *Player) LevelUp() {
	gameLog.Update("Level up!")
	p.Level += 1
	p.Experience = 0
	p.MaxHitPoints += 2
	p.HitPoints = p.MaxHitPoints
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
func newPlayer(name string, spriteIndex, x, y int) *Player {
	p := &Player{
		BasicEntity: ecs.NewBasic(),
		CameraComponent: CameraComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
		X:            x,
		Y:            y,
		Name:         name,
		Hostility:    "neutral",
		HitPoints:    20,
		MaxHitPoints: 20,
		Level:        1,
		Experience:   0,
	}
	p.RenderComponent = common.RenderComponent{
		Drawable: gameSpritesChar.Cell(spriteIndex),
		Scale:    engo.Point{charScale, charScale},
	}
	p.RenderComponent.SetZIndex(zCharacter)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			(charSizeX * float32(x)) + charOffsetX,
			(charSizeY * float32(y)) + charOffsetY,
		},
		Width:  charSizeX,
		Height: charSizeY,
	}
	gameGrid.AddCharacter(p, x, y)
	return p
}

// movePlayer reads keyboard input, checks level bounds, and processes Player
// movement.
// FIXME: Is the CameraEntity parameter needed? We can refer to player globally.
// FIXME: Only call movePlayer from scene when a key has recently been pressed.
// FIXME: Make this a function bound to *Player instead
func movePlayer(e CameraEntity) {
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
		if gamePlayer.GetX() == gameGrid.MinX {
			//gameLog.Update(failedMovementMsg)
			return
		}
		gameGrid.MoveCharacter(gamePlayer, gamePlayer.GetX()-1, gamePlayer.GetY())
	case "right":
		if gamePlayer.GetX() == gameGrid.MaxX {
			//gameLog.Update(failedMovementMsg)
			return
		}
		gameGrid.MoveCharacter(gamePlayer, gamePlayer.GetX()+1, gamePlayer.GetY())
	case "up":
		if gamePlayer.GetY() == gameGrid.MinY {
			//gameLog.Update(failedMovementMsg)
			return
		}
		gameGrid.MoveCharacter(gamePlayer, gamePlayer.GetX(), gamePlayer.GetY()-1)
	case "down":
		if gamePlayer.GetY() == gameGrid.MaxY {
			//gameLog.Update(failedMovementMsg)
			return
		}
		gameGrid.MoveCharacter(gamePlayer, gamePlayer.GetX(), gamePlayer.GetY()+1)
	}
	posX := float32(gamePlayer.GetX()) * charSizeX
	posY := float32(gamePlayer.GetY()) * charSizeY
	if gamePlayer.GetHitPoints() > 0 {
		// The gravestone is a differently-sized asset which doesn't need an offset.
		posX += charOffsetX
		posY += charOffsetY
	}
	e.SpaceComponent.Position.X = posX
	e.SpaceComponent.Position.Y = posY
}

// loadExperienceTable reads the ExperienceTable from data/experience.json.
func loadExperienceTable() error {
	b, err := ioutil.ReadFile("data/experience.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &gameExperienceTable); err != nil {
		return err
	}
	return nil
}
