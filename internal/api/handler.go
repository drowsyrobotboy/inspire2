package api

import (
	"net/http"

	"github.com/drowsyrobotboy/inspire2/internal/fetcher"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func RegisterRoutes(r chi.Router) {
	// Define separate routes for Hacker News and Lobsters
	r.Get("/api/hackernews", getHackerNews)
	r.Get("/api/lobsters", getLobsters)

	// Serve static files
	htmlDir := http.Dir("ui")
	r.Handle("/*", http.StripPrefix("/", http.FileServer(htmlDir)))
}

func getHackerNews(w http.ResponseWriter, r *http.Request) {
	hackerNews, err := fetcher.FetchHackerNews()
	if err != nil {
		http.Error(w, "Failed to fetch Hacker News", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, hackerNews)
}

func getLobsters(w http.ResponseWriter, r *http.Request) {
	lobsters, err := fetcher.FetchLobsters()
	if err != nil {
		http.Error(w, "Failed to fetch Lobsters", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, lobsters)
}
