package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tiaguinho/gosoap"
)

type restApple struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		WrapperType            string    `json:"wrapperType"`
		Kind                   string    `json:"kind,omitempty"`
		ArtistID               int       `json:"artistId"`
		CollectionID           int       `json:"collectionId"`
		TrackID                int       `json:"trackId,omitempty"`
		ArtistName             string    `json:"artistName"`
		CollectionName         string    `json:"collectionName"`
		TrackName              string    `json:"trackName,omitempty"`
		CollectionCensoredName string    `json:"collectionCensoredName"`
		TrackCensoredName      string    `json:"trackCensoredName,omitempty"`
		ArtistViewURL          string    `json:"artistViewUrl"`
		CollectionViewURL      string    `json:"collectionViewUrl"`
		TrackViewURL           string    `json:"trackViewUrl,omitempty"`
		PreviewURL             string    `json:"previewUrl"`
		ArtworkURL30           string    `json:"artworkUrl30,omitempty"`
		ArtworkURL60           string    `json:"artworkUrl60"`
		ArtworkURL100          string    `json:"artworkUrl100"`
		CollectionPrice        float64   `json:"collectionPrice,omitempty"`
		TrackPrice             float64   `json:"trackPrice,omitempty"`
		ReleaseDate            time.Time `json:"releaseDate"`
		CollectionExplicitness string    `json:"collectionExplicitness"`
		TrackExplicitness      string    `json:"trackExplicitness,omitempty"`
		DiscCount              int       `json:"discCount,omitempty"`
		DiscNumber             int       `json:"discNumber,omitempty"`
		TrackCount             int       `json:"trackCount"`
		TrackNumber            int       `json:"trackNumber,omitempty"`
		TrackTimeMillis        int       `json:"trackTimeMillis,omitempty"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		PrimaryGenreName       string    `json:"primaryGenreName"`
		ContentAdvisoryRating  string    `json:"contentAdvisoryRating,omitempty"`
		IsStreamable           bool      `json:"isStreamable,omitempty"`
		CollectionHdPrice      float64   `json:"collectionHdPrice,omitempty"`
		TrackHdPrice           float64   `json:"trackHdPrice,omitempty"`
		ShortDescription       string    `json:"shortDescription,omitempty"`
		LongDescription        string    `json:"longDescription,omitempty"`
		CollectionArtistID     int       `json:"collectionArtistId,omitempty"`
		CollectionArtistName   string    `json:"collectionArtistName,omitempty"`
		Description            string    `json:"description,omitempty"`
	} `json:"results"`
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

type SearchLyricResult struct {
	XMLName       xml.Name `xml:"SearchLyricResult"`
	TrackChecksum string   `xml:"TrackChecksum"`
	TrackId       int      `xml:"TrackId"`
	LyricId       int      `xml:"LyricId"`
	SongUrl       string   `xml:"SongUrl"`
	ArtistUrl     string   `xml:"ArtistUrl"`
	Artist        string   `xml:"Artist"`
	Song          string   `xml:"Song"`
	SongRank      int      `xml:"SongRank"`
}

type SearchLyricResponse struct {
	XMLName           xml.Name            `xml:"SearchLyricResponse"`
	SearchLyricResult []SearchLyricResult `xml:"SearchLyricResult>SearchLyricResult"`
}

var (
	r SearchLyricResult
)

func clientSoap() []song {
	songs := []song{}

	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	// set custom envelope
	gosoap.SetCustomEnvelope("soapenv", map[string]string{
		"xmlns:soapenv": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:tem":     "http://tempuri.org/",
	})

	soap, err := gosoap.SoapClient("http://api.chartlyrics.com/apiv1.asmx?wsdl", httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	// Use gosoap.ArrayParams to support fixed position params
	params := gosoap.Params{
		"artist": "artist",
		"song":   "song",
	}

	res, err := soap.Call("SearchLyric", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}
	res.Unmarshal(&r)

	var response SearchLyricResponse
	err2 := xml.Unmarshal(res.Body, &response)
	if err2 != nil {
		fmt.Println("Error al analizar el XML:", err2)
		return songs
	}
	singleSong := song{}
	// Imprimir los resultados
	for _, v := range response.SearchLyricResult {
		singleSong = song{
			id:       strconv.FormatInt(int64(v.TrackId), 10),
			name:     v.Song,
			artist:   v.Artist,
			duration: "0:00",
			album:    "",
			artwork:  v.ArtistUrl,
			price:    "",
			origin:   "",
		}
		songs = append(songs, singleSong)
	}
	return songs
}

func convertMillisToMinutes(millis int) float64 {
	if millis == 0 {
		return 0
	} else {
		return float64(millis) / 60000
	}
}
func main() {

	songs := []song{}
	//call to apple's api
	response, err := http.Get("https://itunes.apple.com/search?term=Nirvana")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//get response by apple's api
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//convert response to string and convert []byte to json
	s := string(responseData)
	jsonMainApple := restApple{}
	err2 := json.Unmarshal([]byte(s), &jsonMainApple)
	if err2 != nil {
		fmt.Println(err2)
	}

	//create a single song and insert into general slice songs
	singleSong := song{}
	for _, v := range jsonMainApple.Results {
		singleSong = song{
			id:       strconv.FormatInt(int64(v.TrackID), 10),
			name:     v.TrackName,
			artist:   v.ArtistName,
			duration: strconv.FormatFloat(convertMillisToMinutes(v.TrackTimeMillis), 'E', -1, 64),
			album:    v.CollectionName,
			artwork:  v.ArtworkURL100,
			price:    strconv.FormatFloat(v.TrackPrice, 'E', -1, 64),
			origin:   v.Country,
		}
		songs = append(songs, singleSong)
	}
	resultSoap := clientSoap()
	songs = append(songs, resultSoap...)
	fmt.Println(songs)

}
