package api

import (
	"fmt"
	"net/http"
)

var Str string
var name string

func init() {
	fmt.Println("init()")
	Str = "HoHo"
	name = "Micael"
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello "+name)
	fmt.Fprintf(w, "'%s'", Str)
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>"+Str)
}
