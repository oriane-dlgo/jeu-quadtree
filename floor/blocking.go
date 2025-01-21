package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	if configuration.Global.BlockWater {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1 || f.content[relativeYPos-1][relativeXPos] == 5 || f.content[relativeYPos-1][relativeXPos] == 6
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1 || f.content[relativeYPos][relativeXPos+1] == 5 || f.content[relativeYPos][relativeXPos+1] == 6
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1 || f.content[relativeYPos+1][relativeXPos] == 5 || f.content[relativeYPos+1][relativeXPos] == 6
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1 || f.content[relativeYPos][relativeXPos-1] == 5 || f.content[relativeYPos][relativeXPos-1] == 6
	}
	if !configuration.Global.BlockWater {
		blocking[0] = relativeYPos <= 0 || f.content[relativeYPos-1][relativeXPos] == -1
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.content[relativeYPos][relativeXPos+1] == -1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.content[relativeYPos+1][relativeXPos] == -1
		blocking[3] = relativeXPos <= 0 || f.content[relativeYPos][relativeXPos-1] == -1
	}

	return blocking
}
