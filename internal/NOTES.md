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

* Character interface has changed, necessitating some updates
* NewPlayer no longer calls Grid.AddOccupant
  * Do this from scene/game.go
* CameraComponent no longer destroyed in Player.Destroy()
  * When porting the combat logic, do this from scene/game.go
* MoveCamera() is no longer called in input.go
  * Do this from scene/game.go
* HandleInput() is no longer called in camera.Update()
  * Do this from scene/game.go
