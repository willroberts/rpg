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

type DefaultScene struct{}

var GameWorld *ecs.World

func (scene *DefaultScene) Preload() {
	log.Println("[assets] preloading resources")
	PreloadMapAssets("maps/stone.tmx")

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
	err = engo.Files.Load("fonts/Roboto-Regular.ttf")
	if err != nil {
		panic(err)
	}
}

func (scene *DefaultScene) Setup(w *ecs.World) {
	log.Println("[setup] setting up scene")
	GameWorld = w
	common.SetBackground(color.Black)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&ControlSystem{})

	log.Println("[setup] loading map")
	level, tiles, err := LoadMap("maps/stone.tmx")
	if err != nil {
		panic(err)
	}

	log.Println("[setup] processing grid")
	GameGrid = NewGrid(level.Width(), level.Height())

	log.Println("[setup] creating player")
	player = NewPlayer(1, 1, spriteWhiteZombie)

	log.Println("[setup] creating enemies")
	err = loadEnemyTypes()
	if err != nil {
		panic(err)
	}
	enemies := []*Enemy{
		newEnemy("skeleton", spriteSkeleton, 2, 7),
		newEnemy("skeleton", spriteSkeleton, 8, 6),
		newEnemy("skeleton", spriteSkeleton, 5, 5),
		newEnemy("goblin", spriteGoblin, 4, 11),
		newEnemy("goblin", spriteGoblin, 7, 12),
		newEnemy("bear", spriteBear, 6, 17),
		newEnemy("demon", spriteDemon, 10, 22),
	}

	log.Println("[setup] configuring systems")
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			log.Println("[setup] configuring render system")
			for _, t := range tiles {
				sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
			sys.Add(&player.BasicEntity, &player.RenderComponent,
				&player.SpaceComponent)
			for _, e := range enemies {
				sys.Add(&e.BasicEntity, &e.RenderComponent, &e.SpaceComponent)
			}
		case *ControlSystem:
			log.Println("[setup] configuring control system")
			sys.Add(&player.BasicEntity, &player.ControlComponent,
				&player.SpaceComponent)
		}
	}
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &player.SpaceComponent,
		TrackingBounds: level.Bounds(),
	})

	log.Println("[setup] creating hud")
	GameHUD, err = newHUD()
	if err != nil {
		panic(err)
	}

	log.Println("[setup] binding controls")
	BindControls()
}

func (scene *DefaultScene) Type() string {
	return "DefaultScene"
}
