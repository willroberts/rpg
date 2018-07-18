package char

import (
	"encoding/json"
	"io/ioutil"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/willroberts/rpg/internal/camera"
)

type enemy struct {
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

func (e *enemy) GetName() string          { return e.name }
func (e *enemy) GetX() int                { return e.x }
func (e *enemy) SetX(x int)               { e.x = x }
func (e *enemy) GetY() int                { return e.y }
func (e *enemy) SetY(y int)               { e.y = y }
func (e *enemy) GetHostility() bool       { return e.hostility }
func (e *enemy) SetHostility(h bool)      { e.hostility = h }
func (e *enemy) GetHitPoints() int        { return e.hitPoints }
func (e *enemy) SetHitPoints(hp int)      { e.hitPoints = hp }
func (e *enemy) GetDamage() int           { return e.damage }
func (e *enemy) GetExperiencePoints() int { return e.experience }

func (e *enemy) GetEntity() *ecs.BasicEntity {
	return &e.BasicEntity
}

func (e *enemy) GetRenderComponent() *common.RenderComponent {
	return &e.RenderComponent
}

func (e *enemy) GetSpaceComponent() *common.SpaceComponent {
	return &e.SpaceComponent
}

func (e *enemy) GetCameraComponent() *camera.CameraComponent {
	return &e.CameraComponent
}

func (e *enemy) Destroy() {
	// Remove the Enemy from the CameraSystem.
	//for _, sys := range scene.World.Systems() {
	//  switch s := sys.(type) {
	//  case *camera.CameraSystem:
	//    s.Remove(e.BasicEntity)
	//  }
	//}
}

func NewEnemy(name string, x int, y int, attrs *EnemyAttributes, sprite common.Drawable) Character {
	e := &enemy{
		BasicEntity: ecs.NewBasic(),

		name:       name,
		hostility:  true,
		x:          x,
		y:          y,
		hitPoints:  attrs.HitPoints,
		damage:     attrs.Damage,
		experience: attrs.ExperienceGranted,
	}

	e.RenderComponent = common.RenderComponent{
		Drawable: sprite,
		Scale: engo.Point{
			X: DefaultCharacterScale,
			Y: DefaultCharacterScale,
		},
	}
	e.RenderComponent.SetZIndex(DefaultCharacterZIndex)

	e.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: DefaultCharacterSize*float32(x) + DefaultCharacterOffsetX,
			Y: DefaultCharacterSize*float32(y) + DefaultCharacterOffsetY,
		},
		Width:  DefaultCharacterSize,
		Height: DefaultCharacterSize,
	}

	e.CameraComponent = camera.CameraComponent{
		SchemeHoriz: "horizontal",
		SchemeVert:  "vertical",
	}

	return e
}

type EnemyAttributes struct {
	HitPoints         int `json:"hitpoints" validate:"nonzero"`
	Damage            int `json:"damage" validate:"nonzezro"`
	ExperienceGranted int `json:"experience_granted" validate:"nonzero"`
}

func LoadEnemies(filename string) (map[string]*EnemyAttributes, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := make(map[string]*EnemyAttributes, 0)
	if err = json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return data, nil
}
