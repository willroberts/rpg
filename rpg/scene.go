// scene.go

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
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var GameWorld *ecs.World

// DefaultScene is our first and only scene at the moment. It includes the first
// map, a static set of enemies, and only one room.
type DefaultScene struct{}

// Preload validates and loads assets. In can cause panics, since the game cannot
// run without its assets.
func (scene *DefaultScene) Preload() {
	log.Println("[assets] preloading resources")
	preloadMapAssets("maps/stone.tmx")

	log.Println("[assets] loading sprites")
	err := engo.Files.Load("spritesheets/characters-32x32.png")
	if err != nil {
		panic(err)
	}
	err = engo.Files.Load("spritesheets/decoration-20x20-40x40.png")
	if err != nil {
		panic(err)
	}
	charSpritesheet = common.NewSpritesheetFromFile(
		charSpritesheetPath,
		charSpritesheetWidth,
		charSpritesheetHeight)
	decorationSpritesheet = common.NewSpritesheetFromFile(
		decorationSpritesheetPath,
		decorationSpritesheetWidth,
		decorationSpritesheetHeight)

	log.Println("[assets] loading fonts")
	err = engo.Files.Load("fonts/hud.ttf")
	if err != nil {
		panic(err)
	}
	err = engo.Files.Load("fonts/combatlog.ttf")
	if err != nil {
		panic(err)
	}
}

// Setup initializes all systems necessary for the game to function. It can
// panic, since the game cannot run without these systems.
func (scene *DefaultScene) Setup(w *ecs.World) {
	GameWorld = w
	log.Println("[setup] setting up scene")
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&ControlSystem{})
	log.Println("[setup] loading map")
	level, tiles, err := loadMap("maps/stone.tmx")
	if err != nil {
		panic(err)
	}
	log.Println("[setup] processing grid")
	GameGrid = newGrid(level.Width(), level.Height())
	log.Println("[setup] creating player")
	player = newPlayer("Edmund", spriteWhiteZombie, 1, 1)
	log.Println("[setup] creating enemies")
	if err = loadEnemyTypes(); err != nil {
		panic(err)
	}
	enemies := []*Enemy{
		newEnemy("Skeleton", spriteSkeleton, 2, 7),
		newEnemy("Skeleton", spriteSkeleton, 8, 6),
		newEnemy("Skeleton", spriteSkeleton, 5, 5),
		newEnemy("Goblin", spriteGoblin, 4, 11),
		newEnemy("Goblin", spriteGoblin, 7, 12),
		newEnemy("Bear", spriteBear, 6, 17),
		newEnemy("Demon", spriteDemon, 10, 22),
	}
	log.Println("[setup] configuring systems")
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			log.Println("[setup] configuring render system")
			for _, t := range tiles {
				s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			s.Add(&player.BasicEntity, &player.RenderComponent, &player.SpaceComponent)
			for _, e := range enemies {
				s.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}
		case *ControlSystem:
			log.Println("[setup] configuring control system")
			s.Add(&player.BasicEntity, &player.ControlComponent,
				&player.SpaceComponent)
		}
	}
	log.Println("[setup] configuring camera")
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &player.SpaceComponent,
		TrackingBounds: level.Bounds(),
	})
	log.Println("[setup] creating hud")
	initializeFonts()
	GameHUD, err = newHUD()
	if err != nil {
		panic(err)
	}
	newCombatLog()
	log.Println("[setup] binding controls")
	bindControls()
}

// Type returns the name of the scene. This is used to satisfy engo's Scene
// interface.
func (scene *DefaultScene) Type() string {
	return "DefaultScene"
}
