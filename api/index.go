package api

import (
	"fmt"
	"net/http"
)

var packageApi string
var RootMain string
var RootInit string

func init() {
	packageApi = "init() from package api"
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nPackage Api: "+packageApi)
	fmt.Fprintf(w, "\nRoot main: "+RootMain)
	fmt.Fprintf(w, "\nRoot init: "+RootInit)
}
