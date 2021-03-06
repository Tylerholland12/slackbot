package slack

import (
	"net/url"
)

type splashImage struct {
	Data []struct {
		Type         string `json:"type"`
		ID           string `json:"id"`
		URL          string `json:"url"`
		Query        string `json:"query"`
		Page         string `json:"page"`
		OrderBy      string `json:"order_by"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		CollectionID string `json:"collection_id"`
		PhotoID      string `json:"photo_id"`
	}
}

func getImage(key string) {
	baseURL, err := url.Parse("https://api.unsplash.com/search/photos?page=1&query=office")
	if err != nil {
		panic(err)
	}

}

// rel="first"
