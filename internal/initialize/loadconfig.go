package initialize

import (
	"fmt"
	"go-backend-api/global"

	"github.com/spf13/viper"
)

func Loadconfig() {
	viper := viper.New()
	viper.AddConfigPath(("./config/")) //cấu hình đường đẫn
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v\n", err)
	}
}
