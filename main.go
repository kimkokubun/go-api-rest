package main

import (
	"fmt"
	"github.com/kimkokubun/go-api-rest/models"
	"github.com/kimkokubun/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "PERSONALIDADE 1", Historia: "Historia 1"},
		{Id: 2, Nome: "PERSONALIDADE 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
