package persona

import (
	"fmt"
	"net/http"
)

type CardInfo struct {
	Name string
	Job  string
	Cost int
}

func DeckMake(s string) CardInfo {
	var ci CardInfo

	ci.Name = s
	createSVG(ci)
	return ci
}

// Handler is Vercelにデプロイした時に「/api」でここが呼ばれる
func Handler(w http.ResponseWriter, r *http.Request) {

	myURL := r.URL.Path

	q := r.URL.Query()
	v := q.Get("actor")

	fv := r.FormValue("actorItem")

	mainD := DeckMake(v)

	startPage := "<h1>Hey from Go!</h1>" + mainD.Name + v + fv

	// HTMLを描画
	fmt.Fprintf(w, startPage)

	if myURL != "/api" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
}

// SvgHandler is Vercelにデプロイした時に「/api」でここが呼ばれる
func SvgHandler(w http.ResponseWriter, r *http.Request) {

	myURL := r.URL.Path

	q := r.URL.Query()
	v := q.Get("actor")

	fv := r.FormValue("actorItem")

	mainD := DeckMake(v)

	startPage := "<h1>Hey!! from Go!</h1>" + mainD.Name + v + fv

	// HTMLを描画
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
