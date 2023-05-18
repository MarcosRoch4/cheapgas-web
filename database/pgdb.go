package database

import (
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {

	dbConect := "host=mouse.db.elephantsql.com user=lgngbpdo password=5NqqDGWFkazYmraBjtPhJAeUoJzFJ5Kt dbname=lgngbpdo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dbConect), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o Banco de dados")
	}
}
