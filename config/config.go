package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	PostgresHost string
	PostgresPort int
	PostgresUser string
	PostgresPass string
	PostgresDB   string
	Hash     string
	HashComplixity interface{}
	JWTSecret string
	JWTExpiry int
	AppPort int
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't read config:", err)
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Can't unmarshal config:", err)
	}
	return &cfg
}