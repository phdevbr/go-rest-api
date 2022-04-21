package main

import (
	"fmt"

	"github.com/phdevbr/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor REST com Go...")
	routes.HandleRequest()
}
