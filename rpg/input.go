package rpg

import "engo.io/engo"

// bindControls assigns the arrow keys to Player movement.
func bindControls() {
	engo.Input.RegisterButton("moveup", engo.KeyArrowUp)
	engo.Input.RegisterButton("movedown", engo.KeyArrowDown)
	engo.Input.RegisterButton("moveleft", engo.KeyArrowLeft)
	engo.Input.RegisterButton("moveright", engo.KeyArrowRight)
}

func (g *GameScene) HandleInput() {
	var d string
	if engo.Input.Button("moveleft").JustPressed() {
		d = "left"
	} else if engo.Input.Button("moveright").JustPressed() {
		d = "right"
	} else if engo.Input.Button("moveup").JustPressed() {
		d = "up"
	} else if engo.Input.Button("movedown").JustPressed() {
		d = "down"
	}
	if d == "" {
		// Don't process unbound keypresses.
		return
	}

	switch d {
	case "left":
		if g.Player.GetX() == g.Grid.MinX {
			return
		}
		g.Grid.MoveCharacter(g.Player, g.Player.GetX()-1, g.Player.GetY())
	case "right":
		if g.Player.GetX() == g.Grid.MaxX {
			return
		}
		g.Grid.MoveCharacter(g.Player, g.Player.GetX()+1, g.Player.GetY())
	case "up":
		if g.Player.GetY() == g.Grid.MinY {
			return
		}
		g.Grid.MoveCharacter(g.Player, g.Player.GetX(), g.Player.GetY()-1)
	case "down":
		if g.Player.GetY() == g.Grid.MaxY {
			return
		}
		g.Grid.MoveCharacter(g.Player, g.Player.GetX(), g.Player.GetY()+1)
	}

	x := float32(g.Player.GetX()) * charSizeX
	y := float32(g.Player.GetY()) * charSizeY

	// If the player is alive, give the character a position offset. The player
	// turns into a tombstone on death, and the tombstone asset does not need an
	// offset.
	if g.Player.GetHitPoints() > 0 {
		x += charOffsetX
		y += charOffsetY
	}

	// Move the camera.
	for _, sys := range g.World.Systems() {
		switch s := sys.(type) {
		case *CameraSystem:
			for _, e := range s.Entities {
				e.SpaceComponent.Position.X = x
				e.SpaceComponent.Position.Y = y
			}
		}
	}
}
