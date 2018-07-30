package text

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Label represents a drawable text label on the screen.
type Label struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// NewLabel creates and returns a Label.
func NewLabel(contents string, font *common.Font, x, y float32) *Label {
	l := &Label{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: engo.Point{
				X: x,
				Y: y,
			},
		},
	}
	l.RenderComponent.Drawable = common.Text{Font: font, Text: contents}
	l.SetShader(common.HUDShader)
	return l
}
