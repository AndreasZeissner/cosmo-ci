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
		log.Fatalf("Error reading file, %s", err)
	}

	fmt.Println("Getting Values:")
	fmt.Println("Value:", viper.GetString("foo"))
	fmt.Println("Value:", viper.GetString("bar"))
	fmt.Println("Value:", viper.GetString("baz"))

	values := []string{viper.GetString("foo"), viper.GetString("bar"), viper.GetString("baz")}
	for _, value := range values {
		fmt.Println("Feature Value:", value)
	}
}
