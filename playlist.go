package gotify

// import (
// 	"encoding/json"
// )

type Playlist interface {
	RemoveTrack(t *Track) (err error)
	AddTrack(t *Track) (err error)
}

func CreatePlaylist(name string, public bool, auth Auth) (Playlist, error) {
	panic("not yet implemented")
}

func GetPlaylist(name string, auth Auth) (Playlist, error) {
	panic("not yet implmented")
}

type SpotifyPlaylist struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Tracks []Track `json:"-"`
	Public bool    `json:"public"`
	Owner  string  `json:"owner"`
	Url    string  `json:"uri"`
}

func (p *SpotifyPlaylist) AddTrack(t *Track) {
	panic("not yet implemented")
}

func (p *SpotifyPlaylist) RemoveTrack(t *Track) {
	panic("not yet implemented")
}
