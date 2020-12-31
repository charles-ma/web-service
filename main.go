package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(100) + 1
	fmt.Fprintf(w, "Look what number you got from Docker!\n%d\n", n)

	if n > 90 {
		fmt.Fprintf(w, "\nperfect!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
