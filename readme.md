# RPG

A 2D game written in Go, using the [engo](https://www.github.com/EngoEngine/engo)
engine.

## Design

Inspired by
[Castle of the Winds](https://en.wikipedia.org/wiki/Castle_of_the_Winds): an
orthogonal tile-based RPG in which the game advances when the player moves. There
are no animations or sound effects.

## Milestones

Project started July 12, 2016.

### I. Player and Movement (Reached July 14 2016 - Day 3)

The player should be able to control a character with the arrow keys, and
move around within a graphical tiled map. The camera should follow the
player during movement, allowing exploration of the map when the map is
larger than the game window. The player should not be able to move outside
the boundary of the map.

Bugs:

-	[ ] Tiles are not accurately displayed by the engine, and moving the player
	sometimes changes the tiles in a level. Waiting on upstream to fix this.

### II. Combat System (Reached July 17 2016 - Day 6)

The player should be able to encounter and 'fight' enemies by moving in
their direction. When the collision is detected, movement will not be
allowed, but a round of combat will trigger. The player and the enemy will
calculate and deal damage at the same time, before either entity has a
chance to die. When the enemy reaches zero hit points and dies, its sprite should
be removed from the scene. Player hit points should be displayed on screen, and
combat activities should be displayed in a log window.

Bugs:

- [ ] The player can move under the HUD elements. Need to make the game
	boundaries smaller than the game window.

### III. Role-Playing Elements

- [x] Experience Points
- [x] Leveling
- [ ] Character Attributes (STR: ATK, DEX: ACC, INT: MP, END/VIT: HP)
- [ ] Character Creation

### IV. Items

- [ ] Randomized Gold Drops
- [ ] Inventory
- [ ] Randomized Item Drops
- [ ] Equipment

### V. Combat Depth

- [ ] Defensive Spells
- [ ] Offensive Spells
- [ ] Enemy AI
- [ ] Fog of War

### VI. World Depth

- [ ] Starting Town
- [ ] NPCs
- [ ] Vendors
- [ ] Score Tracking and Leaderboards

### VII. Randomized Content

- [ ] "Random" option during Character Creation
- [ ] Randomized Enemy Spawns
- [ ] Randomized Levels, Rooms, Walls, Layouts, etc.

### VIII. Polish

- [ ] Create a better HUD
- [ ] Create a Start Menu
- [ ] Minimal Animations ("Slash" animation for Melee combat, etc.)
- [ ] Thin Fog layer over game window

### IX. Technical Quality

- [ ] Address all TODO and FIXME
- [ ] Proofread and edit documentation
- [ ] Describe a file's purpose on the second comment line (after the filename)

### X. Bonus Content

- [ ] Recreate specific maps from Castle of the Winds
- [ ] Music
- [ ] Sound Effects

## Credits

Art Assets: Jesse Freeman, Tile Crusader
http://jessefreeman.com/articles/free-game-art-tile-crusader
