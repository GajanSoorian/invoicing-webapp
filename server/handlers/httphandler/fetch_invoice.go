package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for endpoint invoice-display

func GetInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"Invoice ID": " 1",
			"cost":       "500",
		})
	}
}
