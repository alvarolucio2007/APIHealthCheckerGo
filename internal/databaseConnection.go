package internal

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBanco() {
	dsn := "host=localhost user=user_api password=123456 dbname=db_api port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Não foi possível concectar ao Postgres!", err)
	}
	fmt.Println("Banco conectado com sucesso!")
	err = DB.AutoMigrate(&ChecagemAPI{})
	if err != nil {
		log.Fatal("Falha ao rodar as migrações", err)
	}
}
