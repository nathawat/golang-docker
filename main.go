package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Hello, This is SCK GoW project. \n")
		fmt.Fprintf(w, "your envoirment is %q \n", os.Getenv("DBHOST"))
		fmt.Fprintf(w, "your database host is %q \n", os.Getenv("DBHOST"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func Sum(x int, y int) int {
	return x + y
}
