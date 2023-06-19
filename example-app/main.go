package main

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/spf13/viper"
	"net/http"
	"os"
)

func main() {
	DumpConfig()
	DumpEnv()
	RunBasicHTTP()
}

func DumpConfig() {
	viper.AddConfigPath("/configs")
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	viper.ReadInConfig()

	// Dump all config
	fmt.Println(viper.AllSettings())
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
