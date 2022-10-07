package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

//Todo : mock DB call
/*func TestGetInvoiceById(t *testing.T) {
	var db *sql.DB
	invoiceRec := GetInvoiceById(db)
	assert.Equal(t, http.StatusOK, invoiceRec.)

}
*/

// Helper function to create a router during testing
func getRouter() *gin.Engine {
	r := gin.Default()
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	router := getRouter()
	var db *sql.DB
	router.GET("/v1/invoice/view/1235", GetInvoiceById(db))

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/v1/invoice/view/12345", nil)
	assert.Equal(t,http.StatusOK, req.Response.StatusCode)
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}
