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
	"context"
	"net/http"
	"errors"
)

// StatementsApiService is a service that implents the logic for the StatementsApiServicer
// This service should implement the business logic for every endpoint for the StatementsApi API. 
// Include any external packages or services that will be required by this service.
type StatementsApiService struct {
}

// NewStatementsApiService creates a default api service
func NewStatementsApiService() StatementsApiServicer {
	return &StatementsApiService{}
}

// GetAccountStatements - Return a list of statements for a given account
func (s *StatementsApiService) GetAccountStatements(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetAccountStatements with the required logic for this service method.
	// Add api_statements_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, StatementsList{}) or use other options such as http.Ok ...
	//return Response(200, StatementsList{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountStatements method not implemented")
}

