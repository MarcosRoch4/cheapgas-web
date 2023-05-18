package routes

import (
	"net/http"

	"github.com/MarcosRoch4/cheapgas-web/controller"
	//"github.com/MarcosRoch4/cheapgas-web/middleware"
	//"github.com/gorilla/handlers"
	//"github.com/gorilla/mux"
)

func HandleRequest() {
	//r := mux.NewRouter()
	//r.Use(middleware.ContentTypeMiddleware)

	// Tela Inicial
	http.HandleFunc("/", controller.Index)

	// Gas Station
	http.HandleFunc("/gas_station", controller.PageGas_Station)
	http.HandleFunc("/new_gs", controller.PageNew)
	http.HandleFunc("/insert_gs", controller.InsertGasStation)
	http.HandleFunc("/delete_gs", controller.Delete)
	http.HandleFunc("/edit_gs", controller.PageEdit)
	http.HandleFunc("/update_gs", controller.Update)

	// Fuel
	http.HandleFunc("/fuel", controller.PageFuel)
	http.HandleFunc("/new_f", controller.PageNewFuel)
	http.HandleFunc("/insert_f", controller.InsertFuel)
	http.HandleFunc("/delete_f", controller.DeleteFuel)
	http.HandleFunc("/edit_f", controller.PageEditFuel)
	http.HandleFunc("/update_f", controller.UpdateFuel)

	// Category
	http.HandleFunc("/category", controller.PageCategory)
	http.HandleFunc("/new_c", controller.PageNewCategory)
	http.HandleFunc("/insert_c", controller.InsertCategory)
	http.HandleFunc("/delete_c", controller.DeleteCategory)
	http.HandleFunc("/edit_c", controller.PageEditCategory)
	http.HandleFunc("/update_c", controller.UpdateCategory)

	// Fuel Value
	http.HandleFunc("/fuel_value", controller.PageFuelValue)
	http.HandleFunc("/new_fv", controller.PageNewFuelValue)
	http.HandleFunc("/fuel_fv", controller.PageFuel)
	http.HandleFunc("/gs_fuel", controller.Gas_Station)
	http.HandleFunc("/insert_fv", controller.InsertFuelValue)

	http.ListenAndServe(":8000", nil)

	//log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
