package char

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/willroberts/rpg/internal/camera"
)

type player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	camera.CameraComponent

	name       string
	x          int
	y          int
	hostility  bool
	hitPoints  int
	damage     int
	experience int
}

func (p *player) GetName() string          { return p.name }
func (p *player) GetX() int                { return p.x }
func (p *player) SetX(x int)               { p.x = x }
func (p *player) GetY() int                { return p.y }
func (p *player) SetY(y int)               { p.y = y }
func (p *player) GetHostility() bool       { return p.hostility }
func (p *player) SetHostility(h bool)      { p.hostility = h }
func (p *player) GetHitPoints() int        { return p.hitPoints }
func (p *player) SetHitPoints(hp int)      { p.hitPoints = hp }
func (p *player) GetDamage() int           { return p.damage }
func (p *player) GetExperiencePoints() int { return p.experience }

func (p *player) GetEntity() *ecs.BasicEntity {
	return &p.BasicEntity
}

func (p *player) GetRenderComponent() *common.RenderComponent {
	return &p.RenderComponent
}

func (p *player) GetSpaceComponent() *common.SpaceComponent {
	return &p.SpaceComponent
}

func (p *player) GetCameraComponent() *camera.CameraComponent {
	return &p.CameraComponent
}

func (p *player) Destroy() {
	// Remove the Player from the CameraSystem.
	//for _, sys := range scene.World.Systems() {
	//  switch s := sys.(type) {
	//  case *camera.CameraSystem:
	//    s.Remove(p.BasicEntity)
	//  }
	//}
}

func NewPlayer(name string, x int, y int, sprite common.Drawable) Character {
	p := &player{
		BasicEntity: ecs.NewBasic(),

		name:       name,
		x:          x,
		y:          y,
		hitPoints:  20,
		experience: 0,
	}

	p.RenderComponent = common.RenderComponent{
		Drawable: sprite,
		Scale: engo.Point{
			X: DefaultCharacterScale,
			Y: DefaultCharacterScale,
		},
	}
	p.RenderComponent.SetZIndex(DefaultCharacterZIndex)

	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: DefaultCharacterSize*float32(x) + DefaultCharacterOffsetX,
			Y: DefaultCharacterSize*float32(y) + DefaultCharacterOffsetY,
		},
		Width:  DefaultCharacterSize,
		Height: DefaultCharacterSize,
	}

	p.CameraComponent = camera.CameraComponent{
		SchemeHoriz: "horizontal",
		SchemeVert:  "vertical",
	}

	return p
}
