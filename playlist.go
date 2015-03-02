package gotify

type Playlist interface {
	RemoveTrack(t *Track) (err error)
	AddTrack(t *Track) (err error)
}

func CreatePlaylist(name string, auth Auth) (Playlist, error) {
	panic("not yet implemented")
}

func GetPlaylist(name string, auth Auth) (Playlist, error) {
	panic("not yet implmented")
}

type SpotifyPlaylist struct {
}

func (p *SpotifyPlaylist) AddTrack(t *Track) {
	panic("not yet implemented")
}

func (p *SpotifyPlaylist) RemoveTrack(t *Track) {
	panic("not yet implemented")
}
