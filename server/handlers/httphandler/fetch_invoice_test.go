package httphandler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetInvoiceById(t *testing.T) {
	invoiceRec := GetInvoiceById()
	//assert.Equal(t, http.StatusOK, invoiceRec.)
	fmt.Println("This is tested!!", invoiceRec)

}

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

	router.GET("/invoice-display", GetInvoiceById())

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/invoice-display", nil)

	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}
