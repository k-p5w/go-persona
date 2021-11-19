package main

import (
	"net/http"

	persona "github.com/k-p5w/go-persona/api"
)

func main() {
	http.HandleFunc("/start", persona.Handler)
	port := "8080"
	// 公開する
	http.ListenAndServe(":"+port, nil)
}
