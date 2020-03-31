package api

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("init()")
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
