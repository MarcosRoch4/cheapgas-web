package controller

import (
	"fmt"
	"net/http"

	"github.com/MarcosRoch4/cheapgas-web/models"
)

func PageCategory(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Category", models.ConsultaCategory())
}

func PageNewCategory(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New_C", nil)
}

func PageEditCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("ID da categoria:", id)
	fuel := models.EditCategory(id)
	temp.ExecuteTemplate(w, "Edit_C", fuel)
	http.Redirect(w, r, "/category", 301)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("deletando categoria com ID -->", id)
	models.DeleteCategory(id)
	http.Redirect(w, r, "/category", 301)
}

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("category_name")

		models.CreateCategory(name)
	}

	http.Redirect(w, r, "/category", 301)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("category_name")

		fmt.Println("atualizando primeira instancia", name)
		models.UpdateCategory(id, name)
	}

	http.Redirect(w, r, "/category", 301)

}
