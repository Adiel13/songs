package logic

import "time"

type song struct {
	Id       string
	Name     string
	Artist   string
	Duration string
	Album    string
	Artwork  string
	Price    string
	Origin   string
	Fuente   int
}

type TransaccionSong struct {
	ID         uint   `gorm:"column:id_transaccion_song"`
	TrackID    string `gorm:"column:id_track"`
	NombreSong string `gorm:"column:nombre_song"`
	Artist     string
	Duration   string
	Album      string
	URLArtWork string `gorm:"column:url_art_work"`
	Price      float64
	Origin     string
	Fuente     int `gorm:"column:fuente"`
	Fecha      time.Time
}

type Result struct {
	Id       string
	Name     string
	Artist   string
	Duration string
	Album    string
	Artwork  string
	Price    string
	Origin   string
}
