package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcosRoch4/cheapgas-web/models"
)

func PageFuelValue(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "FuelValue", models.ConsultaFuelValue())
}

func PageNewFuelValue(w http.ResponseWriter, r *http.Request) {
	//fuels := models.ConsultaFuel()
	//fmt.Println("Combustíveis:", fuels)

	gas := models.ConsultaFuelValue()
	//gStation := models.ConsultaGS()
	//	fmt.Println("Posto de Combustível:", gStation)
	fmt.Println("HINT:", gas)

	temp.ExecuteTemplate(w, "New_FV", gas)

	//temp.Execute(w, fuels)
	//temp.Execute(w, gStation)

	//temp.ExecuteTemplate(w, "New_FV", fuels)

}

func InsertFuelValue(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		fuel_id := r.FormValue("fuel_id")
		price := r.FormValue("price")
		fmt.Printf("\n")
		fmt.Println("--------# Inserindo valor ao Combustível #--------")
		fmt.Printf("\n")

		fmt.Println("o Retorno da tela: ", id, fuel_id, price)
		models.CreateFuelValue(id, fuel_id, price)

	}

	http.Redirect(w, r, "/fuel_value", 301)

}
