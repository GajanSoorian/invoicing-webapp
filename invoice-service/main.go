package main

import (
	"log"
	"os"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/api/handlers"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	env := readEnvVariables()

	// Creates a gin router - gin is a micro-framework for creating REST APIs
	router := gin.Default()

	api := router.Group(env.ApiVersion)
	{
		api.GET("/invoice/ping", handlers.PingGet())
		api.POST("/invoice/save", handlers.SaveInvoice(env))
		api.GET("/invoice/view/:id", handlers.GetInvoiceById(env))
	}
	router.Run(os.Getenv("SERVER_ADDRESS") + os.Getenv("SERVER_PORT"))
}

// Util function to read environment variables from env file.
func readEnvVariables() *config.Config {
	currentPath, _ := os.Getwd()
	env, err := config.GetEnvVariables(currentPath + "/invoice-service/invoice-service.env")
	if err != nil {
		log.Fatal("Error config file format not support, please check invoice-service.env", err)
	}
	return env
}