package char

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

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
func (p *player) GetDamage() int           { return p.Damage }
func (p *player) GetExperiencePoints() int { return p.experience }
func (p *player) Destroy()                 {}

func NewPlayer(name string, x int, y int, spriteID int) Character {
	p := player{
		BasicEntity: ecs.NewBasic(),

		name:       name,
		x:          x,
		y:          y,
		hitPoints:  20,
		experience: 0,
	}

	p.RenderComponent = common.RenderComponent{
		// Requires spritesheet.
		//Drawable: scene.Sprites.Characters.Cell(spriteID),
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

	return p
}
