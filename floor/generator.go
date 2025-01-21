package floor

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// GenerateRandomFloor génère un terrain aléatoire avec plusieurs types de sols différents.
func (f *Floor) GenerateRandomFloor(width, height, soilTypes int) {
	// Crée un tableau représentant un terrain aléatoire.
	randomFloorArray := createRandomFloorArray(width, height, soilTypes)

	// Initialise explicitement fullContent avec le tableau généré aléatoirement.
	f.fullContent = randomFloorArray
	log.Printf("fullContent généré avec succès : %dx%d\n", len(f.fullContent), len(f.fullContent[0]))

	// Génère un quadtree à partir du tableau généré.
	f.quadtreeContent = quadtree.MakeFromArray(randomFloorArray)
	f.randomFloorGenerated = true

	log.Printf("fullContent après génération : %v\n", f.fullContent)

	// Vérifie que fullContent n'est pas vide après la génération.
	if len(f.fullContent) == 0 {
		log.Println("Erreur : fullContent est vide après la génération")
	}
}

// createRandomFloorArray crée un tableau représentant un terrain aléatoire.
func createRandomFloorArray(width, height, soilTypes int) [][]int {
	// Initialise un nouveau générateur de nombres aléatoires.
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Crée le tableau de terrain aléatoire.
	floorArray := make([][]int, height)
	for y := 0; y < height; y++ {
		floorArray[y] = make([]int, width)
		for x := 0; x < width; x++ {
			// Assigne un type de sol aléatoire à chaque position.
			floorArray[y][x] = r.Intn(soilTypes)
		}
	}

	// Si BlockWater est activé, régénère le terrain s'il commence dans l'eau ou est entouré par de l'eau.
	if configuration.Global.BlockWater {
		for floorArray[0][0] == 5 || isSurroundedByWater(floorArray, 0, 0) {
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					floorArray[y][x] = r.Intn(soilTypes)
				}
			}
		}
	}

	// Si SaveFloorGenerated et EnableRandomTerrain sont activés, enregistre le terrain généré.
	if configuration.Global.SaveFloorGenerated && configuration.Global.EnableRandomTerrain {
		filename := filepath.Join("../floor-files", "reussite.txt")
		if err := savefloor(floorArray, filename); err != nil {
			log.Println("Erreur lors de la sauvegarde du terrain généré :", err)
		}
	}
	return floorArray
}

// savefloor enregistre les données générées dans un fichier.
func savefloor(floorgenerated [][]int, filename string) error {
	// Crée un nouveau fichier pour sauvegarder le terrain.
	myFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	// Écrit chaque ligne du tableau dans le fichier.
	for _, value := range floorgenerated {
		line := ""
		for _, val := range value {
			line += strconv.Itoa(val) + " "
		}
		_, err = fmt.Fprintln(myFile, line)
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture dans le fichier : %w", err)
		}
	}

	return nil
}

// isSurroundedByWater vérifie si une position est entourée d'eau.
func isSurroundedByWater(floorArray [][]int, x, y int) bool {
	// Vérifie les cases autour de (x, y) (haut, droite, bas, gauche) tout en prenant en compte les bords du terrain.
	return (y > 0 && floorArray[y-1][x] == 5) || // Case du dessus
		(x < len(floorArray[0])-1 && floorArray[y][x+1] == 5) || // Case à droite
		(y < len(floorArray)-1 && floorArray[y+1][x] == 5) || // Case du dessous
		(x > 0 && floorArray[y][x-1] == 5) // Case à gauche
}
