package gotify

import (
	"bytes"
	"encoding/json"
	"github.com/antonholmquist/jason"
	"log"
	"net/http"
)

type Playlist interface {
	RemoveTrack(t *Track) (err error)
	AddTrack(t *Track) (err error)
}

func CreatePlaylist(name string, public bool, auth Auth) (Playlist, error) {
	playlist := SpotifyPlaylist{}
	playlist.Name = name
	playlist.Public = public

	b, err := json.Marshal(playlist)
	if b != nil {
		log.Fatal(err)
	}

	url := "https://api.spotify.com/v1/users/"
	url = url + auth.OwnerId + "/playlists"

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth.AccessToken)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	json, err := jason.NewObjectFromReader(res.Body)
	defer res.Body.Close()

	nPlaylist := parsePlaylist(json)

	return &nPlaylist, err
}

func GetPlaylist(name string, auth Auth) (Playlist, error) {
	panic("not yet implmented")
}

type SpotifyPlaylist struct {
	Id     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Tracks []Track `json:"-"`
	Public bool    `json:"public"`
	Owner  string  `json:"owner,omitempty"`
	Url    string  `json:"uri,omitempty"`
}

func (p *SpotifyPlaylist) AddTrack(t *Track) error {
	panic("not yet implemented")
}

func (p *SpotifyPlaylist) RemoveTrack(t *Track) error {
	panic("not yet implemented")
}
