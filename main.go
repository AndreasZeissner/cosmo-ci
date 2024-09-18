package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	fmt.Println("Value:", viper.GetString("foo"))
	fmt.Println("Value:", viper.GetString("bar"))
	fmt.Println("Value:", viper.GetString("baz"))
}
