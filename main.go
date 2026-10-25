package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/apostrophes"
	"go-reloaded/articles"
	"go-reloaded/commandes"
	"go-reloaded/ponctuation"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage : go run . entree.txt sortie.txt")
		os.Exit(1)
	}

	fichierEntree := os.Args[1]
	fichierSortie := os.Args[2]

	contenu, err := os.ReadFile(fichierEntree)
	if err != nil {
		fmt.Println("erreur : impossible de lire", fichierEntree)
		os.Exit(1)
	}

	texte := string(contenu)
	lignes := strings.Split(texte, "\n")

	for i := 0; i < len(lignes); i++ {
		lignes[i] = corrigerLigne(lignes[i])
	}

	resultat := strings.Join(lignes, "\n")

	err = os.WriteFile(fichierSortie, []byte(resultat), 0644)
	if err != nil {
		fmt.Println("erreur : impossible d'écrire", fichierSortie)
		os.Exit(1)
	}
}

func corrigerLigne(ligne string) string {
	mots := strings.Fields(ligne)

	mots = commandes.Appliquer(mots)
	mots = articles.Corriger(mots)
	mots = ponctuation.Corriger(mots)
	mots = apostrophes.Corriger(mots)

	return strings.Join(mots, " ")
}
