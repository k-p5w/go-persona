package persona

import (
	"fmt"
	"net/http"

	"github.com/k-p5w/go-persona/api/corelogic"
)

// Handler is Vercelにデプロイした時に「/api」でここが呼ばれる
func Handler(w http.ResponseWriter, r *http.Request) {

	mainD := corelogic.DeckMake()
	myURL := r.URL.Path
	subD := "Sub"
	startPage := "<h1>Hey from Go!</h1>" + mainD.Name + subD
	fmt.Fprintf(w, startPage)

	if myURL != "/api" {
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
