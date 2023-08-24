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
	RunBasicHTTP()
}

func ReadConfig() {
	// Loading env config of path /configs/config.yaml
	viper.AddConfigPath("/configs")
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("yaml")   // Look for specific type
	viper.ReadInConfig()

	// Try to Dump all config from file by viper
	fmt.Println(viper.AllSettings())

	// Loading env config of key SOMECONFIGKEY
	someConfig := os.Getenv("SOMECONFIGKEY")
	fmt.Printf("Env Config is : %s", someConfig)

}

func RunBasicHTTP() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
