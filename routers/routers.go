package routers

import (
	"aluguel/controllers"
	"aluguel/utils"
	"github.com/gin-gonic/gin"
	"html/template"
)

func IniciaRoteamento() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formataDataUS": utils.FormataDataUS,
		"formataDataBR": utils.FormataDataBR,
		"formataFloat":  utils.FormataFloat,
	})
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.LoadHTMLGlob("templates/*/**")
	r.GET("/", controllers.Home)
	r.GET("/form_cliente", controllers.FormCliente)
	r.GET("/clientes", controllers.Clientes)
	r.POST("/clientes", controllers.Clientes)
	r.GET("/imoveis", controllers.Imoveis)
	r.POST("/imoveis", controllers.Imoveis)
	r.GET("/form_imovel", controllers.FormImovel)
	r.POST("/recibos", controllers.Recibos)

	r.Run(":4000")
}
