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
	tracks := SearchTrack("Incidental")

	for _, track := range tracks {
		// fmt.Println("Track name:", track.Name)
		for _, artist := range track.Artists {
			fmt.Println("Artists:", artist.Name)
		}
		// fmt.Println("---")
	}
}

func TestParseTrack(t *testing.T) {
}

func TestParseArtist(t *testing.T) {
}

func TestParseAlbum(t *testing.T) {
}
