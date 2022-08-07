package main

import (
	"RoadToTribal2.0/internal/adaptors/api"
	"log"
	"net/http"
)

func main() {
	r := api.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
