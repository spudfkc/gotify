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

	url += "?type="
	url += typeParam
	url += "&q="
	url += what

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	json, err := jason.NewObjectFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	result := SearchResult{}

	if contains(types, TRACK) {
		unparsedTracks, err := json.GetObjectArray("tracks", "items")
		if err == nil {
			result.Tracks = parseTracks(unparsedTracks)
		}
	}

	if contains(types, ARTIST) {
		unparsedArtists, err := json.GetObjectArray("artists", "items")
		if err == nil {
			result.Artists = parseArtists(unparsedArtists)
		}
	}

	if contains(types, ALBUM) {
		unparsedAlbums, err := json.GetObjectArray("albums", "items")
		if err == nil {
			result.Albums = parseAlbums(unparsedAlbums)
		}
	}

	return result
}

func SearchTrack(s string) []Track {
	return Search(s, []SearchType{TRACK}).Tracks
}

func SearchArtist(s string) []Artist {
	return Search(s, []SearchType{ARTIST}).Artists
}

func SearchAlbum(s string) []Album {
	return Search(s, []SearchType{ALBUM}).Albums
}

func contains(t []SearchType, s SearchType) bool {
	for _, item := range t {
		if item == s {
			return true
		}
	}
	return false
}
