// hud.go
package rpg

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type TextLabel struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func ConfigureHUD() error {
	fnt := &common.Font{
		URL:  "fonts/Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 64,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	// Create HP display
	hpLabel := TextLabel{BasicEntity: ecs.NewBasic()}
	hpLabel.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: fmt.Sprintf("Hit Points: %d", 10),
	}
	hpLabel.SetShader(common.HUDShader)

	for _, system := range GameWorld.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&hpLabel.BasicEntity,
				&hpLabel.RenderComponent,
				&hpLabel.SpaceComponent)
		}
	}

	return nil
}
