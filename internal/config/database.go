package config

import (
	"log"
	"myapi/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Conexão com o Postgres (usando host "db" pois o docker-compose cria essa rede)
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o BD: %v", err)
	}
	DB = db

	// AutoMigrate para criar/ajustar tabelas
	if err = DB.AutoMigrate(&models.Item{}); err != nil {
		log.Fatalf("Erro durante a miggração do Item: %v", err)
	}
	if err = DB.AutoMigrate(&models.Categoria{}); err != nil {
		log.Fatalf("Erro durante a migração da Categoria: %v", err)
	}
}
