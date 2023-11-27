package main

import (
	"fmt"
	pendu "hangman/Hangman-main"
	"net/http"
	"strings"
	"text/template"
)

type Pendu struct {
	Pseudo      string
	Nberreur    int
	Motatrouver string
}

var (
	pseudo string
	guess  string
)

func main() {
	// Registre le handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/pseudo", pseudoHandler)
	http.HandleFunc("/traitement", traitementHandler)
	http.HandleFunc("/pendu", penduHandler)
	http.HandleFunc("/pendutraitement", pendutraitementHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("asset"))))

	// Start le server sur le port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "page1", nil)
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{"join": join}).ParseFiles("Template/" + tmplName + ".html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func join(s []string, sep string) string {
	return strings.Join(s, sep)
}

func pseudoHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "pseudo", nil)
}

func traitementHandler(w http.ResponseWriter, r *http.Request) {
	pseudo = r.FormValue("pseudo")
	fmt.Println(r.FormValue("pseudo"))
	http.Redirect(w, r, "pendu", http.StatusSeeOther)
}

func penduHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(pseudo, pendu.Faute.Faute, pendu.Placement)
	data := Pendu{
		Pseudo:      pseudo,
		Nberreur:    pendu.Faute.Faute,
		Motatrouver: concat(pendu.Placement),
	}
	renderTemplate(w, "pendu", data)

}

func pendutraitementHandler(w http.ResponseWriter, r *http.Request) {
	guess = r.FormValue("guess")
	if len(guess) >= 2 {
		pendu.Motcomplet(guess)
	} else if len(guess) == 1 {
		pendu.LettresProposes(guess)

	}
	http.Redirect(w, r, "pendu", http.StatusSeeOther)
}

func concat(s []string) string {
	slice := s

	result := strings.Join(slice, " ")
	return result
}
