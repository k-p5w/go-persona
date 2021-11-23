package main

import (
	"fmt"
	"net/http"
	"os"

	persona "github.com/k-p5w/go-persona/api"
	"github.com/k-p5w/go-persona/api/corelogic"
)

func main() {
	var ci corelogic.CardInfo

	ci.Name = "main()"
	// vercel だとgo runされないっぽいな、wasm的な感じをつくればいいのかなぁ
	http.Handle("/data", http.FileServer(http.Dir("static")))

	http.HandleFunc("/start", persona.Handler)
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "9999"
	}
	// 公開する
	http.ListenAndServe(":"+port, nil)
}
