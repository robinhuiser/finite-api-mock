/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// A StatementsApiController binds http requests to an api service and writes the service results to the http response
type StatementsApiController struct {
	service StatementsApiServicer
}

// NewStatementsApiController creates a default api controller
func NewStatementsApiController(s StatementsApiServicer) Router {
	return &StatementsApiController{service: s}
}

// Routes returns all of the api route for the StatementsApiController
func (c *StatementsApiController) Routes() Routes {
	return Routes{
		{
			"GetAccountStatements",
			strings.ToUpper("Get"),
			"/cloud/v1/account/{accountId}/statements",
			c.GetAccountStatements,
		},
	}
}

// GetAccountStatements - Return a list of statements for a given account
func (c *StatementsApiController) GetAccountStatements(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	accountId := params["accountId"]
	mask, _ := strconv.ParseBool(query.Get("mask"))
	enhance, _ := strconv.ParseBool(query.Get("enhance"))
	xTRACEID := r.Header.Get("X-TRACE-ID")
	xTOKEN := r.Header.Get("X-TOKEN")
	result, err := c.service.GetAccountStatements(r.Context(), accountId, mask, enhance, xTRACEID, xTOKEN)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}