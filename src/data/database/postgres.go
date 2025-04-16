package database

import (
	"fmt"
	"log"
	"time"

	"github.com/arash2007mahdavi/web-api-1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DbName, cfg.Postgres.Port, cfg.Postgres.Sslmode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDb, _ := dbClient.DB()

	err = sqlDb.Ping()
	if err != nil { 
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("Db connection ok")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	if dbClient == nil {
        fmt.Println("Database connection is nil, nothing to close")
        return
    }
    con, err := dbClient.DB()
    if err != nil {
        fmt.Println("Error getting database instance:", err)
        return
    }
    con.Close()
}
