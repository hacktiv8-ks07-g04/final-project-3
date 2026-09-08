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

// ISSUE: Package-level mutable global variable for the DB connection (service locator
// pattern instead of dependency injection). Makes testing difficult and couples all
// consumers to this singleton. Should be passed explicitly via DI.
var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	config := config.GetConfig()
	// ISSUE: DSN hardcodes sslmode=disable and TimeZone=Asia/Shanghai.
	// These should be configurable via environment variables.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	// ISSUE: No connection pooling configuration (SetMaxOpenConns, SetMaxIdleConns,
	// SetConnMaxLifetime). Defaults may be insufficient for production traffic.
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// ISSUE: panic() used instead of returning an error / log.Fatal. Also the
		// message "failed to connect database" gives no hint about WHICH env var
		// or DB param is wrong.
		panic("failed to connect database")
	}
}

func Migration() {
	// AutoMigrate your models
	// ISSUE: AutoMigrate runs on every startup. It is not safe for production —
	// it can alter/drop columns, has no rollback, and doesn't support versioned
	// migrations. A tool like golang-migrate or goose should be used instead.
	err := db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{})
	if err != nil {
		// Handle error
		log.Fatalf("Error migrating: %v", err)
	}

	// Check if the User table exists and is empty
	// ISSUE: The count-then-insert seed logic is not atomic. In a multi-instance
	// deployment, two instances could both see count==0 and both insert, causing
	// a duplicate-key error.
	if db.Migrator().HasTable(&entity.User{}) {
		var count int64
		// ISSUE: Return value (including the error) from GORM's Count() is ignored.
		// If the query fails, count stays 0 and an unnecessary admin is created.
		db.Model(&entity.User{}).Count(&count)
		if count == 0 {
			// Insert seed data
			createUserAdmin()
		}
	}
}

func createUserAdmin() {
	// ISSUE: Admin credentials are hardcoded in source code and in README.md.
	// Any deploy starts with a known password (admin123). These should come from
	// environment variables or a one-time setup script.
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
