package rpg

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

// Creates and returns a new Label.
func NewLabel(text string, font *common.Font, x, y float32) *Label {
	l := &Label{BasicEntity: ecs.NewBasic()}
	l.RenderComponent.Drawable = common.Text{Font: font, Text: text}
	l.SetShader(common.HUDShader)
	l.SpaceComponent = common.SpaceComponent{Position: engo.Point{x, y}}
	return l
}
