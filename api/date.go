package handler

import (
	"fmt"
	"net/http"
	"time"
)

var initFn string
var mainFn string

func init() {
	initFn = "init()"
}

type Test struct{}

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
	fmt.Fprintf(w, initFn)
	fmt.Fprintf(w, mainFn)
}

func main() {
	mainFn = "init()"
}
