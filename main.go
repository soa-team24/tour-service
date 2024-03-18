package main

import (
	"tour-service/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:super@tcp(localhost:3306)/soa?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Equipment{}, &model.Checkpoint{}, &model.Tour{})

	return database
}

func main() {

	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

}
