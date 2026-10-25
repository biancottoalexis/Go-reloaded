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

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mots := strings.Fields(string(data))

	mots = commandes.Appliquer(mots)
	mots = articles.Corriger(mots)
	mots = ponctuation.Corriger(mots)
	mots = apostrophes.Corriger(mots)

	resultat := strings.Join(mots, " ")
	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
