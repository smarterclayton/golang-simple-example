package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portStr := os.Getenv("PORT")
	if len(portStr) == 0 {
		portStr = strconv.FormatInt(8080, 10)
	}
	addr := fmt.Sprintf("%s:%s", os.Getenv("ADDR"), portStr)
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hello")
	})); err != nil {
		log.Fatalf("Unable to listen: %v", err)
	}
}
