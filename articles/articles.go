package articles

import "strings"

func Corriger(mots []string) []string {
	for i, mot := range mots {
		if (mot == "a" || mot == "A") && i+1 < len(mots) {
			lettre := mots[i+1][:1]
			if strings.ContainsAny(lettre, "aeiouhAEIOUH") {
				mots[i] = mot + "n"
			}
		}
	}
	return mots
}
