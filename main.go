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

// GetIPLocationResponse will hold the Soap response
type GetSearchLyricResult struct {
	SearchLyricResult string `xml:"SearchLyricResult"`
}

// GetIPLocationResult will
type GetIPLocationResult struct {
	XMLName xml.Name `xml:"GeoIP"`
	Country string   `xml:"Country"`
	State   string   `xml:"State"`
}

var (
	r GetSearchLyricResult
)

func clientSoap() {
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
	fmt.Println(string(res.Body[:]))
	//res.Unmarshal(&r)
	fmt.Printf("%T", res)

	/*// GetIpLocationResult will be a string. We need to parse it to XML
	result := GetIPLocationResult{}
	err = xml.Unmarshal([]byte(r.GetIPLocationResult), &result)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	if result.Country != "US" {
		log.Fatalf("error: %+v", r)
	}

	log.Println("Country: ", result.Country)
	log.Println("State: ", result.State)*/
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
	//fmt.Println(songs)

	clientSoap()
}
