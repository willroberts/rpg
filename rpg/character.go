package rpg

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type Character struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	ControlComponent
}
