package floor

import (
	"os"
	"testing"
)

// TestGenerateRandomFloor vérifie que GenerateRandomFloor génère un contenu valide.
func TestGenerateRandomFloor(t *testing.T) {
	f := Floor{}
	width, height, soilTypes := 10, 10, 6
	f.GenerateRandomFloor(width, height, soilTypes)

	// Vérifie que fullContent n'est pas vide.
	if len(f.fullContent) == 0 || len(f.fullContent[0]) == 0 {
		t.Errorf("fullContent devrait être généré, mais il est vide")
	}

	// Vérifie que chaque case a un type de sol valide.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if f.fullContent[y][x] < 0 || f.fullContent[y][x] >= soilTypes {
				t.Errorf("Type de sol invalide à la position (%d, %d): %d", x, y, f.fullContent[y][x])
			}
		}
	}
}

// TestIsSurroundedByWater vérifie que isSurroundedByWater retourne les résultats attendus.
func TestIsSurroundedByWater(t *testing.T) {
	tests := []struct {
		name       string
		floorArray [][]int
		x, y       int
		expected   bool
	}{
		{"Non-surrounded", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 1, 1, false},
		{"Surrounded by water", [][]int{{5, 5, 5}, {5, 1, 5}, {5, 5, 5}}, 1, 1, true},
		{"Edge case", [][]int{{5, 0}, {0, 5}}, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSurroundedByWater(tt.floorArray, tt.x, tt.y); got != tt.expected {
				t.Errorf("isSurroundedByWater() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

// TestSaveFloor vérifie que savefloor enregistre correctement les données générées dans un fichier.
func TestSaveFloor(t *testing.T) {
	floorArray := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	filename := "testfile.txt"

	// Appelle la fonction savefloor.
	err := savefloor(floorArray, filename)
	if err != nil {
		t.Fatalf("Erreur lors de l'appel à savefloor : %v", err)
	}
	defer os.Remove(filename) // Supprime le fichier de test après le test

	// Lit le fichier et vérifie son contenu.
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Erreur lors de la lecture du fichier généré : %v", err)
	}

	expectedContent := "0 1 2 \n3 4 5 \n6 7 8 \n"
	if string(content) != expectedContent {
		t.Errorf("Contenu du fichier généré incorrect, attendu : %v, obtenu : %v", expectedContent, string(content))
	}
}
