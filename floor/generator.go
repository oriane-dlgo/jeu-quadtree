package floor

import (
	"math/rand"
	"time"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// GenerateRandomFloor génère un terrain aléatoire avec 5 types de sols différents.
func (f *Floor) GenerateRandomFloor(width, height, soilTypes int) {
	f.quadtreeContent = quadtree.MakeFromArray(createRandomFloorArray(width, height, soilTypes))
	f.randomFloorGenerated = true
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
	return floorArray
}
