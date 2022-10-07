package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/service"
	"github.com/gin-gonic/gin"
)

// Handler for endpoint invoice/view/:id
func GetInvoiceById(env *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := repository.GetDbHandle(env)
		invoiceNumber, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			fmt.Println("error with string to into: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		in := service.FindInvoice(db, invoiceNumber)
		if in == nil {
			c.IndentedJSON(http.StatusBadRequest, models.NewInvoice())
		}
		fmt.Println("invoice returned :", in)
		c.IndentedJSON(http.StatusOK, in)
	}
}

// Handler for testing endpoint connection.
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name":  "testingConnection",
			"email": "tester@test.com",
		})
	}
}
