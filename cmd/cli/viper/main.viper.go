package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath(("../../../config/")) //cấu hình đường đẫn
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("JWT::", viper.GetString("security.jwt.key"))

	// configuration structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
	fmt.Println("Config Port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("databases User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
