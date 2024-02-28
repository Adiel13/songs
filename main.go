package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"songs/logic"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var tokens map[string]bool

type Request struct {
	Artist string `json:"artist"`
	Song   string `json:"song"`
}

func searchSongs(artist string, song string) string {
	resultSongs := logic.ClientSoapSong(artist, song)
	resultApple := logic.ClientRestSongs(artist, song)
	resultSongs = append(resultSongs, resultApple...)

	x := logic.InsertSong(resultSongs)

	jsonData, err := json.Marshal(x)
	if err != nil {
		fmt.Println("hubo un error")
	}

	return string(jsonData)
}

func postSerachSongs(w http.ResponseWriter, r *http.Request) {

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Token de autorización requerido", http.StatusUnauthorized)
		return
	}

	if !tokens[authToken] {
		http.Error(w, "Token de autorización inválido", http.StatusUnauthorized)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, searchSongs(req.Artist, req.Song))
}

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error cargando el archivo .env: %v", err)
	}

	token := os.Getenv("TOKEN")
	tokens = make(map[string]bool)
	tokens[token] = true

	router := mux.NewRouter()
	router.HandleFunc("/search/song", postSerachSongs).Methods("POST")
	fmt.Println("Servicio songs cargado en el puerto")
	log.Fatal(http.ListenAndServe(":8080", router))
}
