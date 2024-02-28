package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"songs/logic"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func searchSongs(artist string, song string) string {
	resultSongs := logic.ClientSoapSong(artist, song)
	resultApple := logic.ClientRestSongs(artist, song)
	resultSongs = append(resultSongs, resultApple...)

	logic.InsertSong(resultSongs)
	jsonData, err := json.Marshal(resultSongs)
	if err != nil {
		fmt.Println("hubo un error")
	}
	//songs := string(jsonData)
	fmt.Println(jsonData)
	songsString := string(jsonData)
	fmt.Println(songsString)
	return string(jsonData)
}

type Payload struct {
	Artist string `json:"artist"`
	Song   string `json:"song"`
}

func postSerachSongs(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Mensaje recibido:", payload.Artist)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, searchSongs(payload.Artist, payload.Song))
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error cargando el archivo .env: %v", err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/search/song", postSerachSongs).Methods("POST")
	fmt.Println("Servicio levantado")
	log.Fatal(http.ListenAndServe(":8080", router))
}
