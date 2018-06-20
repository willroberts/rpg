package rpg

import "fmt"

// Given pointers to two entities, resolve combat between them. Applies damage
// simultaneously and then updates the HUD before destroying any Characters.
// FIXME: Combat should always assume c1=player. c2 is always enemy.
func handleCombat(c1, c2 Character) {
	// Apply damage values simultaneously.
	c1.ModifyHitPoints(-c2.GetDamage())
	c2.ModifyHitPoints(-c1.GetDamage())

	// Write damage dealt to the activity log and update the HUD.
	gameScene.Log.Update(fmt.Sprintf("%s hits %s for %d damage!", c1.GetName(),
		c2.GetName(), c1.GetDamage()))
	gameScene.Log.Update(fmt.Sprintf("%s hits %s for %d damage!", c2.GetName(),
		c1.GetName(), c2.GetDamage()))
	gameScene.HUD.Update()

	// Destroy any entities with no HP remaining, and determine if the Player died.
	var playerDestroyed bool
	var enemyDestroyed bool
	var xpBonus int

	if c1.GetHitPoints() == 0 {
		c1.Destroy()
		if c1.GetName() == gameScene.Player.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = c1.GetXPBonus()
		}
	}
	if c2.GetHitPoints() == 0 {
		c2.Destroy()
		if c2.GetName() == gameScene.Player.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = c2.GetXPBonus()
		}
	}

	// If the Player was not destroyed, grant Experience Points and update the HUD.
	if enemyDestroyed && !playerDestroyed {
		// FIXME: Add enemy.GetName()
		gameScene.Log.Update("Enemy was destroyed!")
		gameScene.Log.Update(fmt.Sprintf("You gained %d experience points.", xpBonus))
		gameScene.Player.ModifyExperience(xpBonus)
		gameScene.HUD.Update()
	}

	if playerDestroyed {
		gameScene.Log.Update("You died!")
	}
}
