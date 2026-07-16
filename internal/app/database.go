package app

import (
	"fmt"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB() (*gorm.DB, error) {
	host := GetEnv("DB_HOST", "localhost")
	user := GetEnv("DB_USER", "postgres")
	password := GetEnv("DB_PASSWORD", "1234")
	dbname := GetEnv("DB_NAME", "db_belajar")
	port := GetEnv("DB_PORT", "5432")
	sslmode := GetEnv("DB_SSLMODE", "disable")
	timezone := GetEnv("DB_TIMEZONE", "Asia/Makassar")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
	
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke datatabse %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("gagal memanggil sql.DB %w", err)
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}


