package models

import (
	"gorm.io/gorm"
	"time"
)

type Cliente struct {
	gorm.Model
	Nome           string    `json:"nome"`
	DataNascimento time.Time `json:"data_nascimento"`
	CI             string    `json:"documento_identidade"`
	CPF            string    `json:"cpf"`
	Telefone1      string    `json:"telefone_1"`
	Telefone2      string    `josn:"telefone_2"`
}
