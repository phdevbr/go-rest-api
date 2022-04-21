package main

import (
	"fmt"

	"github.com/phdevbr/go-rest-api/models"
	"github.com/phdevbr/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome1", Historia: "Historia1"},
		{Nome: "Nome2", Historia: "Historia2"},
	}

	fmt.Println("Iniciando o servidor REST com Go...")
	routes.HandleRequest()
}
