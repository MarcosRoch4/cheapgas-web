package main

import (
	"fmt"

	"github.com/MarcosRoch4/cheapgas-web/migrate"
	"github.com/MarcosRoch4/cheapgas-web/routes"
)

func main() {
	migrate.IniMigrate()

	fmt.Println("Iniciando o servidor Rest com GO!")
	routes.HandleRequest()

}
