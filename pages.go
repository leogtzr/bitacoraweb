package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"encoding/csv"
	"encoding/json"
)

// templateHandler represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
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
