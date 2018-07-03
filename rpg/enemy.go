package rpg

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// EnemyAttributes stores the game data which are imported from external sources,
// such as JSON or a database. Once a set of attributes is defined, attributes
// may be modified freely without modifying the game code.
type EnemyAttributes struct {
	HitPoints         int `json:"hitpoints" validate:"nonzero"`
	Damage            int `json:"damage" validate:"nonzero"`
	ExperienceGranted int `json:"experience_granted" validate:"nonzero"`
}

// An Enemy is a non-player Character which is hostile by default.
type Enemy struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	Name              string
	Hostile           bool
	HitPoints         int
	Damage            int
	ExperienceGranted int

	X, Y int
}

// Destroy removes an enemy from the Grid and from the RenderSystem.
func (e *Enemy) Destroy() {
	gameScene.Grid.RemoveCharacter(e.GetX(), e.GetY())
	for _, sys := range gameScene.World.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			s.Remove(e.BasicEntity)
		}
	}
}

// GetDamage returns the damage dealt by this Enemy.
func (e *Enemy) GetDamage() int { return e.Damage }

// GetHitPoints returns the current HP for this Enemy.
func (e *Enemy) GetHitPoints() int { return e.HitPoints }

// IsHostile returns the enemy's hostility state.
func (e *Enemy) IsHostile() bool { return e.Hostile }

// GetName returns the name of this Enemy type, which is used by newEnemy to
// retrieve its EnemyAttributes.
func (e *Enemy) GetName() string { return e.Name }

// GetXPBonus returns the number of experience points granted upon killing this
// enemy type.
func (e *Enemy) GetXPBonus() int { return e.ExperienceGranted }

// GetX returns the Enemy's X coordinate.
func (e *Enemy) GetX() int { return e.X }

// GetY returns the Enemey's Y coordinate.
func (e *Enemy) GetY() int { return e.Y }

// ModifyHitPoints adds a number to an Enemy's hit points. To deal damage to an
// Enemy, provide a negative number.
func (e *Enemy) ModifyHitPoints(amount int) {
	e.HitPoints += amount
	if e.HitPoints < 0 {
		e.HitPoints = 0
	}
}

// SetHostility modifies whether or not the enemy is hostile.
func (e *Enemy) SetHostility(h bool) { e.Hostile = h }

// SetX updates the Enemy's X coordinate.
func (e *Enemy) SetX(x int) { e.X = x }

// SetY updates the Enemy's Y coordinate.
func (e *Enemy) SetY(y int) { e.Y = y }

// loadEnemyTypes reads EnemyAttributes from file, and populates the global map
// gameScene.EnemyTypes.
func loadEnemyTypes() error {
	b, err := ioutil.ReadFile("assets/data/enemies.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &gameScene.EnemyTypes); err != nil {
		return err
	}
	return nil
}

// newEnemy creates and returns an Enemy.
func newEnemy(name string, spriteIndex, x, y int) *Enemy {
	e := &Enemy{
		BasicEntity:       ecs.NewBasic(),
		Name:              name,
		Hostile:           true,
		HitPoints:         gameScene.EnemyTypes[name].HitPoints,
		Damage:            gameScene.EnemyTypes[name].Damage,
		ExperienceGranted: gameScene.EnemyTypes[name].ExperienceGranted,
		X:                 x,
		Y:                 y,
	}
	enemyTexture := gameScene.Sprites.Characters.Cell(spriteIndex)
	e.RenderComponent = common.RenderComponent{
		Drawable: enemyTexture,
		Scale:    engo.Point{X: 2, Y: 2},
	}
	e.RenderComponent.SetZIndex(zCharacter)
	e.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: (charSizeX * float32(x)) + charOffsetX,
			Y: (charSizeY * float32(y)) + charOffsetY,
		},
		Width:  charSizeX,
		Height: charSizeY,
	}
	if err := gameScene.Grid.AddOccupant(e, x, y); err != nil {
		log.Println("failed to add occupant to grid:", err)
	}
	return e
}
