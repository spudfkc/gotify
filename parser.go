package gotify

import (
	"github.com/antonholmquist/jason"
)

func parseAlbum(o *jason.Object) Album {
	albumName, _ := o.GetString("name")
	albumUri, _ := o.GetString("uri")

	album := Album{albumName, albumUri}

	return album
}

func parseArtist(o *jason.Object) Artist {
	artistName, _ := o.GetString("name")
	artistUri, _ := o.GetString("uri")

	artist := Artist{artistName, artistUri}

	return artist
}

func parseTrack(o *jason.Object) Track {
	trackName, _ := o.GetString("name")
	trackUri, _ := o.GetString("uri")

	jsonAlbum, err := o.GetObject("album")

	var album Album
	if err == nil {
		album = parseAlbum(jsonAlbum)
	}

	jsonArtists, err := o.GetObjectArray("artists")

	var artists []Artist
	if err == nil {
		artists = parseArtists(jsonArtists)
	}

	track := Track{trackName, trackUri, album, artists}

	return track
}

func parseTracks(o []*jason.Object) []Track {
	tracks := []Track{}

	for _, jsonTrack := range o {
		tracks = append(tracks, parseTrack(jsonTrack))
	}

	return tracks
}

func parseAlbums(o []*jason.Object) []Album {
	albums := []Album{}

	for _, jsonAlbum := range o {
		albums = append(albums, parseAlbum(jsonAlbum))
	}

	return albums
}

func parseArtists(o []*jason.Object) []Artist {
	artists := []Artist{}

	for _, jsonArtist := range o {
		artists = append(artists, parseArtist(jsonArtist))
	}

	return artists
}
