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
	panic("not yet implemented")
}

func parseTrack(o *jason.Object) Track {
	trackName, _ := o.GetString("name")
	trackUri, _ := o.GetString("uri")

	jsonAlbum, _ := o.GetObject("album")
	album := parseAlbum(jsonAlbum)

	track := Track{trackName, trackUri, album}

	return track
}

func parseTracks(o *jason.Object) []Track {
	tracks := []Track{}

	jsonTracks, _ := o.GetObjectArray("tracks", "items")
	for _, jsonTrack := range jsonTracks {
		tracks = append(tracks, parseTrack(jsonTrack))
	}

	return tracks
}

func parseAlbums(o *jason.Object) []Album {
	// albums := []Album{}
	panic("not yet implemented")
}

func parseArtists(o *jason.Object) []Artist {
	// artists := []Artist{}
	panic("not yet implemented")
}
