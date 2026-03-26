package main

import (
	"fmt"

	"github.com/spf13/viper"
)


func main() {
	viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//readfileconfig

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("False to read config %w \n", err))
	}

	fmt.Println("server port", viper.GetInt("server.port"))
}