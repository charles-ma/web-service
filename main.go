package main

import (
	"fmt"
	"log"
	"net/http"

	"web-service/mymath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	n := mymath.GenerateRandomNumber(100)
	fmt.Fprintf(w, "Look what number you got from Docker!\n%d\n", n)

	if n >= 90 {
		fmt.Fprintf(w, "\nperfect!\n")
	}

	fmt.Fprintf(w, "Your is address is ", r.RemoteAddr, "\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
