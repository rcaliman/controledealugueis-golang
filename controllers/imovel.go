package controllers

import (
	"aluguel/databases"
	"aluguel/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func insereImovel(c *gin.Context) {
	tipo := c.PostForm("tipo")
	numero := c.PostForm("numero")
	local := c.PostForm("local")
	cliente_id, _ := strconv.Atoi(c.PostForm("cliente_id"))
	valor_aluguel, _ := strconv.ParseFloat(c.PostForm("valor_aluguel"), 64)
	observacao := c.PostForm("observacao")
	dia_base, _ := strconv.Atoi(c.PostForm("dia_base"))

	imovel := models.Imovel{
		Tipo:         tipo,
		Numero:       numero,
		Local:        local,
		ClienteID:    cliente_id,
		ValorAluguel: valor_aluguel,
		Observacao:   observacao,
		DiaBase:      dia_base,
	}
	databases.DB.Create(&imovel)
}
func atualizaImovel(c *gin.Context) {
	ID, _ := strconv.Atoi(c.PostForm("id"))

	var imovel models.Imovel

	databases.DB.First(&imovel, ID)
	imovel.Tipo = c.PostForm("tipo")
	imovel.Numero = c.PostForm("numero")
	imovel.Local = c.PostForm("local")
	imovel.ClienteID, _ = strconv.Atoi(c.PostForm("cliente_id"))
	imovel.ValorAluguel, _ = strconv.ParseFloat(c.PostForm("valor_aluguel"), 64)
	imovel.Observacao = c.PostForm("observacao")
	imovel.DiaBase, _ = strconv.Atoi(c.PostForm("dia_base"))
	databases.DB.Save(&imovel)

}

func apagaImovel(c *gin.Context) {
	apagar, _ := strconv.Atoi(c.Query("apagar"))
	var imovel models.Imovel
	databases.DB.Delete(&imovel, apagar)
}

func Imoveis(c *gin.Context) {

	imovelParaApagar := strings.Contains(c.Request.URL.String(), "apagar")
	imovelParaInserir := c.Request.Method == "POST" && c.PostForm("id") == "0"
	imovelParaAtualizar := c.Request.Method == "POST" && c.PostForm("id") != "0"

	if imovelParaApagar {
		apagaImovel(c)
	} else if imovelParaInserir {
		insereImovel(c)
	} else if imovelParaAtualizar {
		atualizaImovel(c)
	}

	var imoveis []models.Imovel
	databases.DB.Preload("Imovel").Preload("Cliente").Find(&imoveis)
	c.HTML(http.StatusOK, "views/imoveis.html", gin.H{
		"imoveis": imoveis,
	})
}
func FormImovel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("editar"))
	var imovel models.Imovel
	if id > 0 {
		databases.DB.Preload("Cliente").Preload("Imovel").First(&imovel, id)
	}
	clientes := BuscaClientes()
	c.HTML(http.StatusOK, "views/form_imovel.html", gin.H{
		"imovel":   imovel,
		"clientes": clientes,
	})
}

func Recibos(c *gin.Context) {
	c.Request.ParseForm()
	var listaIds []string
	for key, value := range c.Request.PostForm {
		if key != "mes" && key != "ano" {
			listaIds = append(listaIds, value...)
		}
	}

	data := models.Data{
		Mes: c.Request.PostForm["mes"][0],
		Ano: c.Request.PostForm["ano"][0],
	}

	var imoveis []models.Imovel
	databases.DB.Preload("Cliente").Preload("Imovel").Where("id in ?", listaIds).Find(&imoveis)
	c.HTML(http.StatusOK, "views/recibos.html", gin.H{
		"imoveis": imoveis,
		"data":    data,
	})
}
