package gotify

type Auth struct {
	Owner string
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
