# Things cut out of the package move, to be redone elsewhere:

## Player

* failedMovementMsg "You can't go that way!" belongs in whatever prints messages
* MaxHitPoints, ExperienceTable, LevelUp(), Destroy() not implemented
* Grid.AddOccupant() not called in NewPlayer(), should be called from scene
* CameraComponent is missing

## Sprites

## Camera

* scene.HandleInput() no longer called in CameraSystem.Update()
