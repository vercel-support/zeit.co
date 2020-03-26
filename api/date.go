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
	fmt.Print("fmt.Print init")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("log Handler")
	fmt.Fprint(os.Stdout, "os.Stdout Handler")
	fmt.Fprint(os.Stderr, "os.Stderr Handler")
	fmt.Print("fmt.Print Handler")
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
