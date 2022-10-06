package main

import (
	"log"
	"os"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/api/handlers"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/gin-gonic/gin"
)

func main() {

	env := readEnvVariables()

	db, _ := repository.SetupDbConnection(env)
	/*if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}*/

	// Creates a gin router - gin is a micro-framework for RESTful
	router := gin.Default()
	api := router.Group(env.ApiVersion)
	{
		api.GET("/invoice/ping", handlers.PingGet())
		api.POST("/invoice/save", handlers.CreateInvoice(db, models.NewInvoice()))
		api.GET("/invoice/view/:id", handlers.GetInvoiceById(db))
	}
	router.Run(os.Getenv("SERVER_PORT"))
}

// util function to read environment variables from env file.
func readEnvVariables() *config.Config {
	currentPath, _ := os.Getwd()
	env, err := config.GetEnvVariables(currentPath + "/invoice-service/invoice-service.env")
	if err != nil {
		log.Fatal("Error config file format not support, please check invoice-service.env", err)
	}
	return env
}