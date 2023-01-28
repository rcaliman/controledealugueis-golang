package controllers

import (
	"aluguel/databases"
	"aluguel/models"
	"aluguel/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func BuscaClientes() []models.Cliente {
	var clientes []models.Cliente
	databases.DB.Preload("Cliente").Preload("Imovel").Find(&clientes)
	return clientes
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "views/index.html", nil)
}

func FormCliente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("editar"))
	var cliente models.Cliente
	if id > 0 {
		databases.DB.First(&cliente, id)
	}
	c.HTML(http.StatusOK, "views/form_cliente.html", gin.H{
		"cliente": cliente,
	})
}

func Clientes(c *gin.Context) {

	clienteParaApagar := strings.Contains(c.Request.URL.String(), "apagar")
	clienteParaInserir := c.Request.Method == "POST" && c.PostForm("id") == "0"
	clienteParaAtualizar := c.Request.Method == "POST" && c.PostForm("id") != "0"

	if clienteParaApagar {
		apagaClientes(c)
	} else if clienteParaInserir {
		insereClientes(c)
	} else if clienteParaAtualizar {
		atualizaClientes(c)
	}

	var clientes []models.Cliente
	databases.DB.Find(&clientes)
	c.HTML(http.StatusOK, "views/clientes.html", gin.H{
		"clientes": clientes,
	})
}

func insereClientes(c *gin.Context) {
	Nome := c.PostForm("nome")
	DataNascimento := utils.TimeParse(c.PostForm("data_nascimento"))
	CI := c.PostForm("ci")
	CPF := c.PostForm("cpf")
	Telefone1 := c.PostForm("telefone_1")
	Telefone2 := c.PostForm("telefone_2")

	cliente := models.Cliente{
		Nome:           Nome,
		DataNascimento: DataNascimento,
		CI:             CI,
		CPF:            CPF,
		Telefone1:      Telefone1,
		Telefone2:      Telefone2,
	}
	databases.DB.Create(&cliente)
}

func atualizaClientes(c *gin.Context) {

	ID, _ := strconv.Atoi(c.PostForm("id"))

	var cliente models.Cliente

	databases.DB.First(&cliente, ID)
	cliente.Nome = c.PostForm("nome")
	cliente.DataNascimento = utils.TimeParse(c.PostForm("data_nascimento"))
	cliente.CI = c.PostForm("ci")
	cliente.CPF = c.PostForm("cpf")
	cliente.Telefone1 = c.PostForm("telefone_1")
	cliente.Telefone2 = c.PostForm("telefone_2")
	databases.DB.Save(cliente)
}

func apagaClientes(c *gin.Context) {
	apagar, _ := strconv.Atoi(c.Query("apagar"))
	var clientes models.Cliente
	databases.DB.Delete(&clientes, apagar)
}
