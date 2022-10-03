package httphandler

import (
	//"parallax-invoicing/models/invoicemodels"

	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/GajanSoorian/parallax-invoicing/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//handler for endpoint: invoice-create

func CreateInvoice(db *sql.DB) gin.HandlerFunc {
	sqlStatement := `
INSERT INTO invoices (id, customer_name, customer_email, created_on, updated_on)
VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, 1, "GajanSoorian", "tes@gmail.com", time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, populateInvoice())
		fmt.Println("Invoice created")

	}

}

func populateInvoice() *models.Invoice {
	invoice := models.NewInvoice()
	// TODO: get these values from frontend
	invoice.CustomerName = "Gajan Soorian"
	invoice.CustomerEmail = "testmail@gmail.com"
	invoice.Id = uuid.New()

	return invoice
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
