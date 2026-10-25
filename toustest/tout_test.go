package tests

import (
	"strings"
	"testing"

	"go-reloaded/apostrophes"
	"go-reloaded/articles"
	"go-reloaded/commandes"
	"go-reloaded/ponctuation"
)

func corriger(texte string) string {
	mots := strings.Fields(texte)
	mots = commandes.Appliquer(mots)
	mots = articles.Corriger(mots)
	mots = ponctuation.Corriger(mots)
	mots = apostrophes.Corriger(mots)
	return strings.Join(mots, " ")
}

func verifier(t *testing.T, entree string, attendu string) {
	obtenu := corriger(entree)
	if obtenu != attendu {
		t.Error("obtenu : " + obtenu + " | attendu : " + attendu)
	}
}

func TestHex(t *testing.T) {
	verifier(t, "1E (hex) files were added", "30 files were added")
}

func TestBin(t *testing.T) {
	verifier(t, "It has been 10 (bin) years", "It has been 2 years")
}

func TestUp(t *testing.T) {
	verifier(t, "Ready, set, go (up) !", "Ready, set, GO!")
}

func TestLow(t *testing.T) {
	verifier(t, "I should stop SHOUTING (low)", "I should stop shouting")
}

func TestCap(t *testing.T) {
	verifier(t, "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge")
}

func TestUpAvecNombre(t *testing.T) {
	verifier(t, "This is so exciting (up, 2)", "This is SO EXCITING")
}

func TestArticles(t *testing.T) {
	verifier(t, "There it was. A amazing rock!", "There it was. An amazing rock!")
}

func TestPonctuation(t *testing.T) {
	verifier(t, "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!")
}

func TestApostrophes(t *testing.T) {
	verifier(t, "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'")
}

func TestToutEnsemble(t *testing.T) {
	verifier(t, "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'")
}
