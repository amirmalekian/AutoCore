package db

import (
	"fmt"
	"log"
	"time"

	"github.com/amirmalekian/Mashinchi/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDB(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DBName, cfg.Postgres.SSLMode)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("DB Connection Established")

	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	conn, _ := dbClient.DB()
	conn.Close()
}
