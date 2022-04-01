package main

import (
	"encoding/json"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(map[string]string{"MenssageType": "mensaje", "Mensagge": "Hola", "Data": "Nuevos datos"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
