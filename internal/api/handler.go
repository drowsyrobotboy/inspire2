package api

import (
	"net/http"

	"github.com/drowsyrobotboy/inspire2/internal/fetcher"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/news", getNews)

	// Serve static files using r.Handle
	htmlDir := http.Dir("ui")
	r.Handle("/*", http.StripPrefix("/", http.FileServer(htmlDir)))
}

func getNews(w http.ResponseWriter, r *http.Request) {
	hackerNews, err := fetcher.FetchHackerNews()
	if err != nil {
		http.Error(w, "Failed to fetch Hacker News", http.StatusInternalServerError)
		return
	}

	lobsters, err := fetcher.FetchLobsters()
	if err != nil {
		http.Error(w, "Failed to fetch Lobsters", http.StatusInternalServerError)
		return
	}

	// Combine results
	allNews := append(hackerNews, lobsters...)

	render.JSON(w, r, allNews)
}
