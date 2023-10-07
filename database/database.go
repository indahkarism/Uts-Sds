package database

import (
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := `root@tcp(localhost:3306)/utssdsindahkarisma`

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(models.User{})
}
