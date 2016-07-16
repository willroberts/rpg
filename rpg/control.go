// control.go
package rpg

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

var (
	heightLimit float32
	widthLimit  float32
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
		heightLimit = levelHeight - e.SpaceComponent.Height
		widthLimit = levelWidth - e.SpaceComponent.Width
		movePlayer(e)
	}
}
