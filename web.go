package main

import (
	"html/template"
	"net/http"
	"os"
)

type Car struct {
	Brand     string
	Model     string
	Power     int
	Available bool
}
type Garage struct {
	Name string
	Cars []Car
}

func main() {
	tmpl := template.Must(template.ParseFiles("template/page1.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Garage{
			Name: "le pendu",
			Cars: []Car{
				{Brand: "Audi", Model: "TT", Power: 245, Available: false},
				{Brand: "Lamborghini", Model: "Aventador SVJ", Power: 770, Available: true},
				{Brand: "Ferrari", Model: "F8 Spider", Power: 720, Available: true},
			},
		}
		tmpl.Execute(w, data)
	})
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.ListenAndServe(":80", nil)
}
