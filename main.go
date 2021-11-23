package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// func main() {

// 	// vercel だとgo runされないっぽいな、wasm的な感じをつくればいいのかなぁ
// 	http.Handle("/", http.FileServer(http.Dir("static")))

// 	http.HandleFunc("/start", persona.MyHandler)
// 	http.HandleFunc("/home", homeHandler)
// 	http.HandleFunc("/eg", smthHandler)
// 	port := os.Getenv("PORT")
// 	fmt.Println(port)
// 	if port == "" {
// 		port = "9999"
// 	}
// 	// 公開する
// 	http.ListenAndServe(":"+port, nil)
// }

func addServedHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Served-Date", time.Now().String())
}

func makeRequestHandler(middleware http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware(w, r)

		w.Write([]byte("OK"))
	}
}

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	next := addServedHeader
	http.HandleFunc("/", makeRequestHandler(next))
	log.Fatal(s.ListenAndServe())
}

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
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
