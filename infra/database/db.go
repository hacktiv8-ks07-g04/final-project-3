package database

import (
	"fmt"
	"log"

	"github.com/hacktiv8-ks07-g04/final-project-3/entity"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/config"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
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
}

func Migration() {
	// AutoMigrate your models
	err := db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{})
	if err != nil {
		// Handle error
		log.Fatalf("Error migrating: %v", err)
	}

	// Check if the User table exists and is empty
	if db.Migrator().HasTable(&entity.User{}) {
		var count int64
		db.Model(&entity.User{}).Count(&count)
		if count == 0 {
			// Insert seed data
			createUserAdmin()
		}
	}
}

func createUserAdmin() {
	adminUser := &entity.User{
		FullName: "Admin",
		Email:    "admin@gmail.com",
		Password: "admin123",
		Role:     "admin",
	}

	err := adminUser.HashPassword()
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	result := db.Create(&adminUser)
	if result.Error != nil {
		err = errs.NewInternalServerError(result.Error.Error())
		log.Fatalf("Error creating user: %v", err)
	}

}

func GetDbInstance() *gorm.DB {
	if db == nil {
		log.Fatal("Database Instance is not initialized")
	}
	return db
}

func InitializedDatabase() {
	ConnectDB()
	Migration()
}
