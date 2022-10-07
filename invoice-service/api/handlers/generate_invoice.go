package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/gin-gonic/gin"
)

// handler for endpoint: invoice-create
func CreateInvoice(db *sql.DB, invoice *models.Invoice) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Bind(invoice); err != nil {
			fmt.Println("Error binding", err)
		}
		fmt.Println("Invoice created", invoice)
		c.JSON(http.StatusOK, populateInvoice(db)) //subPopulate(invoice)) //

	}

}

func populateInvoice(db *sql.DB) *models.Invoice {
	invoice := models.NewInvoice()
	// TODO: get these values from frontend
	invoice.CustomerName = "Gajan Soorian"
	invoice.Email = "testmail@gmail.com"
	//invoice.Id = uuid.New()
	/*sqlStatement := `INSERT INTO invoices (id, customer_name, customer_email, created_on, updated_on)
	VALUES ($1, $2, $3, $4, $5)`
		_, err := db.Exec(sqlStatement, 1, "GajanSoorian", "tes@gmail.com", time.Now(), time.Now())
		if err != nil {
			panic(err)
		}*/
	return invoice
}

func PostInvoice(c *gin.Context) {
	/*newInvoice := invoicemodels.New()

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(newInvoice); err != nil {*/
}
