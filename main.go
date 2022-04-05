package main

import (
	"fmt"
	"github.com/kimkokubun/api-go-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
