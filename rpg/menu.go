package rpg

import (
	"image/color"
	"log"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type MenuScene struct {
	Width  int
	Height int
}

func (s *MenuScene) Preload() {
	log.Println("preloading menu fonts")
	var err error
	gameFonts, err = PreloadFonts()
	if err != nil {
		panic(err)
	}
	log.Println("preloading menu sprites")
	gameSprites, err = PreloadSprites()
	if err != nil {
		panic(err)
	}
}

func (s *MenuScene) Setup(w *ecs.World) {
	common.SetBackground(color.RGBA{40, 44, 40, 255})
	w.AddSystem(&common.RenderSystem{})

	// Create title label and "Portrait:" label.
	// FIXME: Determine width automatically - via t.SpaceComponent.Width?
	labels := []*Label{
		NewLabel("Create Character", gameFonts.TitleFont, 128, 50),
		NewLabel("Name:", gameFonts.MenuFont, 40, 150),
		NewLabel("Portrait:", gameFonts.MenuFont, 40, 230),
		NewLabel("Attributes:", gameFonts.MenuFont, 40, 380),
		NewLabel("STR: XX", gameFonts.MenuFont, 180, 460),
		NewLabel("DEX: XX", gameFonts.MenuFont, 180, 540),
		NewLabel("INT: XX", gameFonts.MenuFont, 400, 460),
		NewLabel("VIT: XX", gameFonts.MenuFont, 400, 540),
		NewLabel("Start", gameFonts.TitleFont, 280, 620),
	}

	// draw a character image selector
	csp := NewCharSelectPanel()

	// draw a character stats selector

	// render everything
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			for _, l := range labels {
				s.Add(&l.BasicEntity, &l.RenderComponent, &l.SpaceComponent)
			}
			s.Add(&csp.BasicEntity, &csp.RenderComponent, &csp.SpaceComponent)
			for _, p := range csp.Portraits {
				s.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
			}
		}
	}
}

func (s *MenuScene) Type() string {
	return "MenuScene"
}
