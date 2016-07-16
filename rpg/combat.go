package rpg

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type CombatSystem struct{}

// FIXME: Player is moved to the tile under the enemy on attack.
func (c *CombatSystem) New(*ecs.World) {
	collisionHandler := func(message engo.Message) {
		msg, isCollision := message.(common.CollisionMessage)
		if isCollision {
			// Create a combat event.
			combatMessage := engo.Message(CombatMessage{
				Initiator: msg.Entity.ID(),
				Target:    msg.To.ID(),
			})
			engo.Mailbox.Dispatch(combatMessage)

			// TODO: Cancel movement.
		}
	}
	engo.Mailbox.Listen("CollisionMessage", collisionHandler)

	combatHandler := func(message engo.Message) {
		if message.Type() == "CombatMessage" {
			msg, _ := message.(CombatMessage)
			if msg.Initiator == playerEntityID {
				log.Println("You attacked!")
			}
			if msg.Target == playerEntityID {
				log.Println("You were hit!")
				// TODO: Deduct hit points.
			}
		}
	}
	engo.Mailbox.Listen("CombatMessage", combatHandler)
}

func (c *CombatSystem) Remove(ecs.BasicEntity) {}

func (c *CombatSystem) Update(float32) {
}

type CombatMessage struct {
	Initiator uint64
	Target    uint64
}

func (c CombatMessage) Type() string {
	return "CombatMessage"
}
