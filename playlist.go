package gotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/antonholmquist/jason"
	"log"
	"net/http"
	"strconv"
)

type Playlist interface {
	RemoveTrack(t *Track) (err error)
	AddTrack(t *Track) (err error)
}

func CreatePlaylist(name string, public bool, auth Auth) (Playlist, error) {
	log.Println("creating playlist")
	playlist := SpotifyPlaylist{}
	playlist.Name = name
	playlist.Public = public

	b, err := json.Marshal(playlist)

	// FIXME why is this failing with <nil> ???
	// if b != nil {
	// 	log.Fatal("failed to marshal playlist", err)
	// }

	url := "https://api.spotify.com/v1/users/"
	url = url + auth.OwnerId + "/playlists"

	client := &http.Client{}

	log.Println("building request")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth.AccessToken)

	log.Println("doing request")
	res, err := client.Do(req)

	if err != nil {
		log.Fatal("http request failed", err)
	}

	log.Println("reading json")

	json, err := jason.NewObjectFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var nPlaylist SpotifyPlaylist

	errorJson, err := json.GetObject("error")
	if err == nil {
		message, _ := errorJson.GetString("message")
		status, _ := errorJson.GetInt64("status")
		err = errors.New(strconv.FormatInt(status, 10) + ": " + message)
	} else {
		nPlaylist = parsePlaylist(json)
	}

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
