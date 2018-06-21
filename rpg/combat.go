package rpg

import "fmt"

// HandleCombat resolves combat between the Player and an Enemy. Applies damage
// simultaneously and then updates the HUD before destroying any Characters.
func (g *GameScene) HandleCombat(enemy Character) {
	// Apply damage values simultaneously.
	g.Player.ModifyHitPoints(-enemy.GetDamage())
	enemy.ModifyHitPoints(-g.Player.GetDamage())

	// Write damage dealt to the activity log and update the HUD.
	gameScene.Log.Update(fmt.Sprintf("%s hits %s for %d damage!", g.Player.GetName(),
		enemy.GetName(), g.Player.GetDamage()))
	gameScene.Log.Update(fmt.Sprintf("%s hits %s for %d damage!", enemy.GetName(),
		g.Player.GetName(), enemy.GetDamage()))
	gameScene.HUD.Update()

	// Destroy any entities with no HP.
	var playerDestroyed bool
	var enemyDestroyed bool
	var xpBonus int

	if g.Player.GetHitPoints() == 0 {
		g.Player.Destroy()
		if g.Player.GetName() == gameScene.Player.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = g.Player.GetXPBonus()
		}
	}
	if enemy.GetHitPoints() == 0 {
		enemy.Destroy()
		if enemy.GetName() == gameScene.Player.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = enemy.GetXPBonus()
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
