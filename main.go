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
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/eg", smthHandler)
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "9999"
	}
	// 公開する
	http.ListenAndServe(":"+port, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome home")
}

func smthHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/smth/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "welcome smth")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
