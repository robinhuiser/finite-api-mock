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
	"errors"
	"net/http"
)

// StatementApiService is a service that implents the logic for the StatementApiServicer
// This service should implement the business logic for every endpoint for the StatementApi API.
// Include any external packages or services that will be required by this service.
type StatementApiService struct {
}

// NewStatementApiService creates a default api service
func NewStatementApiService() StatementApiServicer {
	return &StatementApiService{}
}

// GetAccountStatementItem - Return a binary item for a statement
func (s *StatementApiService) GetAccountStatementItem(ctx context.Context, accountId string, statementId string, itemId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountStatementItem method not implemented")
}

// GetAccountStatementItems - Return the list of content items associated with a statement
func (s *StatementApiService) GetAccountStatementItems(ctx context.Context, accountId string, statementId string, inline bool, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountStatementItems method not implemented")
}

// GetStatementItem - Return a binary item for a statement
func (s *StatementApiService) GetStatementItem(ctx context.Context, statementId string, itemId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("GetStatementItem method not implemented")
}

// GetStatementItems - Return the list of content items associated with a statement
func (s *StatementApiService) GetStatementItems(ctx context.Context, statementId string, inline bool, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {

	return Response(http.StatusNotImplemented, nil), errors.New("GetStatementItems method not implemented")
}
