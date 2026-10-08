package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello this is Jami and the server is running\n")
	fmt.Println(r.Method)
	if(r.Method == "POST"){
		var req URLRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Fprintf(w, "Invalid JSON")
			return
		}

		fmt.Fprintln(w, "Received URL:", req.URL)

	
	}
}

type URLRequest struct {

	URL string `json:"url"`

}

func main() {

	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/shorten",handler)
	http.ListenAndServe(":8080", nil)
}

