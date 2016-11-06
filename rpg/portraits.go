package rpg

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Portrait struct {
	Name string

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewPortrait(name string, sprite int, x, y float32) *Portrait {
	p := &Portrait{
		Name:        name,
		BasicEntity: ecs.NewBasic(),
	}
	p.RenderComponent = common.RenderComponent{
		Drawable: gameSprites.Characters.Cell(sprite),
		Scale:    engo.Point{charScale, charScale},
	}
	p.RenderComponent.SetZIndex(1)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{x, y},
		Width:    charSizeX,
		Height:   charSizeY,
	}
	return p
}

type CharSelectPanel struct {
	Portraits []*Portrait

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewCharSelectPanel() *CharSelectPanel {
	c := &CharSelectPanel{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Drawable: common.Rectangle{},
			Color:    color.RGBA{35, 35, 35, 255},
		},
	}
	c.RenderComponent.SetZIndex(0)
	c.SetShader(common.HUDShader)
	c.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{120, 280},
		Width:    490,
		Height:   80,
	}
	c.AddPortraits()
	return c
}

func (c *CharSelectPanel) AddPortraits() {
	c.Portraits = append(c.Portraits, NewPortrait("Human", spriteHuman,
		134, 284))
	c.Portraits = append(c.Portraits, NewPortrait("Zombie", spriteGreenZombie,
		216, 284))
	c.Portraits = append(c.Portraits, NewPortrait("Zombie", spriteOliveZombie,
		296, 284))
	c.Portraits = append(c.Portraits, NewPortrait("Zombie", spriteWhiteZombie,
		376, 284))
	c.Portraits = append(c.Portraits, NewPortrait("Mummy", spriteMummy,
		456, 284))
	c.Portraits = append(c.Portraits, NewPortrait("Gargoyle", spriteGargoyle,
		536, 284))
}
