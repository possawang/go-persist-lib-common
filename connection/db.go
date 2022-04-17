package connection

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionAndMigration(dst []interface{}) error {
	err := connecting()
	if err != nil {
		return err
	}
	for _, i := range dst {
		DB.AutoMigrate(&i)
	}
	return err
}

func connecting() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	user, pass, host, port, dbname := os.Getenv("DB.USER"), os.Getenv("DB.PASS"), os.Getenv("DB.HOST"), os.Getenv("DB.PORT"), os.Getenv("DB.NAME")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	return err
}
