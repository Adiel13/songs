package logic

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Definir un modelo para la tabla
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

func InsertSong(s song) {
	dsn := "root:songs@tcp(db_songs:3306)/songs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	newSong := TransaccionSong{
		TrackID:    s.id,
		NombreSong: s.name,
		Artist:     s.artist,
		Duration:   s.duration,
		Album:      s.album,
		URLArtWork: s.artwork,
		Price:      0.0,
		Origin:     s.origin,
		Fuente:     s.fuente,
		Fecha:      time.Now(),
	}
	result := db.Create(&newSong)
	if result.Error != nil {
		panic(result.Error)
	}

	if result.RowsAffected > 0 {
	} else {
		println("No se insertó ninguna fila")
	}

}
