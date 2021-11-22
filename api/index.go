package persona

import (
	"fmt"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
