package rpg

import (
	"engo.io/engo"
	"engo.io/engo/common"
)

func (p *Player) Destroy() {
	for _, sys := range gameScene.World.Systems() {
		switch s := sys.(type) {
		case *CameraSystem:
			s.Remove(p.BasicEntity)
		case *common.RenderSystem:
			s.Remove(p.BasicEntity)
			p.RenderComponent = common.RenderComponent{
				Drawable: gameScene.Sprites.Decorations.Cell(spriteGravestone),
				Scale:    engo.Point{X: charScale, Y: charScale},
			}
			p.RenderComponent.SetZIndex(zCharacter)
			s.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
		}
	}
}
