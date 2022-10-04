package main

import (
	"fmt"
	"github.com/spf13/viper"
	"julo-test"
	"julo-test/internal/config"
)

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	cfg := config.InitConfig()

	server := julo.NewServer(cfg)

	err = server.Start(cfg.ServerAddress())

	if err != nil {
		cfg.Logger().Fatal("Port already used")
	}
}
