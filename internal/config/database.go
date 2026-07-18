package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ilmu-merah/be-article/internal/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sqlDB *sql.DB


func InitDB() (*gorm.DB, error) {
	host := app.GetEnv("DB_HOST", "localhost")
	user := app.GetEnv("DB_USER", "postgres")
	password := app.GetEnv("DB_PASSWORD", "1234")
	dbname := app.GetEnv("DB_NAME", "db_belajar")
	port := app.GetEnv("DB_PORT", "5432")
	sslmode := app.GetEnv("DB_SSLMODE", "disable")
	timezone := app.GetEnv("DB_TIMEZONE", "Asia/Makassar")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
	
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke datatabse %w", err)
	}

	var errDB error
	sqlDB, errDB = db.DB()
	if errDB != nil {
		return nil, fmt.Errorf("gagal memanggil sql.DB %w", err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func CloseDB() {
	if sqlDB != nil {
		sqlDB.Close()

		log.Println("Koneksi database berhasil ditutup")
	}
}


