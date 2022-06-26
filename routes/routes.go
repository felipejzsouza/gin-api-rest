package routes

import (
	"github.com/felipejzsouza/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRquests() {
	r := gin.Default()
	r.GET("/alunos", controllers.BuscarAlunos)
	r.GET("/aluno/:id", controllers.BuscarAluno)
	r.POST("/aluno", controllers.CadastrarAluno)
	r.DELETE("/aluno/:id", controllers.DeletarAluno)
	r.PATCH("/aluno/:id", controllers.EditarAluno)
	r.GET("/aluno/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	r.Run()
}
