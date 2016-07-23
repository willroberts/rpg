# RPG

A 2D game written in Go, with the engo engine.

## Design

Basic recreation of Castle of the Winds (1993): an orthogonal tile-based RPG in
which the game advances when the player moves. There are no animations,
and there is no audio.

## Milestones

### I. Interaction

- [x] Tiled Maps
- [x] Controllable Player
- [x] Player-Centric Camera
- [x] Window Bounds Checking

The player should be able to control a character with the arrow keys, and
move around within a graphical tiled map. The camera should follow the
player during movement, allowing exploration of the map when the map is
larger than the game window. The player should not be able to move outside
the boundary of the map.

### II. Combat

- [x] Enemies
- [x] Enemy Collision
- [x] Hit Points
- [x] Combat Calculation
- [x] Enemy Death
- [x] Player Death
- [x] HP Display
- [x] On-Screen Combat Log

The player should be able to encounter and 'fight' enemies by moving in
their direction. When the collision is detected, movement will not be
allowed, but a round of combat will trigger. The player and the enemy will
calculate and deal damage at the same time, before either entity has a
chance to die. When the enemy dies, its sprite should be removed from the
scene.

### III. Increased Depth

- [x] Experience Points and Leveling
- [ ] Enemies drop Gold
	- [ ] Player can collect Gold
- [ ] Enemies drop Items
	- [ ] Player can store Items
	- [ ] Player can equip Items
- [ ] Character Creation
- [ ] Character Attribute: Endurance
	-	[ ] +1 HP per Something
- [ ] Character Attribute: Strength
	- [ ] +1 Damage per Something
- [ ] Character Attribute: Dexterity
	- [ ] Accuracy System
- [ ] Character Attribute: Intelligence
	- [ ] Mana
	- [ ] +1 Mana per Something
	- [ ] Defensive Spells
	- [ ] Offensive Spells

### IV. Improvements

- [ ] Make HUD not part of the moveable surface
- [ ] Make HUD look better
- [ ] Start Menu
- [ ] Artificial Intelligence
	- [ ] Enemies can see Player within N tiles
	- [ ] Enemies advance one step toward the player when spotted
	-	[ ] Enemies have limited decision making and may flee
- [ ] Fog of War
- [ ] Address all FIXME and TODO notes
- [ ] Add a comment under filename.go summarizing what it contains
- [ ] Fix tile rendering bug

### V. Replayability

- [ ] Generate randomized Characters
- [ ] Generate randomized Items
- [ ] Generate randomized Enemies
- [ ] Generate randomized Vendors
- [ ] Generate randomized Maps
	- [ ] Implement Walls and Wall collision

### VI. Bonus Features

- [ ] Recreate the map and enemies from the first Castle of the Winds episode
- [ ] Score Tracking
	- [ ] High Scores
- [ ] Minimal Combat Animations
	-	[ ] Slight nudge of sprite towards occupied tile for <1s
	- [ ] "Slash" sprite over damaged enemy
- [ ] Music
- [ ] Sound Effects
- [ ] Thin fog layer over the entire game window

## Credits

Art Assets: Jesse Freeman, Tile Crusader
http://jessefreeman.com/articles/free-game-art-tile-crusader
