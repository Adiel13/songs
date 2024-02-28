package logic

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InsertSong(songs []song) []Result {

	res := []Result{}

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

		duracion, _ := strconv.ParseFloat(s.Duration[:len(s.Duration)-1], 64)
		precio, _ := strconv.ParseFloat(s.Duration[:len(s.Duration)-1], 64)
		y := fmt.Sprintf("%.2f", precio)
		x, _ := strconv.ParseFloat(y, 64)

		newSong := TransaccionSong{
			TrackID:    s.Id,
			NombreSong: s.Name,
			Artist:     s.Artist,
			Duration:   fmt.Sprintf("%.2f", duracion),
			Album:      s.Album,
			URLArtWork: s.Artwork,
			Price:      x,
			Origin:     s.Origin,
			Fuente:     s.Fuente,
			Fecha:      time.Now(),
		}
		font := ""
		if s.Fuente == 1 {
			font = "Apple"
		} else {
			font = "ChartLyrics"
		}

		r := Result{
			Id:       s.Id,
			Name:     s.Name,
			Artist:   s.Artist,
			Duration: fmt.Sprintf("%.2f", duracion),
			Album:    s.Album,
			Artwork:  s.Artwork,
			Price:    fmt.Sprintf("Q. %s", y),
			Origin:   font,
		}

		res = append(res, r)

		result := db.Create(&newSong)
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected > 0 {
		} else {
			println("No se insertó ninguna fila")
		}
	}
	return res
}
