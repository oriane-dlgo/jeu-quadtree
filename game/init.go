package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	if configuration.Global.CameraMode != 2 {

		g.character.Init(g.floor.GetWidth(), g.floor.GetHeight())
		g.camera.Init(g.character.X, g.character.Y)
	}
	if configuration.Global.CameraMode == 2 {

		if configuration.Global.EnableRandomTerrain {
			g.floor.GenerateRandomFloor(configuration.Global.NumTileX*2, configuration.Global.NumTileY*2, 6)

			// Initialisation du personnage après le terrain pour obtenir les dimensions correctes
			g.character.Init(g.floor.GetWidth(), g.floor.GetHeight())

			// Initialisation de la caméra après le personnage pour obtenir les coordonnées initiales correctes
			g.camera = camera.Camera{
				X:           g.character.X,
				Y:           g.character.Y,
				FullContent: g.floor,
			}
		} else {
			// Initialisation du personnage après le terrain pour obtenir les dimensions correctes
			g.character.Init(g.floor.GetWidth(), g.floor.GetHeight())

			// Initialisation de la caméra après le personnage pour obtenir les coordonnées initiales correctes
			g.camera = camera.Camera{
				X:           g.character.X,
				Y:           g.character.Y,
				FullContent: g.floor,
			}

		}
	}

}
