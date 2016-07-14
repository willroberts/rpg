# rpg

* 2d graphics
  * tile based
  * not isometric
  * game logic advances when player moves
  * castle of the winds
* game logic written in go
* EngoEngine/engo as the engine
* Tiled for level management

## milestones

* [x] tiled maps
* [x] characters
* window collision
* enemies
* enemy collision
* walls
* wall collision
* combat (start with unarmed)
* experience and leveling
* items and loot
* high scores

## bugs and to do

* retain character centeredness after hitting a window border
* look into Key.JustPressed() bug
  * https://github.com/EngoEngine/engo/issues/279
  * https://github.com/glfw/glfw/issues/747
  * ignore for now
* why does engo add some number to the tile id for maps?
  * selecting grey stone gives green stone, etc.
  * punt on this for now
* add more logging
* add more documentation

## art assets

* http://jessefreeman.com/articles/free-game-art-tile-crusader

## other projects

* drawing with opengl, doesn't have to be a game
  * learn more about opengl, graphics, game programming in the process
  * use go-gl 3.3 (generated from c bindings)
