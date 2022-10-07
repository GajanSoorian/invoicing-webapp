package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository/models"
	"github.com/gin-gonic/gin"
)

// Handler for endpoint invoice-display

func GetInvoiceById(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("invoiceNumber"), 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		/*rows, err := repository.FindInvoice(db)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			fmt.Println(rows)
			err := rows.Scan(&id, &name, &email, &createdOn, &updatedOn)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		} */
		in := models.NewInvoice()
		in.CustomerName = "god"
		in.Email = "god@god.com"
		in.DueDate = time.Now()
		in.InvoiceNumber = id
		in.Products = []models.Item{{"pen", "good pen", 12.3}, {"pencil", "good Pencil", 10}}
		in.TotalAmount = 22
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
