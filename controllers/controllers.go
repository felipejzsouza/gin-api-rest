package controllers

import (
	"net/http"

	"github.com/felipejzsouza/gin-api-rest/database"
	"github.com/felipejzsouza/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

// @Summary Sauda o chamador com o nome informado
// @ID SaudarAluno
// @Param nome path string true "Nome"
// @Produce json
// @Success 200 {object} string
// @Router /saudacao/{nome} [get]
func SaudarAluno(c *gin.Context) {
	nome := c.Param("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz": "Eai " + nome + ", tudo beleza?",
	})
}

// @Summary Retorna todos os alunos
// @ID BuscarAlunos
// @Produce json
// @Success 200 {object} []models.Aluno
// @Router /alunos [get]
func BuscarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

// @Summary Retorna uma aluno por ID
// @ID BuscarAluno
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /aluno/{id} [get]
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

// @Summary Cadastra um aluno
// @ID CadastrarAluno
// @Param Aluno body models.Aluno true "Aluno"
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /aluno [post]
func CadastrarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if err := models.Validar(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// @Summary Deleta o aluno indicado
// @ID DeletarAluno
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} string
// @Router /aluno/{id} [delete]
func DeletarAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, aluno)
}

// @Summary Edita um aluno
// @ID EditarAluno
// @Param id path string true "ID"
// @Param Aluno body models.Aluno true "Aluno"
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /aluno/{id} [patch]
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
	if err := models.Validar(&aluno); err != nil {
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

func ExibirPagina(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RetornarNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
