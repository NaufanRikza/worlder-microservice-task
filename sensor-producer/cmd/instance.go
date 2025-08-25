package cmd

import (
	"sensor-producer/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseInstance(config config.DatabaseConfig) (*gorm.DB, error) {
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
