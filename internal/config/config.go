package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
