package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	httpRouter := mux.NewRouter().StrictSlash(false)

	httpRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"response": "test"}
		respBytes, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
		return
	})

	log.Fatal(http.ListenAndServe(":80", httpRouter))
}
