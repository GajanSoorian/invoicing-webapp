package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/gin-gonic/gin"
)

// Handler for endpoint invoice-display

func GetInvoiceById(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("invoiceNumber")

		fmt.Println("Calling FetchRows for id ", id)
		/*rows, err := repository.FindInvoice(db)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(rows)
		for rows.Next() {
			fmt.Println(rows)
			err := rows.Scan(&id, &name, &email, &createdOn, &updatedOn)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name, email, createdOn, updatedOn)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		} */
		in := models.NewInvoice()
		in.CustomerName = "god"
		in.Email = "god@god.com"
		in.DueDate = time.Now()
		in.InvoiceNumber = 12345
		in.Products = []models.Item{{"pen", "good pen", 12.3}, {"pencil", "good Pencil", 10}}
		in.TotalAmount = 22
		fmt.Println("sending to front end: ", in)
		c.IndentedJSON(http.StatusOK, in)
	}
}

func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name":  "Gajan",
			"email": "My@gmail.com",
		})
	}
}
