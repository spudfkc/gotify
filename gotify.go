package gotify

import (
	"golang.org/x/oauth2"
	"net/http"
	// "log"
	"os"
)

type Auth struct {
	OwnerName    string
	OwnerId      string
	ClientId     string
	ClientSecret string
	AccessToken  string
}

type Album struct {
	Name string
	Uri  string
}

type Artist struct {
	Name string
	Uri  string
}

type Track struct {
	Name    string
	Uri     string
	Album   Album
	Artists []Artist
}

type Oauth2Authenticator struct {
	Conf oauth2.Config
}

type SpotifyClient struct {
	Client Client
}

const (
	ENV_CLIENT_ID     = "GOTIFY_CLIENT_ID"
	ENV_CLIENT_SECRET = "GOTIFY_CLIENT_SECRET"
	ENV_REDIRECT_URL  = "GOTIFY_REDIRECT_URL"

	SPOTIFY_AUTH_URL  = "https://accounts.spotify.com/authorize"
	SPOTIFY_TOKEN_URL = "https://accounts.spotify.com/api/token"
)

func CreateAuthenticator() *Oauth2Authenticator {
	conf := oauth2.Config{
		ClientID:     os.Getenv(ENV_CLIENT_ID),
		ClientSecret: os.Getenv(ENV_CLIENT_SECRET),
		RedirectURL:  os.Getenv(ENV_REDIRECT_URL),
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  SPOTIFY_AUTH_URL,
			TokenURL: SPOTIFY_TOKEN_URL,
		},
	}

	return &Oauth2Authenticator{conf}
}

func CreateClient(res *http.Request) {
	params := res.URL.Query()

}

// url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
// log.Println("Visit the URL for auth dialog", url)

// var code string
// if _, err := fmt.Scan(&code); err != nil {
// 	log.Fatal(err)
// }

// token, err := conf.Exchange(oauth2.NoContext, code)
// if err != nil {
// 	log.Fatal(err)
// }

// client := conf.Client(oauth2.NoContext, token)
// client.Get("code")
