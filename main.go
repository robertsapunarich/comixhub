package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	renderJSON(writer, "hello!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/hello", helloHandler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), router))
}
