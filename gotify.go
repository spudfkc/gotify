package gotify

import (
	"github.com/antonholmquist/jason"
	"log"
	"net/http"
)

type Auth struct {
	Owner string
}

type Album struct {
	Name string
	Uri  string
}

type Artist struct {
	Name string
	Uri  string
}

type Track struct {
	Name  string
	Uri   string
	Album Album
}

type SearchResult struct {
	Tracks  []Track
	Albums  []Album
	Artists []Artist
}

func Search(s string) SearchResult {
	panic("not yet implemented")
}

func SearchTrack(s string) []Track {
	url := "https://api.spotify.com/v1/search?type=track"
	url += "&q="
	url += s // FIXME need to encode this URL

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	json, err := jason.NewObjectFromReader(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return parseTracks(json)
}
