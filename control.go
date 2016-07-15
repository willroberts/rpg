package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type ControlComponent struct {
	SchemeVert  string
	SchemeHoriz string
}

type ControlSystem struct {
	entities []controlEntity
}

type controlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
	*common.SpaceComponent
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, control *ControlComponent,
	space *common.SpaceComponent) {
	c.entities = append(c.entities, controlEntity{basic, control, space})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	del := -1

	// Determine if the requested entity is in our entities slice.
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			del = index
			break
		}
	}

	// If we found the entity, delete it.
	if del >= 0 {
		c.entities = append(c.entities[:del], c.entities[del+1:]...)
	}
}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		// Move the character.
		// This currently works because the character is the only entity.
		// As soon as we add another entity, the movement needs to be changed
		// to a speed model (based on dt) instead of a fixed rate model.
		if engo.Input.Button("moveup").JustPressed() {
			e.SpaceComponent.Position.Y -= 80
		}
		if engo.Input.Button("movedown").JustPressed() {
			e.SpaceComponent.Position.Y += 80
		}
		if engo.Input.Button("moveleft").JustPressed() {
			e.SpaceComponent.Position.X -= 80
		}
		if engo.Input.Button("moveright").JustPressed() {
			e.SpaceComponent.Position.X += 80
		}

		// Detect when the player attempts to move outside the game window.
		var heightLimit float32 = levelHeight - e.SpaceComponent.Height
		var widthLimit float32 = levelWidth - e.SpaceComponent.Width

		if e.SpaceComponent.Position.Y < 0 {
			// Dirty hack: retain the character's Y offset
			e.SpaceComponent.Position.Y = 4
		} else if e.SpaceComponent.Position.Y > heightLimit {
			// Dirty hack: retain the character's Y offset
			e.SpaceComponent.Position.Y = heightLimit + 4
		}

		if e.SpaceComponent.Position.X < 0 {
			// Dirty hack: retain the character's X offset
			e.SpaceComponent.Position.X = 6
		} else if e.SpaceComponent.Position.X > widthLimit {
			// Dirty hack: retain the character's X offset
			e.SpaceComponent.Position.X = widthLimit + 6
		}
	}
}
