package main

import (
	"log"
	"net/http"
	"os"

	routerFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/router"
)

func main() {
	router := routerFactory.GetRouter()
	rawPort := os.Getenv("SERVER_CONTAINER_PORT")
	port := ":" + rawPort

	routerFactory.BootControllers()

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Println(err.Error())
	}
}
