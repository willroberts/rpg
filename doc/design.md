# Design Notes

## Movement and Combat Interactions

### Messy Notes

- Player movement should be in units of 1 tile
- Use some kind of grid data structure to represent this
  - http://stackoverflow.com/questions/1391381/what-data-structures-can-efficiently-store-2-d-grid-data
  - I'm currently using the "Array of Arrays" approach
  - Consider the 2D doubly-linked list if needed

- What's the line of responsibility between Tiled and game code?
  - Tiled for static graphics, game code for dynamic graphics
    - Static graphics: floors, walls, stairs, furniture, etc.
    - Dynamic graphics: characters, enemies, loot, doors?
  - Tiled appears to do more than this (it can handle objects, etc.)
		- Keep it simple for now?

- How do we detect entity collision?
- How does the combat system get notified?
- How does the combat system access the memory of the character and enemy?
- Does the grid keep track of objects?
- Do objects keep track of their position in the grid?

- When an entity attempts movement:
	1. Check the grid to determine if anything is there yet

- The grid is the source of truth. We ask the grid what's next to us
- Entities must tell the grid when they move

### Level Design

1. A level consists of a Tiled map, and a grid of objects over that map
2. The Tiled map only contains static features
3. The grid is an array of arrays which stores entities and objects
4. A tile may have one colliding entity and one non-colliding object
