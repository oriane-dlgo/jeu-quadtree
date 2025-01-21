package camera

import (
	"log"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour la position de la caméra à chaque pas de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY, camXPos, camYPos int) {
	switch configuration.Global.CameraMode {
	case Static:
		// Appel de la fonction pour une caméra statique
		c.updateStatic()
	case FollowCharacter:
		// Appel de la fonction pour suivre le personnage
		c.updateFollowCharacter(characterPosX, characterPosY)
	case BlockBorder:
		// Appel de la fonction pour bloquer les bords
		c.updateblockborder(characterPosX, characterPosY)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste toujours à la position (0,0).
func (c *Camera) updateStatic() {
	// Caméra statique, ne fait rien car la position reste toujours à (0,0)
}

// updateFollowCharacter est la mise-à-jour d'une caméra qui suit toujours le personnage.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int) {
	// La caméra suit directement les coordonnées du personnage
	c.X = characterPosX
	c.Y = characterPosY
}

// updateblockborder vérifie que le personnage ne soit pas bloqué par des bordures ou des sols non valides.
func (c *Camera) updateblockborder(characterPosX, characterPosY int) {
	// Obtenir le contenu complet de la carte
	fullContent := c.FullContent.GetFullContent()

	// Vérifier si le contenu est valide
	if len(fullContent) == 0 || len(fullContent[0]) == 0 {
		log.Println("Erreur: fullContent est vide ou mal initialisé dans la caméra")
		return
	}

	// Calculer les bords du terrain
	maxCamX := len(fullContent[0]) - configuration.Global.NumTileX
	maxCamY := len(fullContent) - configuration.Global.NumTileY

	// Calculer les marges pour éviter les zones noires à l'écran
	marginX := configuration.Global.ScreenCenterTileX
	marginY := configuration.Global.ScreenCenterTileY

	// Ajuster la position de la caméra en fonction de la position du personnage et des marges
	if characterPosX < marginX {
		c.X = marginX
	} else if characterPosX > maxCamX+marginX {
		c.X = maxCamX + marginX
	} else {
		c.X = characterPosX
	}

	if characterPosY < marginY {
		c.Y = marginY
	} else if characterPosY > maxCamY+marginY {
		c.Y = maxCamY + marginY
	} else {
		c.Y = characterPosY
	}

	// Afficher la position de la caméra avec les bordures dans le log
	log.Printf("Camera position (with boundaries): %d, %d", c.X, c.Y)

	// Calculer les coordonnées du coin supérieur gauche de la caméra
	topLeftX := c.X - configuration.Global.ScreenCenterTileX
	topLeftY := c.Y - configuration.Global.ScreenCenterTileY

	// Vérifier si le personnage est hors des bordures valides
	if topLeftX < 0 || topLeftY < 0 || characterPosX >= len(fullContent[0]) || characterPosY >= len(fullContent) {
		c.updateStatic()
		return
	}

	// Vérifier si le personnage est sur un sol invalide (type 5) ou entouré d'eau (sol de type -1)
	if fullContent[characterPosY][characterPosX] == 5 ||
		fullContent[characterPosY][characterPosX] == -1 ||
		(characterPosY > 0 && fullContent[characterPosY-1][characterPosX] == 5) ||
		(characterPosX < len(fullContent[0])-1 && fullContent[characterPosY][characterPosX+1] == -1) ||
		(characterPosY < len(fullContent)-1 && fullContent[characterPosY+1][characterPosX] == -1) ||
		(characterPosX > 0 && fullContent[characterPosY][characterPosX-1] == -1) {
		c.updateStatic()
		return
	}
}
