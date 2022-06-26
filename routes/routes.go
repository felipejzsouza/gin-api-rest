package routes

import (
	"github.com/felipejzsouza/gin-api-rest/controllers"
	_ "github.com/felipejzsouza/gin-api-rest/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRquests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/saudacao/:nome", controllers.SaudarAluno)
	r.GET("/alunos", controllers.BuscarAlunos)
	r.GET("/aluno/:id", controllers.BuscarAluno)
	r.POST("/aluno", controllers.CadastrarAluno)
	r.DELETE("/aluno/:id", controllers.DeletarAluno)
	r.PATCH("/aluno/:id", controllers.EditarAluno)
	r.GET("/aluno/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	r.GET("/index", controllers.ExibirPagina)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.NoRoute(controllers.RetornarNaoEncontrada)
	r.Run()
}
