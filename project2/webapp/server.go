package main

import (
	"log"
	"net/http"

	"github.com/mcnamarabrian/manning-web-application-scaffolding/project2/webapp/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.SetupHandlers(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
