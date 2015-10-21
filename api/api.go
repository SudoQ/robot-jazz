package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/chord/{root}/{pattern}", ChordHandler)
	r.HandleFunc("/notes/{notes}", NotesHandler)
	http.ListenAndServe(":8080", r)
}

func ChordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	root := vars["root"]
	pattern := vars["pattern"]
	w.Write([]byte(fmt.Sprintf("%s %s", root, pattern)))
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notes := vars["notes"]
	// Parse notes
	w.Write([]byte(fmt.Sprintf("%s", notes)))
}
