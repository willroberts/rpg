package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type CombatSystem struct{}

// FIXME: Player is moved to the tile under the enemy on attack.
// FIXME: Cancel movement when collision is detected.
// FIXME: Or disallow movement by keeping track of map coordinates.
func (c *CombatSystem) New(*ecs.World) {
	engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		_, isCollision := message.(common.CollisionMessage)
		if isCollision {
			log.Println("collision detected")
		}
	})
}

func (c *CombatSystem) Remove(ecs.BasicEntity) {}

func (c *CombatSystem) Update(float32) {}
