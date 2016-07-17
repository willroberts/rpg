// hud.go
package rpg

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

var (
	GameHUD  *HUD
	GameFont *common.Font
)

func NewHUD() (*HUD, error) {
	GameFont = &common.Font{
		URL:  "fonts/Roboto-Regular.ttf",
		BG:   color.Black,
		FG:   color.White,
		Size: 48,
	}
	err := GameFont.CreatePreloaded()
	if err != nil {
		return nil, err
	}

	// Create HP display
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: GameFont,
		Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
	}
	h.SetShader(common.HUDShader)
	h.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{16, 16},
	}

	for _, system := range GameWorld.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&h.BasicEntity,
				&h.RenderComponent,
				&h.SpaceComponent)
		}
	}

	return h, nil
}

func (h *HUD) UpdateHealth() {
	for _, system := range GameWorld.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Remove(h.BasicEntity)
			h.RenderComponent.Drawable = common.Text{
				Font: GameFont,
				Text: fmt.Sprintf("HP: %d", player.GetHitPoints()),
			}
			sys.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
}
