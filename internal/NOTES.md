# Package Migration Checklist

## To Do

* HUD
* Activity Log
* Enemies
* Combat

## Types

* Enemy, EnemyAttributes
* FontSet
* HUD
* Label
* ActivityLog, ActivityLogWindow, ActivityLogMessage

## Functions

* HandleCombat (combat.go)
* LoadEnemyTypes, NewEnemy (enemy.go)
* LoadFont, PreloadFont (font.go)
* NewHUD, UpdateHUD (hud.go)
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

* CameraComponent no longer destroyed in Player.Destroy()
  * When porting the combat logic, do this from scene/game.go
