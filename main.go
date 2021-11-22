package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/start", persona.MyHandler)
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "9999"
	}
	// 公開する
	http.ListenAndServe(":"+port, nil)
}
