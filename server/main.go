package main

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/GajanSoorian/parallax-invoicing/models"
	httphandler "github.com/GajanSoorian/parallax-invoicing/server/handlers/httphandler"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://root:root@172.21.0.2/invoice_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database!!!!")
	}
	defer db.Close()
	// Creates a gin router - gin is a micro-framework for RESTful
	router := gin.Default()

	// display-invoice is the endpoint.
	// GetInvoiceById is the handler.
	router.GET("/invoice-display", httphandler.GetInvoiceById())

	router.POST("/invoice-create", httphandler.CreateInvoice())

	invoice := models.New()
	invoice.CustomerName = "Gajan"
	fmt.Println(invoice.CustomerName)
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
