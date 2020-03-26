package handler

import (
	"net/http"
	"time"	
	"log"
	"fmt"
	"os"
)

func init() {
	log.Print("log init")
	fmt.Fprint(os.Stdout, "os.Stdout init")
	fmt.Fprint(os.Stderr, "os.Stderr init")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("log Handler")
	fmt.Fprint(os.Stdout, "os.Stdout Handler")
	fmt.Fprint(os.Stderr, "os.Stderr Handler")
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
