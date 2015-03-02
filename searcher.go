package gotify

import (
	"github.com/antonholmquist/jason"
	"log"
	"net/http"
)

type SearchType string

const (
	TRACK  SearchType = "track"
	ARTIST SearchType = "artist"
	ALBUM  SearchType = "album"
)

const (
	baseUrl string = "https://api.spotify.com/v1/search"
)

type SearchResult struct {
	Tracks  []Track
	Albums  []Album
	Artists []Artist
}

func Search(what string, types []SearchType) SearchResult {
	url := baseUrl
	var typeParam string

	for _, t := range types {
		typeParam = typeParam + "," + string(t)
	}

	typeParam = typeParam[1:]
	log.Println("typeParam", url+typeParam)

	return SearchResult{}
}

func SearchTrack(s string) []Track {
	url := baseUrl
	url += "?type="
	url += string(TRACK)
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

	unparsedTracks, err := json.GetObjectArray("tracks", "items")
	if err != nil {
		log.Fatal(err)
	}

	return parseTracks(unparsedTracks)
}
