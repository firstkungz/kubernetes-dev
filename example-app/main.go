package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"net/http"
	"os"

	"github.com/spf13/viper"
)

func main() {
	ReadConfig()
	DumpEnv()
	RunBasicHTTP()
}

func ReadConfig() {
	viper.AddConfigPath("/configs")
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	viper.ReadInConfig()

	// Dump all config
	fmt.Println(viper.AllSettings())

	someConfig := os.Getenv("SOMECONFIGKEY")
	fmt.Printf("Config is : %s", someConfig)

}

func DumpEnv() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

func RunBasicHTTP() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
