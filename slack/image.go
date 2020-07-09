package slack

import (
	"net/url"
	"os"
)

type splashImage struct {
	Data []struct {
		Type string `json:"type`
		ID   string `json:id`
		URL  string `json:url`
	}
}

func getImage(key string) {
	baseURL, err := url.Parse("https://api.unsplash.com/search/photos?page=1&query=office")
	if err != nil {
		panic(err)
	}

	param := url.Values{}
	param.Add("api_key", os.Getenv("ACCESS_KEY"))
	param.Add("q", key)

	baseURL.RawQuery = param.Encode()

}

// rel="first"
