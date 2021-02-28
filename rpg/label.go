package rpg

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// Label represents a drawable text label on the screen.
type Label struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// NewLabel creates and returns a Label.
func NewLabel(text string, font *common.Font, x, y float32) *Label {
	l := &Label{BasicEntity: ecs.NewBasic()}
	l.RenderComponent.Drawable = common.Text{Font: font, Text: text}
	l.SetShader(common.HUDShader)
	l.SpaceComponent = common.SpaceComponent{Position: engo.Point{X: x, Y: y}}
	return l
}
