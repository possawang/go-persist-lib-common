package connection

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionAndMigration(dst ...interface{}) error {
	db, err := Connecting()
	if err != nil {
		return err
	}
	db.AutoMigrate(dst...)
	return err
}

func Connecting() (*gorm.DB, error) {
	var db *gorm.DB
	err := godotenv.Load(".env")
	if err != nil {
		return db, err
	}
	user, pass, host, port, dbname := os.Getenv("DB.USER"), os.Getenv("DB.PASS"), os.Getenv("DB.HOST"), os.Getenv("DB.PORT"), os.Getenv("DB.NAME")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	return db, err
}
