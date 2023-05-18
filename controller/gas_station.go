package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/MarcosRoch4/cheapgas-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Index", models.CheapGasHint())
}

func PageGas_Station(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Gas_Station", models.ConsultaGasStation())
}

func Gas_Station(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New_FV", models.ConsultaGS())
}

func PageNew(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New_GS", nil)
}

func PageEdit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("ID do posto:", id)
	gas_station := models.EditGasStation(id)
	temp.ExecuteTemplate(w, "Edit_GS", gas_station)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("deletando posto com ID -->", id)
	models.DeleteGasStation(id)
	http.Redirect(w, r, "/gas_station", 301)
}

func InsertGasStation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		categoryID := r.FormValue("categoryID")

		icategoryID, err := strconv.Atoi(categoryID)
		if err != nil {
			log.Println("Erro na conversão do ID para Int.")
		}

		models.CreateGasStation(name, icategoryID)
	}

	http.Redirect(w, r, "/gas_station", 301)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		categoryID := r.FormValue("categoryID")

		icategoryID, err := strconv.Atoi(categoryID)
		if err != nil {
			log.Println("Erro na conversão do ID para Int.")
		}

		models.UpdateGasStation(id, name, icategoryID)
	}

	http.Redirect(w, r, "/gas_station", 301)

}
