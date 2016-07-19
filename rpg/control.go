// control.go

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

// FIXME: Do these need package scope?
var (
	widthLimit  float32
	heightLimit float32
)

// A ControlComponent is attached to a ControlEntity when being added to the
// ControlSystem.
type ControlComponent struct {
	SchemeVert  string
	SchemeHoriz string
}

// A ControlEntity is anythin which can be tracked by the control system. In our
// case this is usually the player entity.
type ControlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
	*common.SpaceComponent
}

// The ControlSystem keeps track of all controllable entities.
type ControlSystem struct {
	entities []ControlEntity
}

// Add starts tracking an entity in the control system.
func (c *ControlSystem) Add(b *ecs.BasicEntity, ctrl *ControlComponent,
	s *common.SpaceComponent) {
	c.entities = append(c.entities, ControlEntity{b, ctrl, s})
}

// Remove stops tracking an entity in the control system. If the target entity
// is in our tracked entities, remove it.
func (c *ControlSystem) Remove(te ecs.BasicEntity) {
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

// Update processes events for the control system.
func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		heightLimit = levelHeight - e.SpaceComponent.Height
		widthLimit = levelWidth - e.SpaceComponent.Width
		movePlayer(e)
	}
}
