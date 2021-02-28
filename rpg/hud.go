package rpg

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
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
	for _, sys := range gameScene.World.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(h.BasicEntity)
			h.RenderComponent.Drawable = common.Text{
				Font: gameScene.Fonts.HUDFont,
				Text: fmt.Sprintf(hudFormatter, gameScene.Player.GetLevel(),
					gameScene.Player.GetExperience(), gameScene.Player.GetHitPoints(),
					gameScene.Player.GetMaxHitPoints()),
			}
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
}

// newHUD configures and returns a HUD system.
func newHUD() (*HUD, error) {
	h := &HUD{BasicEntity: ecs.NewBasic()}
	h.RenderComponent.Drawable = common.Text{
		Font: gameScene.Fonts.HUDFont,
		Text: fmt.Sprintf(hudFormatter, gameScene.Player.GetLevel(),
			gameScene.Player.GetExperience(), gameScene.Player.GetHitPoints(),
			gameScene.Player.GetMaxHitPoints()),
	}
	h.RenderComponent.SetZIndex(zHUD)
	h.SetShader(common.HUDShader)
	h.SpaceComponent = common.SpaceComponent{Position: engo.Point{X: 0, Y: 0}}
	for _, sys := range gameScene.World.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Add(&h.BasicEntity, &h.RenderComponent, &h.SpaceComponent)
		}
	}
	return h, nil
}
