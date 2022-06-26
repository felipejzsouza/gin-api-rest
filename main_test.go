package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/felipejzsouza/gin-api-rest/controllers"
	"github.com/felipejzsouza/gin-api-rest/database"
	"github.com/felipejzsouza/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func ConfigurarRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do aluno teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletarAlunoMock() {
	aluno := models.Aluno{}
	database.DB.Delete(&aluno, ID)
}

func TestVerificacaoStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := ConfigurarRotasTeste()
	r.GET("/saudacao/:nome", controllers.SaudarAluno)
	req, _ := http.NewRequest("GET", "/saudacao/Felipe", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Status Error: O status de retorno não condiz com o esperado")
	mockDaResposta := `{"API diz":"Eai Felipe, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectarDB()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := ConfigurarRotasTeste()
	r.GET("/alunos", controllers.BuscarAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TesBuscarAlunoPorCPF(t *testing.T) {
	database.ConectarDB()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := ConfigurarRotasTeste()
	r.GET("/aluno/cpf/:cpf", controllers.BuscarAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/aluno/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscarAlunoHandler(t *testing.T) {
	database.ConectarDB()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := ConfigurarRotasTeste()
	r.GET("/aluno/:id", controllers.BuscarAlunoPorCPF)
	pathBusca := "/aluno/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do aluno teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeletarAlunoHandler(t *testing.T) {
	database.ConectarDB()
	CriarAlunoMock()
	r := ConfigurarRotasTeste()
	r.DELETE("/aluno/:id", controllers.DeletarAluno)
	pathChamada := "/aluno/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathChamada, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditarAlunoHandler(t *testing.T) {
	database.ConectarDB()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := ConfigurarRotasTeste()
	r.PATCH("/aluno/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Nome do aluno teste", CPF: "47123456789", RG: "123456700"}
	alunoJson, _ := json.Marshal(aluno)
	pathChamada := "/aluno/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathChamada, bytes.NewBuffer(alunoJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoResposta models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoResposta)
	assert.Equal(t, http.StatusOK, resposta.Code)
	assert.Equal(t, "Nome do aluno teste", alunoResposta.Nome)
	assert.Equal(t, "47123456789", alunoResposta.CPF)
	assert.Equal(t, "123456700", alunoResposta.RG)
}
