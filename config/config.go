package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Address    string `mapstructure:"server_address"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
}

func Load() *Config {
	var config Config
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
