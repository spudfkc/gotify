package gotify

import (
	"fmt"
	"testing"
)

func TestSearchTrack(t *testing.T) {
	trackName := "Spun"
	json := SearchTrack(trackName)
	fmt.Println("%s", json)
}

func TestParseTrack(t *testing.T) {
}

func TestParseArtist(t *testing.T) {
}

func TestParseAlbum(t *testing.T) {
}
