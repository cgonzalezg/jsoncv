package main

import (
	"log"
	"net/http"
	"strconv"
)

var port = strconv.Itoa(5000)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening..." + port)
	http.ListenAndServe(":"+port, nil)
}
