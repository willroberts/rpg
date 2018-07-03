package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const failedMovementMsg string = "You can't go that way!"

// Player is the player-controlled entity in the game.
type Player struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	CameraComponent

	Name            string
	Hostility       string
	HitPoints       int
	MaxHitPoints    int
	Level           int
	Experience      int
	ExperienceTable map[int]int

	X, Y int
}

// Destroy removes a Player entity from the game. When the player dies, we
// replace the character texture and stop processing incoming movement commands.
// Removes the Player from the CameraSystem and the RenderSystem, and then re-add
// the Player to the RenderSystem as a gravestone.
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

// GetDamage returns the damage dealt by the Player.
func (p *Player) GetDamage() int { return 1 }

// GetHitPoints returns the current HP for the Player.
func (p *Player) GetHitPoints() int { return p.HitPoints }

// GetMaxHitPoints returns the maximum HP for the Player.
func (p *Player) GetMaxHitPoints() int { return p.MaxHitPoints }

// GetLevel returns the current Level of the Player.
func (p *Player) GetLevel() int { return p.Level }

// GetExperience returns the current XP for the Player.
func (p *Player) GetExperience() int { return p.Experience }

// GetHostility returns the demeanor of an enemy or NPC.
func (p *Player) GetHostility() string { return p.Hostility }

// GetName returns the name of the Player.
func (p *Player) GetName() string { return p.Name }

// GetXPBonus is stubbed for Player entities. Only needed for Enemy entities, but
// we include it here to satisfy the Character interface.
func (p *Player) GetXPBonus() int { return 0 }

// GetX returns the Player's X coordinate.
func (p *Player) GetX() int { return p.X }

// GetY returns the Player's Y coordinate.
func (p *Player) GetY() int { return p.Y }

// ModifyHitPoints adds a number to the Player's hit points. To deal damage to
// the Player, provide a negative number.
func (p *Player) ModifyHitPoints(amount int) {
	p.HitPoints += amount
	if p.HitPoints < 0 {
		p.HitPoints = 0
	}
}

// ModifyExperience adds a number to the Player's experience points.
func (p *Player) ModifyExperience(amount int) {
	p.Experience += amount
	if p.Experience >= p.ExperienceTable[p.Level+1] {
		p.LevelUp()
	}
}

// LevelUp increases the Player's level, increases MaxHitPoints, and restores HP.
func (p *Player) LevelUp() {
	gameScene.Log.Update("Level up!")
	p.Level++
	p.Experience = 0
	p.MaxHitPoints += 2
	p.HitPoints = p.MaxHitPoints
}

// SetHostility changes the Player's demeanor.
func (p *Player) SetHostility(h string) { p.Hostility = h }

// SetX updates the Player's X coordinate.
func (p *Player) SetX(x int) { p.X = x }

// SetY updates the Player's Y coordinate.
func (p *Player) SetY(y int) { p.Y = y }

// newPlayer creates and returns a Player.
func newPlayer(name string, spriteIndex, x, y int) *Player {
	p := &Player{
		BasicEntity: ecs.NewBasic(),
		CameraComponent: CameraComponent{
			SchemeHoriz: "horizontal",
			SchemeVert:  "vertical",
		},
		X:               x,
		Y:               y,
		Name:            name,
		Hostility:       "player",
		HitPoints:       20,
		MaxHitPoints:    20,
		Level:           1,
		Experience:      0,
		ExperienceTable: defaultPlayerExperienceTable,
	}
	p.RenderComponent = common.RenderComponent{
		Drawable: gameScene.Sprites.Characters.Cell(spriteIndex),
		Scale:    engo.Point{X: charScale, Y: charScale},
	}
	p.RenderComponent.SetZIndex(zCharacter)
	p.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: (charSizeX * float32(x)) + charOffsetX,
			Y: (charSizeY * float32(y)) + charOffsetY,
		},
		Width:  charSizeX,
		Height: charSizeY,
	}
	if err := gameScene.Grid.AddOccupant(p, x, y); err != nil {
		log.Println("failed to add occupant to grid:", err)
	}
	return p
}
