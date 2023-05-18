package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/MarcosRoch4/cheapgas-web/database"
)

type Fuel struct {
	Id       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	FuelName string `json:"fuel_name"`
}

var Fuels []Fuel

type Fuel_Value struct {
	Id             int     `gorm:"primary_key;auto_increment" json:"id"`
	Price          float64 `json:"price"`
	Fuel_Id        int     `gorm:"not null" json:"fuel_id"`
	Gas_Station_Id int     `gorm:"not null" json:"gas_station_id"`
}

var Fuel_Values []Fuel_Value

type GS_Fuel struct {
	Gas_Station_Id uint32
	Name           string
	Fuel_Id        uint32
	FuelName       string
	Price          float64
}

type CheapGas struct {
	Id           int     `json:"gas_station_id"`
	Name         string  `json:"gas_station_name"`
	CategoryID   int     `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Fuel_Id      int     `json:"fuel_id"`
	FuelName     string  `json:"fuel_name"`
	Price        float64 `json:"price"`
}

type CheapGasValue struct {
	Id       int     `json:"gas_station_id"`
	Name     string  `json:"gas_station_name"`
	Fuel_Id  int     `json:"fuel_id"`
	FuelName string  `json:"fuel_name"`
	Price    float64 `json:"price"`
}

func CheapGasHint() []CheapGas {

	cheapHint := []CheapGas{}

	database.DB.Raw("SELECT gas_stations.id,gas_stations.name, gas_stations.category_id, categories.category_name, " +
		" fuels.id, fuels.fuel_name, fuel_values.id, fuel_values.price" +
		" FROM gas_stations " +
		" INNER JOIN fuel_values ON fuel_values.gas_station_id = gas_stations.id" +
		" INNER JOIN fuels ON fuels.id = fuel_values.fuel_id " +
		" INNER JOIN categories ON categories.ID = gas_stations.category_id " +
		" ORDER BY fuel_values.price, fuels.fuel_name, gas_stations.id ASC").
		Scan(&cheapHint)

	item := CheapGas{}
	cheapgas := []CheapGas{}

	for _, v := range cheapHint {
		item.Id = v.Id
		item.Name = v.Name
		item.CategoryID = v.CategoryID
		item.CategoryName = v.CategoryName
		item.Fuel_Id = v.Fuel_Id
		item.FuelName = v.FuelName
		item.Price = v.Price

		cheapgas = append(cheapgas, item)
	}

	//	fmt.Println("CheapGas ===>", cheapHint)

	return cheapgas

}

func ConsultaFuel() []Fuel {

	db_fuel := []Fuel{}

	database.DB.Raw("SELECT fuels.id,fuels.fuel_name FROM fuels ORDER BY fuels.id ASC").
		Scan(&db_fuel)

	item := Fuel{}
	fuel := []Fuel{}

	for _, v := range db_fuel {
		item.Id = v.Id
		item.FuelName = v.FuelName

		fuel = append(fuel, item)
	}

	//	fmt.Println("Combustivel ===>", db_fuel)

	return fuel

}

func CreateFuel(name string) {
	fuel := Fuel{FuelName: name}
	database.DB.Create(&fuel)
	fmt.Println("Inserindo os dados do combutível, nome: ?", name)

}

func EditFuel(id string) Fuel {
	fuel := []Fuel{}
	database.DB.First(&fuel, id)

	item := Fuel{}

	for _, v := range fuel {
		item.Id = v.Id
		item.FuelName = v.FuelName

		fuel = append(fuel, item)
	}

	//	fmt.Println("Combustível #==>", item)

	return item
}

func UpdateFuel(id, name string) {
	fuel := []Fuel{}

	fmt.Println("atualizando segunda instancia", name)
	// função que está atualizando apenas 1 campo
	database.DB.Model(&fuel).Where("id = ?", id).Update("fuel_name", name)

}

func DeleteFuel(id string) {
	fuel := []Fuel{}
	database.DB.Delete(&fuel, id)

}

func ConsultaFuelValue() []CheapGasValue {

	cheapgasvalue := []CheapGasValue{}

	database.DB.Raw("SELECT gas_stations.id,gas_stations.name,  fuels.id, fuels.fuel_name, " +
		"fuel_values.id, fuel_values.price" +
		" FROM gas_stations " +
		" LEFT JOIN fuel_values ON fuel_values.gas_station_id = gas_stations.id" +
		" LEFT JOIN fuels ON fuels.id = fuel_values.fuel_id " +
		" ORDER BY gas_stations.name, fuels.fuel_name, gas_stations.id ASC").
		Scan(&cheapgasvalue)

	item := CheapGasValue{}
	cheapgas := []CheapGasValue{}

	for _, v := range cheapgasvalue {
		item.Id = v.Id
		item.Fuel_Id = v.Fuel_Id
		item.Name = v.Name
		item.FuelName = v.FuelName
		item.Price = v.Price

		cheapgas = append(cheapgas, item)
	}

	fmt.Println("Preço do Combustível ===>", cheapgasvalue)

	return cheapgas

}

func CreateFuelValue(id, fuel_id, price string) {

	id_Converted, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na conversão do ID do Combustível para Int.")
	}

	fuelId_Converted, err := strconv.Atoi(fuel_id)
	if err != nil {
		log.Println("Erro na conversão do ID do Combustível para Int.")
	}

	price_Converted, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Erro na conversão do preço para float64.")
	}

	fuelValue := Fuel_Value{
		Gas_Station_Id: id_Converted,
		Fuel_Id:        fuelId_Converted,
		Price:          price_Converted}
	database.DB.Create(&fuelValue)
	fmt.Println("Inserindo os valores dos combutíveis, posto: ?, combustivel: ?, preço: ?", id, fuel_id, price)

}
