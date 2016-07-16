# Bugs

* On collision, the player is moved to the tile under the enemy
	* No movement should occur, just like running into a wall
  * Use an cancelMovement() func and an engo.Message?
  * Need to efficiently and conveniently store entity locations in the tile map

* `engo.Input.Button("foo").JustPressed()` fires twice
	* Once when pressing the button, and once when releasing it
	* Appears to be a bug with Ubuntu 16.04 and GLFW 3.1.2
	* https://github.com/EngoEngine/engo/issues/279
	* https://github.com/glfw/glfw/issues/747

* The tiles selected in Tiled do not appear in-game
	* Higher-indexed tiles cause an `index out of range` error
