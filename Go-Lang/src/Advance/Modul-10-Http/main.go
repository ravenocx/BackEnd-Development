package main

import (
	"net/http"

	"github.com/ritsuhaaa/BackEnd-Development/Go-Lang/src/Advance/Modul-10-Http/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080",srv)
	
}