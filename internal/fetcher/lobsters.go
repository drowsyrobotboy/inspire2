package fetcher

import (
	"encoding/json"
	"net/http"

	"github.com/drowsyrobotboy/inspire2/internal/models"
)

func FetchLobsters() ([]models.News, error) {
	resp, err := http.Get("https://lobste.rs/hottest.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var items []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, err
	}

	var news []models.News
	for _, item := range items {
		news = append(news, models.News{Title: item.Title, URL: item.URL})
	}

	return news, nil
}
