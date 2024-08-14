package main

import (
	"fmt"
	"net/http"
)

// 01) Basic Way of running and Go server which opens on port localhost:8080 and prints Hello, world whatsup
func main() {
	//This prints hello world in the console
	//fmt.Println("hello, world") 

	//HandleFunc takes two arguments 1)endpoint 2)a function with http ResponseWriter to send response and Request from the web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, world whatsup")
		if err != nil {
			fmt.Printf("error : %v", err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written %d" , n))
	})

	//starts the server on port and nil because HandlerFunc is doing the task needed to be done here.
	http.ListenAndServe(":8080", nil)
}