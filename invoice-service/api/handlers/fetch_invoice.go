package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GajanSoorian/parallax-invoicing/invoice-service/repository"
	"github.com/gin-gonic/gin"
)

// Handler for endpoint invoice-display

func GetInvoiceById(db *sql.DB) gin.HandlerFunc {
	var (
		id        int
		name      string
		email     string
		createdOn time.Time
		updatedOn time.Time
	)

	return func(c *gin.Context) {

		fmt.Println("Calling FetchRows")
		rows, err := repository.FetchRows(db)
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
		}
		c.JSON(http.StatusOK, map[string]string{
			"name":  name,
			"email": email,
		})
	}
}
