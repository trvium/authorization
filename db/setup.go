package db

import (
	"os"

	"github.com/trvium/authorization/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() error {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func Migrate() error {
	m := []interface{}{
		&models.Plan{},
	}

	err := DB.AutoMigrate(m...)
	if err != nil {
		return err
	}

	return nil
}
