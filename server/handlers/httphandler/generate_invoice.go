package httphandler

import (
	//"parallax-invoicing/models/invoicemodels"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//handler for endpoint: invoice-create

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Invoice created")
		c.JSON(http.StatusOK, map[string]string{
			"Invoice ID": " 1",
			"cost":       "500",
		})
	}

}

// postAlbums adds an album from JSON received in the request body.
func PostInvoice(c *gin.Context) {
	/*newInvoice := invoicemodels.New()

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(newInvoice); err != nil {*/
}

// Add the new album to the slice.
/*albums = append(albums, newAlbum)
c.IndentedJSON(http.StatusCreated, newAlbum)*/
