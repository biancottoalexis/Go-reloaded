package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	texte := string(data)
	mots := strings.Fields(texte)
	for i, mot := range mots {
		if mot == "(hex)" {
			mots[i] = ""
			nombre, err := strconv.ParseInt(mots[i-1], 16, 64)
			if err == nil {
				mots[i-1] = strconv.FormatInt(nombre, 10)
			}
		}
		if mot == "(bin)" {
			mots[i] = ""
			nombre, err := strconv.ParseInt(mots[i-1], 2, 64)
			if err == nil {
				mots[i-1] = strconv.FormatInt(nombre, 10)
			}
		}
		if mot == "(up)" {
			mots[i] = ""
			mots[i-1] = strings.ToUpper(mots[i-1])
		}
		if mot == "(low)" {
			mots[i] = ""
			mots[i-1] = strings.ToLower(mots[i-1])
		}
		if mot == "(cap)" {
			mots[i] = ""
			mots[i-1] = strings.ToUpper(mots[i-1][:1]) + strings.ToLower(mots[i-1][1:])
		}
		if mot == "(up," {
			textnombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(textnombre)
			if err == nil {
				mots[i] = ""
				mots[i+1] = ""
				for j := 1; j <= n && i-j >= 0; j++ {
					mots[i-j] = strings.ToUpper(mots[i-j])
				}
			}
		}
		if mot == "(low," {
			textnombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(textnombre)
			if err == nil {
				mots[i] = ""
				mots[i+1] = ""
				for j := 1; j <= n && i-j >= 0; j++ {
					mots[i-j] = strings.ToLower(mots[i-j])
				}
			}
		}
		if mot == "(cap," {
			textnombre := strings.TrimSuffix(mots[i+1], ")")
			n, err := strconv.Atoi(textnombre)
			if err == nil {
				mots[i] = ""
				mots[i+1] = ""
				for j := 1; j <= n && i-j >= 0; j++ {
					if mots[i-j] != "" {
						mots[i-j] = strings.ToUpper(mots[i-j][:1]) + strings.ToLower(mots[i-j][1:])
					}
				}
			}
		}
	}

	var propre []string
	for _, mot := range mots {
		if mot != "" {
			propre = append(propre, mot)
		}
	}

	resultat := strings.Join(propre, " ")
	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
