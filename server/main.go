package main

import (
	"log"

	models "github.com/GajanSoorian/parallax-invoicing/models"
	httphandler "github.com/GajanSoorian/parallax-invoicing/server/handlers/httphandler"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := models.SetupDatabase("postgres", "postgres://root:root@172.20.0.2/invoice_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creates a gin router - gin is a micro-framework for RESTful
	router := gin.Default()

	router.POST("/invoice-create", httphandler.CreateInvoice(db))
	// display-invoice is the endpoint.
	// GetInvoiceById is the handler.
	router.GET("/invoice-display", httphandler.GetInvoiceById(db))

	//router.PUT("/somePut", putting)
	//router.DELETE("/invoice-display", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":3000")
	// router.Run(":3000") for a hard coded port
}
