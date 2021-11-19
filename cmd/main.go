package main

import (
	"net/http"
	"os"

	persona "github.com/k-p5w/go-persona/api"
)

func main() {
	http.HandleFunc("/", persona.Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	// 公開する
	http.ListenAndServe(":"+port, nil)
}
