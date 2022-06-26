package main

import (
	"github.com/felipejzsouza/gin-api-rest/database"
	"github.com/felipejzsouza/gin-api-rest/routes"
)

func main() {
	database.ConectarDB()
	routes.HandleRquests()
}
