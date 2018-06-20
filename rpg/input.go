package rpg

import "engo.io/engo"

// bindControls assigns the arrow keys to Player movement.
func bindControls() {
	engo.Input.RegisterButton("moveup", engo.KeyArrowUp)
	engo.Input.RegisterButton("movedown", engo.KeyArrowDown)
	engo.Input.RegisterButton("moveleft", engo.KeyArrowLeft)
	engo.Input.RegisterButton("moveright", engo.KeyArrowRight)
}
