package cmd

import (
	"sensor-producer/config"
)

func StartHTTPServer() {
	// connect to database
	dbConfig, err := config.GetDatabaseConfig()
	if err != nil {
		panic("Failed to load database configuration: " + err.Error())
	}
	
	db, err := NewDatabaseInstance(dbConfig)
	if err != nil {
		// handle error
		panic("Failed to connect to database: " + err.Error())
	}
}
