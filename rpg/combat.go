// combat.go

// RPG: A 2D game written in Go, with the engo engine.
// Copyright (C) 2016 Will Roberts
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA

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
	gameHUD.UpdateHealth()

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
			xpBonus = c1.GetXPAmount()
		}
	}
	if c2.GetHitPoints() == 0 {
		c2.Destroy()
		if c2.GetName() == gamePlayer.Name {
			playerDestroyed = true
		} else {
			enemyDestroyed = true
			xpBonus = c2.GetXPAmount()
		}
	}

	// If the Player was not destroyed, grant Experience Points.
	if enemyDestroyed && !playerDestroyed {
		gameLog.Update(fmt.Sprintf("You gained %d experience points.", xpBonus))
	}
}
