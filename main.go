// @title Go + Gin API
// @version 1.0
// @description Exemplo de APIs REST para a aprendizado.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
package main

import (
	"github.com/felipejzsouza/gin-api-rest/database"
	"github.com/felipejzsouza/gin-api-rest/routes"
)

func main() {
	database.ConectarDB()
	routes.HandleRquests()
}
