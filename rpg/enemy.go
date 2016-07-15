package rpg

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

type Enemy struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
