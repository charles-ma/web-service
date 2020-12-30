package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Look what number you got!\n%d", rand.Intn(100)+1)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
