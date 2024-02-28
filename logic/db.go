package logic

import (
	"fmt"
	"os"
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

func InsertSong(songs []song) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbDB := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/songs?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbDB, dbHost, dbPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	for _, s := range songs {
		newSong := TransaccionSong{
			TrackID:    s.Id,
			NombreSong: s.Name,
			Artist:     s.Artist,
			Duration:   s.Duration,
			Album:      s.Album,
			URLArtWork: s.Artwork,
			Price:      0.0,
			Origin:     s.Origin,
			Fuente:     s.Fuente,
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
}
