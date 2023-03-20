package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	rawPort := os.Getenv("SERVER_CONTAINER_PORT")
	port := ":" + rawPort

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Println(err.Error())
	}
}
