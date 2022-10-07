package handlers

import (
	"fmt"
	"net/http"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/internal/config"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/GajanSoorian/parallax-invoicing/invoice-service/service"
	"github.com/gin-gonic/gin"
)

// Handler for endpoint: invoice/save.
func SaveInvoice(env *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := repository.GetDbHandle(env)
		invoice := &models.Invoice{}
		if err := c.Bind(invoice); err != nil {
		}
		resultInvoice, saveErr := service.SaveInvoice(db, invoice)
		if saveErr != nil {
			fmt.Println("Error saving invoice", saveErr)
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, resultInvoice)
		}
	}
}
