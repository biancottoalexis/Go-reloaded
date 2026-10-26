package commandes

import (
	"strconv"
	"strings"
)

func Appliquer(mots []string) []string {
	for i := 1; i < len(mots); i++ {
		mot := mots[i]

		if mot == "(hex)" || mot == "(bin)" || mot == "(up)" || mot == "(low)" || mot == "(cap)" {
			mots[i-1] = changer(mots[i-1], mot)
			mots[i] = ""
		}

		if mot == "(up," || mot == "(low," || mot == "(cap," {
			if i+1 < len(mots) {
				n := lireNombre(mots[i+1])
				for j := 1; j <= n && i-j >= 0; j++ {
					mots[i-j] = changer(mots[i-j], mot)
				}
				mots[i] = ""
				mots[i+1] = ""
			}
		}
	}
	return enleverVides(mots)
}

func changer(mot string, commande string) string {
	if commande == "(hex)" {
		return convertir(mot, 16)
	}
	if commande == "(bin)" {
		return convertir(mot, 2)
	}
	if commande == "(up)" || commande == "(up," {
		return strings.ToUpper(mot)
	}
	if commande == "(low)" || commande == "(low," {
		return strings.ToLower(mot)
	}
	if commande == "(cap)" || commande == "(cap," {
		return majuscule(mot)
	}
	return mot
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

func enleverVides(mots []string) []string {
	var propre []string
	for _, mot := range mots {
		if mot != "" {
			propre = append(propre, mot)
		}
	}
	return propre
}
