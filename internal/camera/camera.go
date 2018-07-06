package camera

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type CameraComponent struct {
	SchemeHoriz string
	SchemeVert  string
}

type CameraEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*CameraComponent
}

type CameraSystem struct {
	Entities []CameraEntity
}

func (c *CameraSystem) Add(e *ecs.BasicEntity, sc *common.SpaceComponent, cc *CameraComponent) {
	c.Entities = append(c.Entities, CameraEntity{e, sc, cc})
}

func (c *CameraSystem) Remove(e ecs.BasicEntity) {
	del := -1
	for i, ent := range c.Entities {
		if ent.BasicEntity.ID() == e.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		c.Entities = append(c.Entities[:del], c.Entities[del+1:]...)
	}
}

func (c *CameraSystem) Update(dt float32) {
}
