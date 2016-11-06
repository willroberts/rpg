package rpg

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const (
	zWorld = iota
	zCharacter
	zHUD
	zText

	hudFormatter string = "Level: %d\nXP: %d\nHP: %d/%d"
)

// The HUD contains all on-screen text, controls, and buttons.
type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Update redraws the HUD to use current values.
func (h *HUD) Update() {
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(h.BasicEntity)
			h.RenderComponent.Drawable = common.Text{
				Font: gameFonts.HUDFont,
				Text: fmt.Sprintf(hudFormatter, gamePlayer.GetLevel(),
					gamePlayer.GetExperience(), gamePlayer.GetHitPoints(),
					gamePlayer.GetMaxHitPoints()),
			}
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
}

// newHUD configures and returns a HUD system.
func newHUD() (*HUD, error) {
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: gameFonts.HUDFont,
		Text: fmt.Sprintf(hudFormatter, gamePlayer.GetLevel(),
			gamePlayer.GetExperience(), gamePlayer.GetHitPoints(),
			gamePlayer.GetMaxHitPoints()),
	}
	h.RenderComponent.SetZIndex(zHUD)
	h.SetShader(common.HUDShader)
	h.SpaceComponent = common.SpaceComponent{Position: engo.Point{0, 0}}
	for _, sys := range gameWorld.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
	return h, nil
}
