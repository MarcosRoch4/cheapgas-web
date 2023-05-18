package models

import (
	"fmt"

	"github.com/MarcosRoch4/cheapgas-web/database"
)

type Gas_Station struct {
	Id         int    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

type GStation struct {
	Id   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
}

var Gas_Stations []Gas_Station

type Category struct {
	Id           int    `gorm:"primary_key;auto_increment" json:"id"`
	CategoryName string `json:"category_name"`
}

var Categories []Category

type GS_Category struct {
	Id           int    `json:"gas_station_id"`
	Name         string `json:"gas_station_name"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func Resul() []GStation {

	gs := []GStation{}
	database.DB.Raw("SELECT id,name FROM gas_stations").Scan(&gs)

	item := GStation{}
	gas_station := []GStation{}

	for _, v := range gs {
		item.Id = v.Id
		item.Name = v.Name

		gas_station = append(gas_station, item)
	}

	//	fmt.Println("Resultado:", gas_station)
	return gas_station

}

func ConsultaGS() []GStation {

	db_gasstation := []GStation{}

	database.DB.Raw("SELECT id,name FROM gas_stations ORDER BY id ASC").
		Scan(&db_gasstation)

	item := GStation{}
	gas_station := []GStation{}

	for _, v := range db_gasstation {
		item.Id = v.Id
		item.Name = v.Name

		gas_station = append(gas_station, item)
	}

	//fmt.Println("Posto de gasolina ===>", db_gasstation)

	return gas_station

}

func ConsultaGasStation() []GS_Category {

	gs_category := []GS_Category{}

	database.DB.Raw("SELECT gas_stations.id,gas_stations.name, gas_stations.category_id, categories.category_name FROM gas_stations " +
		"LEFT JOIN categories ON categories.ID = gas_stations.category_id ORDER BY gas_stations.id ASC").
		Scan(&gs_category)

	item := GS_Category{}
	gs_cat := []GS_Category{}

	for _, v := range gs_category {
		item.Id = v.Id
		item.Name = v.Name
		item.CategoryID = v.CategoryID
		item.CategoryName = v.CategoryName

		gs_cat = append(gs_cat, item)
	}

	//	fmt.Println("Posto de Combustível ===>", gs_category)

	return gs_cat

}

func CreateGasStation(name string, categoryID int) {
	gas_station := Gas_Station{Name: name, CategoryID: categoryID}
	database.DB.Create(&gas_station)
	fmt.Println("Inserindo os dados do posto de gasolina, nome: ?, categoria: ?", name, categoryID)

}

func EditGasStation(id string) Gas_Station {
	gas_station := []Gas_Station{}
	database.DB.First(&gas_station, id)

	item := Gas_Station{}

	for _, v := range gas_station {
		item.Id = v.Id
		item.Name = v.Name
		item.CategoryID = v.CategoryID

		gas_station = append(gas_station, item)
	}

	//	fmt.Println("Gas Station ==>", item)

	return item
}

func UpdateGasStation(id, name string, categoryID int) {
	gas_station := []Gas_Station{}
	// função que está atualizando vários campos ao mesmo tempo
	database.DB.First(&gas_station, id).Updates(Gas_Station{Name: name, CategoryID: categoryID})

}

func DeleteGasStation(id string) {
	gas_station := []Gas_Station{}
	database.DB.Delete(&gas_station, id)

}

// Categorias

func ConsultaCategory() []Category {

	db_category := []Category{}

	database.DB.Raw("SELECT categories.id,categories.category_name FROM categories ORDER BY categories.id ASC").
		Scan(&db_category)

	item := Category{}
	category := []Category{}

	for _, v := range db_category {
		item.Id = v.Id
		item.CategoryName = v.CategoryName

		category = append(category, item)
	}

	//	fmt.Println("Categoria ===>", db_category)

	return category

}

func CreateCategory(name string) {

	category := Category{CategoryName: name}
	database.DB.Create(&category)
	fmt.Println("Inserindo os dados da categoria, nome: ?", name)

}

func EditCategory(id string) Category {
	category := []Category{}
	database.DB.First(&category, id)

	item := Category{}

	for _, v := range category {
		item.Id = v.Id
		item.CategoryName = v.CategoryName

		category = append(category, item)
	}

	//	fmt.Println("Categoria #==>", item)

	return item
}

func UpdateCategory(id, name string) {
	category := []Category{}

	//	fmt.Println("atualizando segunda instancia", name)
	// função que está atualizando apenas 1 campo
	database.DB.Model(&category).Where("id = ?", id).Update("category_name", name)

}

func DeleteCategory(id string) {
	category := []Category{}
	database.DB.Delete(&category, id)

}
