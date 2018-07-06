# Package Migration Checklist

## Types

* Enemy, EnemyAttributes
* FontSet
* HUD
* Label
* ActivityLog, ActivityLogWindow, ActivityLogMessage

## Functions

* HandleCombat (combat.go)
* Character Interface methods (enemy.go, player.go)
* LoadEnemyTypes, NewEnemy (enemy.go)
* LoadFont, PreloadFont (font.go)
* NewHUD, UpdateHUD (hud.go)
* BindControls, HandleInput (input.go)
* NewLabel (label.go)
* NewActivityLog, InitializeMessage, DrawMessage, UpdateLog (log.go)

## Variables

* Experience table (experience.go)

## Constants

* Font TTF filenames (font.go)
* HUD format string (hud.go)
* HUD and Text Z-indices (hud.go)
* Log Window dimensions (log.go)
* Failed movement message (player.go)

# Functional Changes

* The InputSystem previously handled out-of-bounds checking for the grid.
  * Now we return an intent with X and Y values instead
	* The Scene and Grid should take this responsibility (func CanMove() bool)
	* Once movement is processed, we also update the SpaceComponent of the Camera entity
* Character interface has changed, necessitating some updates
* CameraComponent no longer destroyed in Player.Destroy()
  * When porting the combat logic, do this from scene/game.go
* MoveCamera() is no longer called in input.go
  * Do this from scene/game.go
* HandleInput() is no longer called in camera.Update()
  * Do this from scene/game.go
