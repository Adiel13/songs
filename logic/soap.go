package logic

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tiaguinho/gosoap"
)

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

func ClientSoapSong(artist string, track string) []song {
	songs := []song{}

	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}

	gosoap.SetCustomEnvelope("soapenv", map[string]string{
		"xmlns:soapenv": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:tem":     "http://tempuri.org/",
	})

	soap, err := gosoap.SoapClient("http://api.chartlyrics.com/apiv1.asmx?wsdl", httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.Params{
		"artist": artist,
		"song":   track,
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
	for _, v := range response.SearchLyricResult {
		singleSong = song{
			Id:       strconv.FormatInt(int64(v.TrackId), 10),
			Name:     v.Song,
			Artist:   v.Artist,
			Duration: "0:00",
			Album:    "",
			Artwork:  v.ArtistUrl,
			Price:    "",
			Origin:   "",
			Fuente:   2,
		}
		songs = append(songs, singleSong)
	}
	return songs
}
