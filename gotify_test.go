package gotify

import (
	"fmt"
	"testing"
)

func TestSearchTypes(t *testing.T) {
	types := []SearchType{TRACK, ARTIST, ALBUM}
	searchResult := Search("Spun", types)
	fmt.Println("%s", searchResult)
}

func TestSearchTrack(t *testing.T) {
	track := SearchTrack("Spun")
	fmt.Println("%s", track)
}

func TestParseTrack(t *testing.T) {

}

func TestParseArtist(t *testing.T) {
}

func TestParseAlbum(t *testing.T) {
}
