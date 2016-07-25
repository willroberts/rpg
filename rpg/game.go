// game.go

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
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo/common"
)

// GameScene is our first and only scene at the moment. It includes the first
// map, a static set of enemies, and only one room.
type GameScene struct{}

// Preload validates and loads assets. In can cause panics, since the game cannot
// run without its assets.
func (scene *GameScene) Preload() {
	log.Println("preloading maps")
	preloadMapAssets("maps/stone.tmx")

	log.Println("preloading sprites")
	if err := preloadSprites(); err != nil {
		panic(err)
	}

	var err error
	gameFonts, err = PreloadFonts()
	if err != nil {
		panic(err)
	}
}

// Setup initializes all systems necessary for the game to function. It can
// panic, since the game cannot run without these systems.
func (scene *GameScene) Setup(w *ecs.World) {
	log.Println("creating scene")
	gameWorld = w
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&CameraSystem{})

	log.Println("loading map")
	level, tiles, err := loadMap("maps/stone.tmx")
	if err != nil {
		panic(err)
	}

	log.Println("creating level grid")
	gameGrid = newGrid(level.Width(), level.Height())

	log.Println("creating player")
	gamePlayer = newPlayer("Edmund", spriteWhiteZombie, 1, 1)
	if err := loadExperienceTable(); err != nil {
		panic(err)
	}

	log.Println("creating enemies")
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

	log.Println("configuring systems")
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			log.Println("configuring render system")
			for _, t := range tiles {
				s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			s.Add(&gamePlayer.BasicEntity, &gamePlayer.RenderComponent, &gamePlayer.SpaceComponent)
			for _, e := range enemies {
				s.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}
		case *CameraSystem:
			log.Println("configuring camera system")
			s.Add(&gamePlayer.BasicEntity, &gamePlayer.CameraComponent,
				&gamePlayer.SpaceComponent)
		}
	}

	log.Println("configuring camera")
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &gamePlayer.SpaceComponent,
		TrackingBounds: level.Bounds(),
	})

	log.Println("creating hud")
	gameHUD, err = newHUD()
	if err != nil {
		panic(err)
	}
	gameLog = newActivityLog()
	gameLog.Update("Welcome to the game.")
	gameLog.Update("There are three skeletons near you.")
	gameLog.Update("Try moving into them to attack.")

	log.Println("binding controls")
	bindControls()
	log.Println("use the arrow keys to move")
}

// Type returns the name of the scene. This is used to satisfy engo's Scene
// interface.
func (scene *GameScene) Type() string {
	return "GameScene"
}
