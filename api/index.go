package api

import (
	"fmt"
	"net/http"
)

var Str string
var name string
var Mon string

func init() {
	fmt.Println("init()")
	Str = "HoHo"
	name = "Micael"
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nHello "+name)
	fmt.Fprintf(w, "\n'%s'", Str)
	fmt.Fprintf(w, "\n'Mon:%s'", Mon)
	fmt.Fprintf(w, "\n<h1>Hello from Go on Now!</h1>"+Str)
}
