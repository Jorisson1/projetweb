package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Draw struct {
	Faute  int
	Lettre []string
}

var Mot string
var Placement []string
var Faute Draw

var Lettre []string

func LettresProposes(guess string) {
	a := 0
	for i := 0; i < len(Lettre); i++ {
		if guess == Lettre[i] {
			fmt.Println("Erreur Vous avez deja selectionner cette lettre")
			Gameplay()
		}
	}
	Lettre = append(Lettre, guess)
	for t := 0; t < len(Faute.Lettre); t++ {

		if guess == Faute.Lettre[t] {
			Placement[t] = guess
		} else {
			a = a + 1
			if a == len(Faute.Lettre) {
				Faute.Faute = Faute.Faute + 1
				break
			}
		}
	}
	Gameplay()
}

func ToLower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		return r + 32
	}
	return r
}

func Motcomplet(guess string) {
	if guess != Mot {
		fmt.Println("Faux")
		Faute.Faute = Faute.Faute + 2

	} else {
		fmt.Println("Vraie")
		for i := 0; i < len(Faute.Lettre); i++ {
			Placement[i] = Faute.Lettre[i]
		}
		fmt.Println(Placement)
		fmt.Println("Tu as trouver le mot tu as gagner")
	}
}

func ReadFileContent(filename string) string {
	var test []string
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}
	fmt.Println("Contenu du fichier :")
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		test = append(test, lines[i])
	}
	mot_random := rand.Intn(len(test))
	return test[mot_random]
}

func Affichage() {
	Faute.Faute = 0
	Solo()
}

func Solo() {
	Mot = ReadFileContent("Jedi.txt")
	lettre_random := rand.Intn(len(Mot))
	for i := 0; i < len(Mot); i++ {
		Placement = append(Placement, "_ ")
		Faute.Lettre = append(Faute.Lettre, string(Mot[i]))
	}
	Placement[lettre_random] = Faute.Lettre[lettre_random]
	Gameplay()
}

func Gameplay() bool {
	if Faute.Faute < 6 {
		return true
	} else {
		return false
	}

}
