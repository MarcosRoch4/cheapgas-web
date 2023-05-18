package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcosRoch4/cheapgas-web/models"
)

func PageFuel(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Fuel", models.ConsultaFuel())
}

func PageNewFuel(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New_F", nil)
}

func PageEditFuel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("ID do combustível:", id)
	fuel := models.EditFuel(id)
	temp.ExecuteTemplate(w, "Edit_F", fuel)
	http.Redirect(w, r, "/fuel", 301)

}

func DeleteFuel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("deletando combustível com ID -->", id)
	models.DeleteFuel(id)
	http.Redirect(w, r, "/fuel", 301)
}

func InsertFuel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("fuel_name")

		models.CreateFuel(name)
	}

	http.Redirect(w, r, "/fuel", 301)

}

func UpdateFuel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("fuel_name")

		fmt.Println("atualizando primeira instancia", name)
		models.UpdateFuel(id, name)
	}

	http.Redirect(w, r, "/fuel", 301)

}
