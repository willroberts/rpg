// enemy.go

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
	"fmt"
	"io/ioutil"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// EnemyAttributes stores the game data which are imported from external sources,
// such as JSON or a database. Once a set of attributes is defined, attributes
// may be modified freely without modifying the game code.
type EnemyAttributes struct {
	HitPoints         int `json:"hitpoints" validate:"nonzero"`
	Damage            int `json:"damage" validate:"nonzero"`
	ExperienceGranted int `json:"experience_granted" validate:"nonzero"`
}

// An Enemy is a non-player Character which is hostile by default.
type Enemy struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Name              string
	Hostility         string
	HitPoints         int
	Damage            int
	ExperienceGranted int

	X, Y int
}

// Destroy removes an enemy from the Grid and from the RenderSystem.
func (e *Enemy) Destroy() {
	gameLog.Update(fmt.Sprintf("%s was destroyed!", e.GetName()))
	gameGrid.RemoveCharacter(e.GetX(), e.GetY())
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(e.BasicEntity)
		}
	}
}

// GetDamage returns the damage dealt by this Enemy.
func (e *Enemy) GetDamage() int { return e.Damage }

// GetHitPoints returns the current HP for this Enemy.
func (e *Enemy) GetHitPoints() int { return e.HitPoints }

// GetHostility returns the demeanor of the enemy for use by the combat system.
// Running into a hostile Character triggers combat, while running into a neutral
// Character does not.
func (e *Enemy) GetHostility() string { return e.Hostility }

// GetName returns the name of this Enemy type, which is used by newEnemy to
// retrieve its EnemyAttributes.
func (e *Enemy) GetName() string { return e.Name }

// GetXPBonus returns the number of experience points granted upon killing this
// enemy type.
func (e *Enemy) GetXPBonus() int { return e.ExperienceGranted }

// GetX returns the Enemy's X coordinate.
func (e *Enemy) GetX() int { return e.X }

// GetY returns the Enemey's Y coordinate.
func (e *Enemy) GetY() int { return e.Y }

// ModifyHitPoints adds a number to an Enemy's hit points. To deal damage to an
// Enemy, provide a negative number.
func (e *Enemy) ModifyHitPoints(amount int) {
	e.HitPoints += amount
	if e.HitPoints < 0 {
		e.HitPoints = 0
	}
}

// SetHostility changes an Enemy's demeanor. This can be used to implement "calm"
// spells against typically hostile enemies, or to create hostility in a
// typically neutral non-player Character.
func (e *Enemy) SetHostility(h string) { e.Hostility = h }

// SetX updates the Enemy's X coordinate.
func (e *Enemy) SetX(x int) { e.X = x }

// SetY updates the Enemy's Y coordinate.
func (e *Enemy) SetY(y int) { e.Y = y }

// loadEnemyTypes reads EnemyAttributes from JSON, and populates the global map
// gameEnemyTypes.
func loadEnemyTypes() error {
	b, err := ioutil.ReadFile("data/enemies.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &gameEnemyTypes); err != nil {
		return err
	}
	return nil
}

// newEnemy creates and returns an Enemy.
func newEnemy(name string, spriteIndex, x, y int) *Enemy {
	e := &Enemy{
		BasicEntity:       ecs.NewBasic(),
		Name:              name,
		Hostility:         "hostile",
		HitPoints:         gameEnemyTypes[name].HitPoints,
		Damage:            gameEnemyTypes[name].Damage,
		ExperienceGranted: gameEnemyTypes[name].ExperienceGranted,
		X:                 x,
		Y:                 y,
	}
	enemyTexture := gameSprites.Characters.Cell(spriteIndex)
	e.RenderComponent = common.RenderComponent{
		Drawable: enemyTexture,
		Scale:    engo.Point{2, 2},
	}
	e.RenderComponent.SetZIndex(zCharacter)
	e.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			(charSizeX * float32(x)) + charOffsetX,
			(charSizeY * float32(y)) + charOffsetY,
		},
		Width:  charSizeX,
		Height: charSizeY,
	}
	gameGrid.AddCharacter(e, x, y)
	return e
}
