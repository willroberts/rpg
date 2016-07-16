// input.go
package rpg

import (
	"log"

	"engo.io/engo"
)

func BindControls() {
	log.Println("[input] binding controls")
	engo.Input.RegisterButton("moveup", engo.ArrowUp)
	engo.Input.RegisterButton("movedown", engo.ArrowDown)
	engo.Input.RegisterButton("moveleft", engo.ArrowLeft)
	engo.Input.RegisterButton("moveright", engo.ArrowRight)
	log.Println("[input] controls bound")
	log.Println("[input] use the arrow keys to move")
}
