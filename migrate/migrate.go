package migrate

import (
	"fmt"

	"github.com/MarcosRoch4/cheapgas-web/database"
)

func IniMigrate() {
	database.ConectaDB()

	fmt.Println("Banco de dados Conectado!")

	//database.DB.AutoMigrate(&models.Fuel_Value{}, &models.Fuel{}, &models.Gas_Station{}, &models.Category{})
}
