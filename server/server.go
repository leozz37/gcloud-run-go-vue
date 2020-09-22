package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", homePageHandler)

	port := os.Args[1]
	fmt.Println("Server listening on port " + port)
	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
