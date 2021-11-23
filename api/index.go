package persona

import (
	"fmt"
	"net/http"
)

// Handler is Vercelにデプロイした時に「/api」でここが呼ばれる
func Handler(w http.ResponseWriter, r *http.Request) {

	myURL := r.URL.Path
	startPage := "<h1>Hey from Go!</h1>" + myURL
	fmt.Fprintf(w, startPage)

	if myURL == "/smth/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
