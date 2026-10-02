package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello this is Jami and the server is running\n")
	fmt.Println(r.Method)
	if(r.Method == "POST"){
		fmt.Fprint(w, "POST request received")
	
	}else {
		fmt.Fprint(w, "Only POST requests allowed")
	}
}

func main() {

	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/shorten",handler)
	http.ListenAndServe(":8080", nil)
}

