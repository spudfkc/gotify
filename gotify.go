package gotify

import (
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id      string
	Name    string
	Product string // premium status?
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
	Conf  *oauth2.Config
	Token *oauth2.Token
}

type Scope string

const (
	PLAYLIST_READ_PRIVATE   Scope = "playlist-read-private"
	PLAYLIST_MODIFY_PRIVATE Scope = "playlist-modify-private"
	PLAYLIST_MODIFY_PUBLIC  Scope = "playlist-modify-public"
	STREAMING               Scope = "streaming"
	USER_FOLLOW_MODIFY      Scope = "user-follow-modify"
	USER_FOLLOW_READ        Scope = "user-follow-read"
	USER_LIBRARY_READ       Scope = "user-library-read"
	USER_LIBRARY_MODIFY     Scope = "user-library-modify"
	USER_READ_PRIVATE       Scope = "user-read-private"
	USER_READ_BIRTHDATE     Scope = "user-read-birthdate"
	USER_READ_EMAIL         Scope = "user-read-email"
)

const (
	ENV_CLIENT_ID     = "GOTIFY_CLIENT_ID"
	ENV_CLIENT_SECRET = "GOTIFY_CLIENT_SECRET"
	ENV_REDIRECT_URL  = "GOTIFY_REDIRECT_URL"
	SPOTIFY_AUTH_URL  = "https://accounts.spotify.com/authorize"
	SPOTIFY_TOKEN_URL = "https://accounts.spotify.com/api/token"
)

var client *http.Client = &http.Client{}
var _auth *Oauth2Authenticator

func CreateAuthenticator(scopes ...Scope) *Oauth2Authenticator {
	scopesStr := make([]string, len(scopes))
	for i, s := range scopes {
		scopesStr[i] = string(s)
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv(ENV_CLIENT_ID),
		ClientSecret: os.Getenv(ENV_CLIENT_SECRET),
		RedirectURL:  os.Getenv(ENV_REDIRECT_URL),
		Scopes:       scopesStr,
		Endpoint: oauth2.Endpoint{
			AuthURL:  SPOTIFY_AUTH_URL,
			TokenURL: SPOTIFY_TOKEN_URL,
		},
	}

	var a Oauth2Authenticator
	a.Conf = conf

	return &a
}

func (auth *Oauth2Authenticator) GetAuthURL() string {
	return auth.Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (auth *Oauth2Authenticator) ParseToken(res *http.Request) {
	params := res.URL.Query()
	code := params.Get("code")

	if code == "" {
		log.Fatal("Unable to retrieve Auth Code from Spotify")
	}

	token, err := auth.Conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal("Failed to get Authorization Token from Spotify", err)
	}

	auth.Token = token
	log.Println("got token ", token)
}

func AuthorizeClient(auth *Oauth2Authenticator) {
	_auth = auth
	client = auth.Conf.Client(oauth2.NoContext, auth.Token)
}
