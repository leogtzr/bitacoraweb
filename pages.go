package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplates, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Print("Error occurred while executing the template or writing its output: ", err)
		return
	}
}

// exportEntriesByEmployee
func exportEntriesByEmployee(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	resp, err := http.Get("http://localhost:" + APIPort + "/api/entries/" + id)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=export.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	io.Copy(w, strings.NewReader(string(body)))
}

func entriesByEmployeePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplates, _ := template.ParseFiles("templates/entries.html")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Print("Error occurred while executing the template or writing its output: ", err)
		return
	}
}

func allEntriesPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplates, _ := template.ParseFiles("templates/allentries.html")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Print("Error occurred while executing the template or writing its output: ", err)
		return
	}
}
