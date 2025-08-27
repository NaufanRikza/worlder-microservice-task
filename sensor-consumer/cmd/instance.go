package cmd

import (
	"fmt"
	"log"
	"os"
	"sensor-consumer/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseInstance(conf *config.DatabaseConfig) (*gorm.DB, error) {
	gormConf := new(gorm.Config)

	gormConf.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	dialector := conf.GetDialector()
	fmt.Println("Connecting to database with DSN:", conf.GetDSN())

	instance, err := gorm.Open(dialector, gormConf)
	if err != nil {
		return nil, err
	}

	//disable in prod
	instance.Debug()

	return instance, nil
}
