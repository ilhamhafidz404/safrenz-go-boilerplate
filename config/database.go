package config

import (
	"fmt"
	"log"
	"safrenz-go-boilerplate/internal/model"
	"safrenz-go-boilerplate/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := utils.GetEnv("DB_HOST", "localhost")
	user := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "postgres")
	dbName := utils.GetEnv("DB_NAME", "safrenz_db")
	port := utils.GetEnv("DB_PORT", "5432")
	sslMode := utils.GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		host, user, password, dbName, port, sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database PostgreSQL: %v", err)
	}

	log.Println("Koneksi ke PostgreSQL berhasil!")

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	log.Println("Auto-Migrate PostgreSQL berhasil!")

	return db
}
