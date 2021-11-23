package main

import (
	"fmt"
	"net/http"
	"os"

	persona "github.com/k-p5w/go-persona/api"
)

func main() {
	var ci persona.CardInfo

	ci.Name = "main()"
	// vercel だとgo runされないっぽいな、wasm的な感じをつくればいいのかなぁ
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/api", persona.Handler)
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "9999"
	}
	// 公開する
	http.ListenAndServe(":"+port, nil)
}
