package handler

import (
	"fmt"
	"net/http"
	"time"
)

var test Test

func init() {
	test = Test{}
}

type Test struct{}

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}

func main() {}
