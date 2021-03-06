# Code Design Documentation

A `main.go` file does the following:

- Configure an Engo window with a title, width, and height
- Create an instance of rpg.GameScene
- Call engo.Run on the game scene

## Type Hierarchy

- GameScene
  - World
	  - Systems
	    - CameraSystem
			  - CameraEntity: Calls movePlayer()
				  - SpaceComponent.Position.{X,Y}: Modified by Player
  - FontSet
	- SpriteSet
	- Grid
	- HUD
	- ActivityLog
	- Player
	  - ExperienceTable
	  - CameraComponent
	- Enemies

### Issues

The hierarchy is generally pretty clean, but there are some improvements to be
made:

1. The camera moves the player by calling its move function, and the player
   moves the camera by setting its position. This is circular, and should be
	 changed so that the input keys move the player without the camera being
	 involved, and the camera's position is determined on demand based on the
	 player's coordinates.

## Game Scene

The `GameScene` type is the entry point for our game. It satisfies
the `engo.Scene` interface.

`GameScene`'s `Preload()` function loads static assets needed by the
game.

`GameScene`'s `Setup()` function initializes its subsystems, creates
the scene (map, grid, etc.) and adds the player and enemies to the
scene. It also configures the camera, the HUD, the log, and the
input bindings.

Once Engo has run the scene, the game engine advances on player
input. The `Player` type call's the `MoveCharacter` method in the
`Grid` interface. If there is a collision with an enemy, the grid
calls `handleCombat`.

## Entities

### Character / Player / Enemy

Files: `character.go`, `player.go`, `enemy.go`

`Character` is a generic interface for a character, whether it is a
player character or an NPC (including enemies). `Player` and `Enemy`
satisfy the `Character` interface.

## Systems using the ECS Paradigm

### Camera

Files: `camera.go`

The camera system controls the player perspective of the map, and
handles player movement events by calling `movePlayer()`.

### Combat

Files: `combat.go`

The combat system handles HP deductions, player and enemy death, and
the granting of experience points.

### Grid

Files: `grid.go`

The grid system stores the actual game state for the given map. It
processes collision events and can have entities added or removed.

### HUD

Files: `hud.go`

The HUD system displays player level, player health, and player
experience points. It updates once per frame, though it should
update more punctually.

### Activity Log

Files: `log.go`

The activity log system is a text log of game events and combat
results displayed on screen during gameplay.

### Tiled (TMX) Tilemaps

Files: `tilemap.go`

The tilemap system is responsible for preloading a Tiled map and
creating ECS entities and components for all necessary objects.

## Helper Systems

### Fonts

Files: `fonts.go`

The fonts system is a set of helpers for preloading fonts which are returned in
a `FontSet` type.

### Input

Files: `input.go`

This is a set of helpers for binding our key controls.

### Label

Files: `label.go`

The Label system is unfinished, but it was intended to facilitate
the creation of text labels for the main menu.

### Sprites

Files: `sprite.go`

The sprites code is a set of helpers for theloading of sprite graphics.
