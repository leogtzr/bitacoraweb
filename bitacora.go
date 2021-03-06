/**
leogtzr | leogutierrezramirez@gmail.com
*/
package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func addRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}
	return router
}

func authorize(handler http.Handler, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(bitacoraUser)) != 1 ||
			subtle.ConstantTimeCompare([]byte(pass), []byte(bitacoraPassword)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func init() {
	// check if the necessary env variables are set:
	if user, isSet := os.LookupEnv(userEnvVar); isSet {
		bitacoraUser = user
	} else {
		log.Fatalf("%s env variable not set.", userEnvVar)
	}
	if password, isSet := os.LookupEnv(passwordEnvVar); isSet {
		bitacoraPassword = password
	} else {
		log.Fatalf("%s env variable not set.", passwordEnvVar)
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(false)
	router = addRoutes(router)

	router.HandleFunc("/", authorize(&templateHandler{filename: "index.html"}, enterYourUserNamePassword))
	router.HandleFunc("/allentries.html", authorize(&templateHandler{filename: "allentries.html"}, enterYourUserNamePassword))
	router.HandleFunc("/entries.html", authorize(&templateHandler{filename: "entries.html"}, enterYourUserNamePassword))
	router.HandleFunc("/export", authorize(http.HandlerFunc(exportEntriesByEmployee), enterYourUserNamePassword))

	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	err := http.ListenAndServe(connHost+":"+connPort, router)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}

}
