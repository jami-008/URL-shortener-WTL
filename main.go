package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"math/rand"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello this is Jami and the server is running\n")
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
		fmt.Fprintln(w, "Received URL:", req.URL)
		shortCode := generateShortCode()
		fmt.Fprintln(w, "Original URL:", req.URL)
		fmt.Fprintln(w, "Short Code:", shortCode)

	}
}

func generateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		RandomIndex := rand.Intn(len(chars))
		code += string(chars[RandomIndex])
	}
	return code
}

type URLRequest struct {
	URL string `json:"url"`
}

func main() {

	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/shorten", handler)



	http.ListenAndServe(":8080", nil)
}
