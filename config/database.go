package config

import (
	"fmt"
	"github.com/poloping/interview_chenping_20240731/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func DatabaseConnection() *gorm.DB {
	host := GetEnvDefault("DB_HOST", "localhost")
	user := GetEnvDefault("DB_USER", "postgres")
	password := GetEnvDefault("DB_PASSWORD", "yourpassword")
	dbName := GetEnvDefault("DB_NAME", "yourdbname")
	port := GetEnvDefault("DB_PORT", "5432")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
