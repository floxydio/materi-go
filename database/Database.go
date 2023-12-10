package database

import (
	"belajarbe/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBInit *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/testdb"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database is Connected")
	DBInit = db
	Migrate()
}

func Migrate() {
	DBInit.AutoMigrate(models.TodoModel{})
}
