# RPG

A 2D game written in Go, with the engo engine.

## Design

Recreation of Castle of the Winds, an orthogonal, tile-based RPG in which
the world advances when the player moves. There are no animations, and
there is no audio.

## Milestones

I. Minimum Interactive Content [x]
	a. Tiled Maps [x]
	b. Controllable Player [x]
	c. Player-Centric Camera [x]
	d. Window Bounds Checking [x]

At this stage, the player should be able to control a character with the
arrow keys, and move them around within a graphical tiled map. The camera
should follow the player during movement, allowing exploration of the map
when the map is larger than the game window. The player should not be able
to move outside the boundary of the map.

II. Combat [ ]
	a. Enemies [ ]
	b. Enemy Collision [ ]
	c. Combat Calculation [ ]
	d. Enemy Death [ ]

At this stage, the player should be able to encounter and 'fight' enemies
by moving in their direction. When the collision is detected, movement
will not be allowed, but a round of combat will trigger. The player and
the enemy will calculate and deal damage at the same time, before either
entity has a chance to die. When the enemy dies, its sprite should be
removed from the scene. Combat should use hardcoded damage values
initially. One approach would be to implement unarmed combat first.

III. Depth [ ]
	a. Character Data Model [ ]
		i. Attributes [ ]
		ii. Level [ ]
		iii. Experience [ ]
    iv. Inventory [ ]
		v. Gold [ ]
  b. Item Data Model [ ]
  c. Loot System [ ]
  d. Intelligence, Mana, and Spells [ ]
     
IV. Polish [ ]
	a. Walls and Wall Collision [ ]
	b. Start Menu [ ]
	c. High Scores [ ]

## Issues

* `engo.Input.Button("foo").JustPressed()` fires twice
	* Once when pressing the button, and once when releasing it
	* Appears to be a bug with Ubuntu 16.04 and GLFW 3.1.2
  * https://github.com/EngoEngine/engo/issues/279
  * https://github.com/glfw/glfw/issues/747
* The tiles selected in Tiled do not appear in-game
  * Higher-indexed tiles cause an `index out of range` error

## Assets

2D art comes from Jesse Freeman's Tile Crusader spritesheets:
http://jessefreeman.com/articles/free-game-art-tile-crusader
