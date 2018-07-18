package rpg

import (
	"engo.io/engo"
	"engo.io/engo/common"
)

const failedMovementMsg string = "You can't go that way!"

type Player struct {
	Experience      int
	ExperienceTable map[int]int
}

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

func (p *Player) ModifyExperience(amount int) {
	p.Experience += amount
	if p.Experience >= p.ExperienceTable[p.Level+1] {
		p.LevelUp()
	}
}

func (p *Player) LevelUp() {
	gameScene.Log.Update("Level up!")
	p.Level++
	p.Experience = 0
	p.MaxHitPoints += 2
	p.HitPoints = p.MaxHitPoints
}
