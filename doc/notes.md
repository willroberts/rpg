# Notes

- Never rely on pixels or graphics for determining entity position
	- Always implement your own position tracking

## tmux for window management

^b + $: rename session
^b + ,: rename window
^b + o: cycle pane
^b + %: split pane (left and right)
^b + ": split pane (top and bottom)

## vim-go for ide features

:GoRun						Run main(), press Enter to return to IDE
:GoBuild					Shows build errors in a quickfix list
:GoTest						Show test errors in a quickfix list
:GoTestFunc				Test the function under the cursor only
:GoTestCompile		Test compilation (doesn't run tests)
:GoCoverage				Visually highlight test coverage
:GoCoverageClear	Remove coverage highlighting
:GoCoverageToggle	Toggle coverage highlighting

## Map Tile Issues

* When the game window is larger than the game world:
	* Areas outside the map should be black
	* Instead, tiles not used by the map are placed there
	* The camera shifts when moving closer to one side
* Without the above issue:
	* A block of 16 tiles (4x4) use the wrong sprite
	* The last tile (last x and y coordinates) uses the wrong sprite
