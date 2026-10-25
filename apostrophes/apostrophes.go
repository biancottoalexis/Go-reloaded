package apostrophes

func Corriger(mots []string) []string {
	var resultat []string
	ouverte := false

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		if mot != "'" {
			resultat = append(resultat, mot)
		} else if ouverte == false {
			if i+1 < len(mots) {
				mots[i+1] = "'" + mots[i+1]
			} else {
				resultat = append(resultat, mot)
			}
			ouverte = true
		} else {
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = resultat[dernier] + "'"
			} else {
				resultat = append(resultat, mot)
			}
			ouverte = false
		}
	}

	return resultat
}
