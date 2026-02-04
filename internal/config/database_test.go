package config

import (
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectDatabase(t *testing.T) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		t.Fatal("DB_DSN não está definida")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar com o BD: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Erro ao obter uma conexão com o banco de dados: %v", err)
	}

	// Configuração do pool (ok para teste)
	sqlDB.SetConnMaxLifetime(time.Minute)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)

	defer func() {
		if err := sqlDB.Close(); err != nil {
			t.Errorf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
	}()

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}
}
