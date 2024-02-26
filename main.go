package main

import (
	"fmt"

	"songs/logic"
)

func main() {

	resultSongs := logic.ClientSoapSong()
	resultApple := logic.ClientRestSongs()

	resultSongs = append(resultSongs, resultApple...)
	fmt.Println(resultSongs)

}
