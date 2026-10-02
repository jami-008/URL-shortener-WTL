package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello this is Jami and the server is running")
}

func main() {

	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/jami",handler)
	http.ListenAndServe(":8080", nil)
}

