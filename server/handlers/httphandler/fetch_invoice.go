package httphandler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GajanSoorian/parallax-invoicing/models"
	"github.com/gin-gonic/gin"
)

// Handler for endpoint invoice-display


func GetInvoiceById(db *sql.DB) gin.HandlerFunc {
	var (
	id int
	name string
	email string
	createdOn time.Time
	updatedOn time.Time
)
fmt.Println("Calling FetchRows")
rows, err := models.FetchRows(db)
	if err!= nil {
        log.Println(err)
    }
		fmt.Println(rows)
for rows.Next() {
	fmt.Println(rows)
	err := rows.Scan(&id, &name, &email, &createdOn, &updatedOn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lets see wjat we get!!!!")
	fmt.Println(id, name, email, createdOn, updatedOn)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name":       name,
			"email":      email,
		})
	}
}
