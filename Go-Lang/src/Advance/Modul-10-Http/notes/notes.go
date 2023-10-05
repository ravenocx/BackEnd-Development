package main

import (
	"net/http"

)

func main(){
	// create new route handler
	http.HandleFunc("/hello-world", func (w http.ResponseWriter, r *http.Request){
		// r -> contain all information about the request
		// w -> response
		w.Write([]byte("Hello World"))
		
	})
	http.ListenAndServe(":8080",nil)// initialize http server -> 1st parameter port
	// for launch default http server on that port
	// run in terminal -> go to localhost:8080

	
}