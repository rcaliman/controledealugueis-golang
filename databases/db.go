package databases

import (
	"aluguel/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	dsn := "root:mysql@tcp(172.17.0.2)/aluguel?parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados: ", err)
	}

	DB.AutoMigrate(&models.Cliente{})
	DB.AutoMigrate(&models.Imovel{})
}
