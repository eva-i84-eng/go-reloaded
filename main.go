package main

import (
	"fmt"
	"os"
	"strings"
)

// lecture (Entrées)
func readInput(path string) (string, error) {
	donnees, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(donnees), nil
}

// traitement global
func processContent(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		lignes[i] = Process(ligne)
	}
	return strings.Join(lignes, "\n")
}

// écriture (Sorties)
func writeOutput(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// fonction main qui orchestre le scénario
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	// 1 : Lecture
	texte, err := readInput(inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur lors de la lecture du fichier:", err)
		os.Exit(1)
	}

	// 2 : Traitement
	texteFinal := processContent(texte)

	// 3 : Écriture
	err = writeOutput(outputPath, texteFinal)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur lors de l'écriture du fichier de sortie:", err)
		os.Exit(1)
	}
}