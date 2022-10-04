package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	User       string
	DbPassword string
	Driver     string
	DbName     string
	Host       string
	Port       string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		User:       viper.GetString("DB_USER"),
		DbPassword: viper.GetString("DB_PASSWORD"),
		DbName:     viper.GetString("DB_NAME"),
		Host:       viper.GetString("DB_HOST"),
		Port:       viper.GetString("DB_PORT"),
	}
}

func (d *DBConfig) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		d.Host,
		d.Port,
		d.User,
		d.DbName,
		d.DbPassword,
	)
}
