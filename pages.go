package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplates, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Print("Error occurred while executing the template or writing its output: ", err)
		return
	}
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
