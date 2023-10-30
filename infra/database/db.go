package database

import (
	// "fmt"
	// "log"

	"fmt"

	"github.com/hacktiv8-ks07-g04/final-project-3/config"
	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	config := config.GetConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.Debug().AutoMigrate(entity.User{}, entity.Category{}, entity.Task{})
}
