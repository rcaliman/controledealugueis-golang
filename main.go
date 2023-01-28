package main

import (
	"aluguel/databases"
	"aluguel/routers"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor")
	databases.ConectaBanco()
	routers.IniciaRoteamento()
}
