package gotify

type Client struct {
	BaseUrl string
	params  map[string]string
	suffix  string
}
