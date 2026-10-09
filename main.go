package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		var req URLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Fprintf(w, "Invalid JSON")
			return
		}
		if req.URL == "" {
			fmt.Fprintln(w, "URL cannot be empty")
			return
		}

		shortCode := generateShortCode()
		urlStore[shortCode] = req.URL
		response := URLResponse{
	    OriginalURL: req.URL,
	    ShortURL:    "http://localhost:8080/" + shortCode,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
}
	}


func generateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(chars))
		code += string(chars[randomIndex])
	}
	return code
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {

	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	originalURL, exists := urlStore[shortCode]

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

type URLRequest struct {
	URL string `json:"url"`
}

type URLResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

var urlStore = make(map[string]string)

func main() {

	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/shorten", handler)
	http.HandleFunc("/", redirectHandler)
	http.ListenAndServe(":8080", nil)
}
