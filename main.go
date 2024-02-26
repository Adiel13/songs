package main

import (
	"songs/logic"
)

func main() {

	resultSongs := logic.ClientSoapSong()
	resultApple := logic.ClientRestSongs()

	resultSongs = append(resultSongs, resultApple...)

	for _, v := range resultSongs {
		logic.InsertSong(v)
	}
}
