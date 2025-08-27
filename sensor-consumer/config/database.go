package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"192.168.1.50"`
	Port     int    `env:"DB_PORT" envDefault:"3306"`
	User     string `env:"DB_USER" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"Tornado090699"`
	Name     string `env:"DB_NAME" envDefault:"sensor_db"`
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	cfg := DatabaseConfig{}
	err := env.Parse(&cfg)
	return &cfg, err
}

func (d DatabaseConfig) GetDialector() gorm.Dialector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)
	return mysql.Open(dsn)
}

func (d DatabaseConfig) GetDSN() string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)
	return dsn
}
