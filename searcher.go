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
	panic("not yet implemented")
}

func SearchTrack(s string) []Track {
	log.Println("starting...")

	url := baseUrl
	url += "?type="
	url += string(TRACK)
	url += "&q="
	url += s // FIXME need to encode this URL

	client := &http.Client{}
	log.Println("httpclient created")

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	log.Println("req created")

	res, err := client.Do(req)
	log.Println("req sent!")

	if err != nil {
		log.Fatal(err)
	}

	json, err := jason.NewObjectFromReader(res.Body)
	log.Println("got reader from res")
	res.Body.Close()
	log.Println("closed")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("getting tracks...")
	unparsedTracks, err := json.GetObjectArray("tracks", "items")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("parsing response...\n%s", unparsedTracks)
	return parseTracks(unparsedTracks)
}
