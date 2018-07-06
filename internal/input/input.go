package input

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/willroberts/rpg/internal/char"
	"github.com/willroberts/rpg/internal/grid"
	"go.uber.org/zap"
)

type InputSystem struct {
	Player char.Character
	Grid   grid.Grid
	Logger *zap.Logger
}

func (i *InputSystem) Remove(e ecs.BasicEntity) {}

func (i *InputSystem) Update(dt float32) {
	delta := ProcessInput()

	// Debug logging.
	if delta.X != 0 || delta.Y != 0 {
		i.Logger.Info("move delta",
			zap.Int("x", delta.X),
			zap.Int("y", delta.Y),
		)
	}

	// Move the player in the Grid.
	toX := i.Player.GetX() + delta.X
	toY := i.Player.GetY() + delta.Y
	if !i.Grid.CanMoveTo(toX, toY) {
		return
	}
	i.Grid.MoveChar(i.Player, toX, toY)
}

type MoveDelta struct {
	X int
	Y int
}

func BindControls() {
	engo.Input.RegisterButton("moveup", engo.KeyArrowUp)
	engo.Input.RegisterButton("movedown", engo.KeyArrowDown)
	engo.Input.RegisterButton("moveleft", engo.KeyArrowLeft)
	engo.Input.RegisterButton("moveright", engo.KeyArrowRight)
}

func ProcessInput() MoveDelta {
	// FIXME: This is too specific to movement. If we add non-movement keyinds,
	// then MoveDelta will not be an appropriate return type here.
	if engo.Input.Button("moveleft").JustPressed() {
		return MoveDelta{X: -1, Y: 0}
	} else if engo.Input.Button("moveright").JustPressed() {
		return MoveDelta{X: 1, Y: 0}
	} else if engo.Input.Button("moveup").JustPressed() {
		return MoveDelta{X: 0, Y: 1}
	} else if engo.Input.Button("movedown").JustPressed() {
		return MoveDelta{X: 0, Y: -1}
	} else {
		// No key was pressed, or a key was pressed which has no binding.
		return MoveDelta{X: 0, Y: 0}
	}
}
