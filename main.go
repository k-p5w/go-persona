package main

import (
	"fmt"
	"net/http"
	"os"

	persona "github.com/k-p5w/go-persona/api"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/start", persona.MyHandler)
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "9999"
	}
	// 公開する
	http.ListenAndServe(":"+port, nil)
}
