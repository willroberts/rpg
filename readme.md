# rpg

* 2d graphics
  * tile based
  * not isometric
  * game logic advances when player moves
  * castle of the winds
* game logic written in go
* EngoEngine/engo as the engine
* Tiled for level management
* you play as a misunderstood skeleton in a dungeon fending off "intruders" (heroes)

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

* figure out why i can still go out of bounds
* look into KeyPressed() bug (extra KeyPressed() event on KeyReleased() event)
* why does engo add some number to the tile id for maps? selecting grey stone gives green stone, etc.
* add logging on everything to make sure i can describe what it's doing

## art assets

* http://jessefreeman.com/articles/free-game-art-tile-crusader

## other projects

* drawing with opengl, doesn't have to be a game
  * learn more about opengl, graphics, game programming in the process
  * use go-gl 3.3 (generated from c bindings)
