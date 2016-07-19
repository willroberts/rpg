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
