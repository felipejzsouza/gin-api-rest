package controllers

import (
	"net/http"

	"github.com/felipejzsouza/gin-api-rest/database"
	"github.com/felipejzsouza/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func BuscarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func BuscarAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func CadastrarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}

func EditarAluno(c *gin.Context) {
	id := c.Param("id")
	var aluno models.Aluno

	database.DB.Find(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscarAlunoPorCPF(c *gin.Context) {
	cpf := c.Params.ByName("cpf")
	var aluno models.Aluno
	database.DB.Where(models.Aluno{CPF: cpf}).Find(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	c.JSON(http.StatusOK, aluno)
}
