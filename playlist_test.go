package gotify

import (
	"log"
	"testing"
)

func TestCreatePlaylist(t *testing.T) {
	auth := Auth{"ncc", "spudfkc", ""}
	_, err := CreatePlaylist("gogoplaylist", true, auth)
	if err != nil {
		log.Println("Error while creating playlist", err)
		t.Fail()
	}
}

func TestGetPlaylist(t *testing.T) {

}

func TestAddTrackToPlaylist(t *testing.T) {

}

func TestRemoveTrackFromPlaylist(t *testing.T) {

}
