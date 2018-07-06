# Things cut out of the package move, to be redone elsewhere:

## Player

* failedMovementMsg "You can't go that way!" belongs in whatever prints messages
* MaxHitPoints, ExperienceTable, LevelUp(), Destroy() not implemented
* Grid.AddOccupant() not called in NewPlayer(), should be called from scene
* CameraComponent is no longer destroyed from Player.Destroy()
* Reverse access to Scene is removed. Need to replace Drawable.

## Sprites

## Input

* MoveCamera() no longer called in input.go

## Camera

* scene.HandleInput() no longer called in CameraSystem.Update()
