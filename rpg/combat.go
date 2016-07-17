// combat.go
package rpg

import "log"

// Given pointers to two entities, resolve combat between them
func InitiateCombat(c1, c2 Character) {
	log.Println("[combat] initiating combat!")

	// Apply damage simultaneously.
	log.Printf("%s hit %s for %d damage!\n", c1.GetType(), c2.GetType(), c2.GetDamage())
	c1.ModifyHitPoints(-c2.GetDamage())

	log.Printf("%s hit %s for %d damage!\n", c2.GetType(), c1.GetType(), c1.GetDamage())
	c2.ModifyHitPoints(-c1.GetDamage())

	// Determine if anyone died.
	if c1.GetHitPoints() <= 0 {
		log.Printf("%s was destroyed!\n", c1.GetType())
		c1.Destroy()
	}
	if c2.GetHitPoints() <= 0 {
		log.Printf("%s was destroyed!\n", c2.GetType())
		c2.Destroy()
	}
}
