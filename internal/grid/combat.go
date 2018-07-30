package grid

import (
	"github.com/willroberts/rpg/internal/char"
)

// HandleCombat resolves combat between the Player and an Enemy. Damage is
// applied simultaneously, so both can die in one round. Updates the HUD and
// destroys Characters.
func (g *grid) HandleCombat(player, enemy char.Character) {
	// Apply damage simultaneously.
	player.DealDamage(enemy.GetDamage())
	enemy.DealDamage(player.GetDamage())

	// Detect game over.
	if player.GetHitPoints() == 0 {
		//g.GameScene.GameOver()
	}

	// Clean up dead Characters.

	// Grant experience points on kill.
	player.AddExperiencePoints(enemy.GetExperiencePoints())
}
