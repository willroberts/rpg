package grid

import "github.com/willroberts/rpg/internal/char"

// HandleCombat resolves combat between the Player and an Enemy. Damage is
// applied simultaneously, so both can die in one round. Updates the HUD and
// destroys Characters.
func (g *grid) HandleCombat(target char.Character) {
	//player := &Player{} // FIXME

	// Apply damage.

	// Write damage to log and update HUD.

	// Detect game over.

	// Clean up dead Characters.

	// Grant experience points on kill.
	//player.AddExperiencePoints(target.GetExperiencePoints())
}
