package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"encoding/csv"
	"encoding/json"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	parsedTemplates, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Print("Error occurred while executing the template or writing its output: ", err)
		return
	}
}

func writeJSONToCSV(data []byte, w http.ResponseWriter) {

	var entries BitacoraEntries
	err := json.Unmarshal(data, &entries)
	if err != nil {
		fmt.Println(err)
	}
	csvWriter := csv.NewWriter(w)

	w.Header().Set("Content-Disposition", "attachment; filename="+entries.Name+".csv")

	csvWriter.Write([]string{"name", "description", "date"})
	for _, entry := range entries.Entries {
		var record []string
		record = append(record, entries.Name)
		record = append(record, strings.Replace(entry.Description, ",", ".", -1))
		record = append(record, entry.Date)
		csvWriter.Write(record)
	}
	csvWriter.Flush()
}

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

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	writeJSONToCSV(body, w)
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
