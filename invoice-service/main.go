package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/api/handlers"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	env := readEnvVariables()

	db, err := repository.SetupDatabase(env)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	// Creates a gin router - gin is a micro-framework for RESTful
	router := gin.Default()

	router.POST("/v1/invoice-create", handlers.CreateInvoice(db))

	router.GET("/v1//invoice-display", handlers.GetInvoiceById(db))

	router.Run(os.Getenv("SERVER_PORT"))
}

func readEnvVariables() *config.Config {
	currentPath, _ := os.Getwd()
	env, err := config.GetEnvVariables(currentPath + "/invoice-service/invoice-service.env")
	if err != nil {
		log.Fatal("Error config file format not support, please check invoice-service.env", err)
	}
	return env
}
