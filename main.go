package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type restApple struct {
	WrapperType            string    `json:"wrapperType"`
	Kind                   string    `json:"kind"`
	ArtistID               int       `json:"artistId"`
	CollectionID           int       `json:"collectionId"`
	TrackID                int       `json:"trackId"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	TrackName              string    `json:"trackName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	TrackCensoredName      string    `json:"trackCensoredName"`
	CollectionArtistID     int       `json:"collectionArtistId"`
	CollectionArtistName   string    `json:"collectionArtistName"`
	ArtistViewURL          string    `json:"artistViewUrl"`
	CollectionViewURL      string    `json:"collectionViewUrl"`
	TrackViewURL           string    `json:"trackViewUrl"`
	PreviewURL             string    `json:"previewUrl"`
	ArtworkURL30           string    `json:"artworkUrl30"`
	ArtworkURL60           string    `json:"artworkUrl60"`
	ArtworkURL100          string    `json:"artworkUrl100"`
	ReleaseDate            time.Time `json:"releaseDate"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackExplicitness      string    `json:"trackExplicitness"`
	DiscCount              int       `json:"discCount"`
	DiscNumber             int       `json:"discNumber"`
	TrackCount             int       `json:"trackCount"`
	TrackNumber            int       `json:"trackNumber"`
	TrackTimeMillis        int       `json:"trackTimeMillis"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	ContentAdvisoryRating  string    `json:"contentAdvisoryRating"`
	IsStreamable           bool      `json:"isStreamable"`
}
type song struct {
	id       string
	name     string
	artist   string
	duration string
	album    string
	artwork  string
	price    string
	origin   string
}

func main() {
	response, err := http.Get("https://itunes.apple.com/search?term=Guatemala")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(responseData)

	jsonApple := restApple{}
	err2 := json.Unmarshal([]byte(s), &jsonApple)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(jsonApple.ArtistID)

}
