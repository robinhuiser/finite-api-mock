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

// AccountsApiService is a service that implents the logic for the AccountsApiServicer
// This service should implement the business logic for every endpoint for the AccountsApi API.
// Include any external packages or services that will be required by this service.
type AccountsApiService struct {
}

// NewAccountsApiService creates a default api service
func NewAccountsApiService() AccountsApiServicer {
	return &AccountsApiService{}
}

// GetAccount - Return a account
func (s *AccountsApiService) GetAccount(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetAccount with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Account{}) or use other options such as http.Ok ...
	//return Response(200, Account{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccount method not implemented")
}

// GetAccountBalances - Return a accounts balances
func (s *AccountsApiService) GetAccountBalances(ctx context.Context, accountId string, mask bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetAccountBalances with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Balances{}) or use other options such as http.Ok ...
	//return Response(200, Balances{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountBalances method not implemented")
}

// GetAccountDetails - Return a accounts details
func (s *AccountsApiService) GetAccountDetails(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetAccountDetails with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Account{}) or use other options such as http.Ok ...
	//return Response(200, Account{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountDetails method not implemented")
}

// GetEntityAccountsList - Return list a of accounts by entity
func (s *AccountsApiService) GetEntityAccountsList(ctx context.Context, entityId string, fields string, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetEntityAccountsList with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, AccountsList{}) or use other options such as http.Ok ...
	//return Response(200, AccountsList{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetEntityAccountsList method not implemented")
}

// PostEntityAccountsList - Return list a of accounts based on a entity search
func (s *AccountsApiService) PostEntityAccountsList(ctx context.Context, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string, searchFilter []SearchFilter) (ImplResponse, error) {
	// TODO - update PostEntityAccountsList with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, AccountsList{}) or use other options such as http.Ok ...
	//return Response(200, AccountsList{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostEntityAccountsList method not implemented")
}

// PutAccount - Update a account
func (s *AccountsApiService) PutAccount(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string, account Account) (ImplResponse, error) {
	// TODO - update PutAccount with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Account{}) or use other options such as http.Ok ...
	//return Response(200, Account{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PutAccount method not implemented")
}

// SearchAccounts - Search for accounts
func (s *AccountsApiService) SearchAccounts(ctx context.Context, fields string, limit int32, cursor string, mask bool, enhance bool, xTRACEID string, xTOKEN string, searchFilter []SearchFilter) (ImplResponse, error) {
	// TODO - update SearchAccounts with the required logic for this service method.
	// Add api_accounts_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, AccountsList{}) or use other options such as http.Ok ...
	//return Response(200, AccountsList{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SearchAccounts method not implemented")
}
