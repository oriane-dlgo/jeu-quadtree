package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	// Générer le terrain aléatoire si activé et non encore généré
	if configuration.Global.EnableRandomTerrain && !f.randomFloorGenerated {
		width := configuration.Global.NumTileX * 2
		height := configuration.Global.NumTileY * 2
		soilTypes := 6
		// Génère un terrain aléatoire et le stock dans fullContent et quadtreeContent
		f.GenerateRandomFloor(width, height, soilTypes)
	}

	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY

	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}

	// Appelle la fonction WaterFramecount() pour incrémenter le compteur d'images d'eau
	f.WaterFramecount()

	// Si l'état de l'eau est actif, appelle la fonction d'animation de l'eau
	if f.WaterState {
		f.WaterAnimation()
	}

}

// Parcours la Matrice et change les 5 en 6 et inversement (ce sont les 2 textures de l'animation de l'eau)
func (f Floor) WaterAnimation() {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			// Change les cellules de la matrice contenant 5 en 6
			if f.content[y][x] == 5 {
				f.content[y][x] = 6
				// Change les cellules de la matrice contenant 6 en 5
			} else if f.content[y][x] == 6 {
				f.content[y][x] = 5
			}
		}
	}
}

// Incrémente le compteur d'images d'eau
func (f *Floor) WaterFramecount() {
	f.WaterFrameTotal += 1

	// Réinitialise le compteur et change l'état de l'eau après un certain nombre d'images
	if f.WaterFrameTotal >= WaterFrameInterval {
		f.WaterFrameTotal = 0
		// Alterne l'état de l'eau entre vrai et faux
		if f.WaterState {
			f.WaterState = false
		} else if !f.WaterState {
			f.WaterState = true
		}
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {

	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX + x
			absY := topLeftY + y
			f.content[y][x] = ((absX % 2) + (absY % 2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
//
// la version actuelle recopie fullContent dans content, ce qui n'est pas
// le comportement attendu dans le rendu du projet
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX + x
			absY := topLeftY + y
			if absY >= 0 && absY < len(f.fullContent) && absX >= 0 && absX < len(f.fullContent[absY]) {
				f.content[y][x] = f.fullContent[absY][absX]
			} else {
				f.content[y][x] = -1
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
