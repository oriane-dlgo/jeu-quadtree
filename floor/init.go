package floor

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
// hhjb
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	// TODO

	var file *os.File
	var err error

	file, err = os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne != "" {
			var tabval []int

			for _, char := range ligne {
				val, err := strconv.Atoi(string(char))
				if err == nil {
					tabval = append(tabval, val)
				}

			}
			floorContent = append(floorContent, tabval)
		}

	}
	return floorContent
}
