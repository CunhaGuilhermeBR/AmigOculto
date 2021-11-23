package main

import (
	"net/http"
	"src/amigOculto/database"
	"src/amigOculto/routes"

	_ "github.com/lib/pq"
)

func main() {
	database.Conect()
	routes.LoadRoutes()
	http.ListenAndServe(":3000", nil)
}
