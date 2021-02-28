package rpg

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

// CameraComponent is attached to a CameraEntity when being added to the
// CameraSystem.
type CameraComponent struct {
	SchemeVert  string
	SchemeHoriz string
}

// CameraEntity is anything which can be tracked by the camera system. In our
// case this is usually the player entity.
type CameraEntity struct {
	*ecs.BasicEntity
	*CameraComponent
	*common.SpaceComponent
}

// CameraSystem keeps track of all camera entities.
type CameraSystem struct {
	Entities []CameraEntity
}

// Add starts tracking an entity in the camera system.
func (c *CameraSystem) Add(b *ecs.BasicEntity, ctrl *CameraComponent,
	s *common.SpaceComponent) {
	c.Entities = append(c.Entities, CameraEntity{b, ctrl, s})
}

// Remove stops tracking an entity in the camera system.
func (c *CameraSystem) Remove(te ecs.BasicEntity) {
	del := -1
	for i, e := range c.Entities {
		if e.BasicEntity.ID() == te.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		c.Entities = append(c.Entities[:del], c.Entities[del+1:]...)
	}
}

// Update processes events for the camera system.
func (c *CameraSystem) Update(dt float32) {
	// FIXME: Move HandleInput to a new InputSystem.
	gameScene.HandleInput()
}
