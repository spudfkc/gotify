package gotify

type Auth struct {
	OwnerName   string
	OwnerId     string
	AccessToken string
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

type Authenticator interface {
	// TODO
}

func requestAuthToken() {
	baseUrl := "https://accounts.spotify.com/authorize"
	clientId := "client_id="
	responseType := "response_type=token"
	redirectUri := "redirect_uri="
	state := "state="
	scope := "scope="
	showDialog := "show_dialog="

	// response will redirect user to redirectUri with info in query params:
	// access_token
	// token_type: Bearer
	// expires_in: time in seconds
	// state: value of "state" provided in first request
	//
	// ERROR VALUES (exist if error)
	// error: reason for error
	// state: value of "state" from above
	//
	//
	// can now user the access_token with header:
	// Authorization: Bearer {accessToken}
}
