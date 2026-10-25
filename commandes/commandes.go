package commandes

import (
	"strconv"
	"strings"
)

func Appliquer(mots []string) []string {
	for i := 1; i < len(mots); i++ {
		mot := mots[i]

		if mot == "(hex)" {
			mots[i-1] = convertir(mots[i-1], 16)
			mots[i] = ""
		}
		if mot == "(bin)" {
			mots[i-1] = convertir(mots[i-1], 2)
			mots[i] = ""
		}
		if mot == "(up)" {
			mots[i-1] = strings.ToUpper(mots[i-1])
			mots[i] = ""
		}
		if mot == "(low)" {
			mots[i-1] = strings.ToLower(mots[i-1])
			mots[i] = ""
		}
		if mot == "(cap)" {
			mots[i-1] = majuscule(mots[i-1])
			mots[i] = ""
		}

		if mot == "(up," && i+1 < len(mots) {
			n := lireNombre(mots[i+1])
			for j := 1; j <= n && i-j >= 0; j++ {
				mots[i-j] = strings.ToUpper(mots[i-j])
			}
			mots[i] = ""
			mots[i+1] = ""
		}
		if mot == "(low," && i+1 < len(mots) {
			n := lireNombre(mots[i+1])
			for j := 1; j <= n && i-j >= 0; j++ {
				mots[i-j] = strings.ToLower(mots[i-j])
			}
			mots[i] = ""
			mots[i+1] = ""
		}
		if mot == "(cap," && i+1 < len(mots) {
			n := lireNombre(mots[i+1])
			for j := 1; j <= n && i-j >= 0; j++ {
				mots[i-j] = majuscule(mots[i-j])
			}
			mots[i] = ""
			mots[i+1] = ""
		}
	}

	var propre []string
	for _, mot := range mots {
		if mot != "" {
			propre = append(propre, mot)
		}
	}
	return propre
}

func convertir(mot string, base int) string {
	nombre, err := strconv.ParseInt(mot, base, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(nombre, 10)
}

func majuscule(mot string) string {
	if mot == "" {
		return mot
	}
	return strings.ToUpper(mot[:1]) + strings.ToLower(mot[1:])
}

func lireNombre(texte string) int {
	texte = strings.TrimSuffix(texte, ")")
	n, err := strconv.Atoi(texte)
	if err != nil {
		return 0
	}
	return n
}
