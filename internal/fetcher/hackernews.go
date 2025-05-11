package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drowsyrobotboy/inspire2/internal/models"
)

func FetchHackerNews() ([]models.News, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int
	if err := json.NewDecoder(resp.Body).Decode(&ids); err != nil {
		return nil, err
	}

	var news []models.News
	for _, id := range ids[:10] { // Fetch top 10 stories
		itemURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
		itemResp, err := http.Get(itemURL)
		if err != nil {
			continue
		}
		defer itemResp.Body.Close()

		var item struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		}
		if err := json.NewDecoder(itemResp.Body).Decode(&item); err != nil {
			continue
		}
		news = append(news, models.News{Title: item.Title, URL: item.URL})
	}

	return news, nil
}
