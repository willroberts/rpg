// camera.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA

package rpg

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

// A CameraComponent is attached to a CameraEntity when being added to the
// CameraSystem.
type CameraComponent struct {
	SchemeVert  string
	SchemeHoriz string
}

// A CameraEntity is anythin which can be tracked by the camera system. In our
// case this is usually the player entity.
type CameraEntity struct {
	*ecs.BasicEntity
	*CameraComponent
	*common.SpaceComponent
}

// The CameraSystem keeps track of all cameralable entities.
type CameraSystem struct {
	entities []CameraEntity
}

// Add starts tracking an entity in the camera system.
func (c *CameraSystem) Add(b *ecs.BasicEntity, ctrl *CameraComponent,
	s *common.SpaceComponent) {
	c.entities = append(c.entities, CameraEntity{b, ctrl, s})
}

// Remove stops tracking an entity in the camera system. If the target entity
// is in our tracked entities, remove it.
func (c *CameraSystem) Remove(te ecs.BasicEntity) {
	del := -1
	for i, e := range c.entities {
		if e.BasicEntity.ID() == te.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		c.entities = append(c.entities[:del], c.entities[del+1:]...)
	}
}

// Update processes events for the camera system.
func (c *CameraSystem) Update(dt float32) {
	for _, e := range c.entities {
		movePlayer(e)
	}
}
