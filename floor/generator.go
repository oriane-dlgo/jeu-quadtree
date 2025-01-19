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

// GenerateRandomFloor génère un terrain aléatoire avec 6 types de sols différents.

func (f *Floor) GenerateRandomFloor(width, height, soilTypes int) {
	randomFloorArray := createRandomFloorArray(width, height, soilTypes)
	f.fullContent = randomFloorArray // Initialisation explicite de fullContent
	log.Printf("fullContent généré avec succès : %dx%d\n", len(f.fullContent), len(f.fullContent[0]))
	f.quadtreeContent = quadtree.MakeFromArray(randomFloorArray)
	f.randomFloorGenerated = true
	log.Printf("fullContent après génération : %v\n", f.fullContent)
	if len(f.fullContent) == 0 {
		log.Println("Erreur : fullContent est vide après la génération")
	}
}

// createRandomFloorArray crée un tableau représentant un terrain aléatoire.
func createRandomFloorArray(width, height, soilTypes int) [][]int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	floorArray := make([][]int, height)
	for y := 0; y < height; y++ {
		floorArray[y] = make([]int, width)
		for x := 0; x < width; x++ {
			floorArray[y][x] = r.Intn(soilTypes)

		}
	}

	if configuration.Global.BlockWater {
		for floorArray[0][0] == 5 || isSurroundedByWater(floorArray, 0, 0) {
			// Regénérer le terrain si nécessaire
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					floorArray[y][x] = r.Intn(soilTypes)
				}
			}
		}
	}

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
	myFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer myFile.Close()

	for _, value := range floorgenerated {
		line := ""
		for _, val := range value {
			line += strconv.Itoa(val) + " "
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = fmt.Fprintln(myFile, line)
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture dans le fichier : %w", err)
		}
	}

	return nil
}

func isSurroundedByWater(floorArray [][]int, x, y int) bool {
	// Vérifier les 4 cases autour de (x, y) (haut, droite, bas, gauche)
	// Prendre en compte les bords du terrain pour éviter des indices hors limites
	return (y > 0 && floorArray[y-1][x] == 5) || // Case du dessus
		(x < len(floorArray[0])-1 && floorArray[y][x+1] == 5) || // Case à droite
		(y < len(floorArray)-1 && floorArray[y+1][x] == 5) || // Case du dessous
		(x > 0 && floorArray[y][x-1] == 5) // Case à gauche
}
