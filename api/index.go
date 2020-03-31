package api

import (
	"fmt"
	"net/http"
)

var Str string

func init() {
	fmt.Println("init()")
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>"+Str)
}
