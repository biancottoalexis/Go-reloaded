package ponctuation

func Corriger(mots []string) []string {
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		k := 0
		for k < len(mot) && estPonctuation(mot[k]) {
			k = k + 1
		}

		debut := mot[:k]
		reste := mot[k:]

		if debut != "" {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = resultat[dernier] + debut
			} else {
				resultat = append(resultat, debut)
			}
		}

		if reste != "" {
			resultat = append(resultat, reste)
		}
	}

	return resultat
}

func estPonctuation(c byte) bool {
	if c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';' {
		return true
	}
	return false
}
