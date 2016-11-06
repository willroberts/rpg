package rpg

import "fmt"

// Given pointers to two entities, resolve combat between them. Applies damage
// simultaneously and then updates the HUD before destroying any Characters.
func initiateCombat(c1, c2 Character) {
	// Apply damage values simultaneously.
	c1.ModifyHitPoints(-c2.GetDamage())
	c2.ModifyHitPoints(-c1.GetDamage())

	// Write damage dealt to the activity log and update the HUD.
	gameLog.Update(fmt.Sprintf("%s hits %s for %d damage!", c1.GetName(),
		c2.GetName(), c1.GetDamage()))
	gameLog.Update(fmt.Sprintf("%s hits %s for %d damage!", c2.GetName(),
		c1.GetName(), c2.GetDamage()))
	gameHUD.Update()

	// Destroy any entities with no HP remaining, and determine if the Player died.
	var playerDestroyed bool
	var enemyDestroyed bool
	var xpBonus int

	if c1.GetHitPoints() == 0 {
		c1.Destroy()
		if c1.GetName() == gamePlayer.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = c1.GetXPBonus()
		}
	}
	if c2.GetHitPoints() == 0 {
		c2.Destroy()
		if c2.GetName() == gamePlayer.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = c2.GetXPBonus()
		}
	}

	// If the Player was not destroyed, grant Experience Points and update the HUD.
	if enemyDestroyed && !playerDestroyed {
		gameLog.Update(fmt.Sprintf("You gained %d experience points.", xpBonus))
		gamePlayer.ModifyExperience(xpBonus)
		gameHUD.Update()
	}
}
