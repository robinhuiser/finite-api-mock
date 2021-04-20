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
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/account"
	"github.com/robinhuiser/fca-emu/util"
)

var clt *ent.Client

// AccountsApiService is a service that implents the logic for the AccountsApiServicer
// This service should implement the business logic for every endpoint for the AccountsApi API.
// Include any external packages or services that will be required by this service.
type AccountsApiService struct {
}

// NewAccountsApiService creates a default api service
func NewAccountsApiService(client *ent.Client) AccountsApiServicer {
	clt = client
	return &AccountsApiService{}
}

// GetAccount - Return a account
func (s *AccountsApiService) GetAccount(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {

	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse("Invalid token")), nil
	}

	// Parse and verify UUID
	u, err := uuid.Parse(accountId)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Account.
		Query().
		Where(account.ID(u)).
		Only(ctx)

	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Retrieve the linked bank
	qba, err := rs.QueryBranch().QueryBranchOwner().Only(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Retrieve the linked branch
	qbr, err := rs.QueryBranch().Only(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Retrieve the linked product
	qpr, err := rs.QueryProduct().Only(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Retrieve the linked account attributes
	atrs, err := rs.QueryPreference().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	preferences := mapPreferences(atrs)

	// Retrieve routing number(s) of the account
	rtns, err := rs.QueryRoutingnumber().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	routingnumbers := mapRoutingNumbers(rtns)

	// Retrieve the owner(s) (entities) of the account
	ents, err := rs.QueryOwner().All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}
	entities := mapEntities(ents, ctx)

	a := Account{
		Id:                rs.ID.String(),
		Type:              rs.Type,
		Number:            isMasked(mask, rs.Number),
		ParentId:          isValidUUID(rs.ParentId.String()),
		Name:              ents[0].Fullname, // TODO: get the primary account holder info here
		Title:             rs.Title,
		DateCreated:       isValidBankDate(rs.DateCreated.Format(util.APIDateFormat)),
		DateOpened:        isValidBankDate(rs.DateOpened.Format(util.APIDateFormat)),
		DateLastUpdated:   isValidBankDate(rs.DateLastUpdated.Format(util.APIDateFormat)),
		DateClosed:        isValidBankDate(rs.DateClosed.Format(util.APIDateFormat)),
		CurrencyCode:      rs.CurrencyCode,
		Status:            rs.Status,
		Source:            rs.Source,
		InterestReporting: rs.InterestReporting,
		Balances: Balances{
			AvailableBalance: rs.AvailableBalance,
			CurrentBalance:   rs.CurrentBalance,
		},
		URI: FiniteUri{
			URL: rs.URL,
		},
		Bank: Bank{
			BranchCode: qbr.BranchCode,
			BankName:   qba.BankName,
			Swift:      qba.Swift,
			URI: FiniteUri{
				URL: qba.URL,
			},
		},
		Entities:       entities,
		Preferences:    preferences,
		Routingnumbers: routingnumbers,
		Product: Product{
			Number:      qpr.Name,
			Type:        string(qpr.Type),
			TypeName:    qpr.TypeName,
			SubType:     qpr.SubType,
			SubTypeName: qpr.SubTypeName,
			URI: FiniteUri{
				URL: qpr.URL,
			},
		},
	}

	return Response(200, a), nil
}

// GetAccountBalances - Return a accounts balances
func (s *AccountsApiService) GetAccountBalances(ctx context.Context, accountId string, mask bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse("Invalid token")), nil
	}

	// Parse and verify UUID
	u, err := uuid.Parse(accountId)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// Perform the search
	rs, err := clt.Account.
		Query().
		Where(account.ID(u)).
		Only(ctx)

	// Error if none or more than one results are returned
	if err != nil {
		return Response(404, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	b := Balances{
		AvailableBalance: rs.AvailableBalance,
		CurrentBalance:   rs.CurrentBalance,
	}

	return Response(200, b), nil
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

// Helper functions for mapping
func mapPreferences(attributes []*ent.Preference) []Attribute {
	preferences := []Attribute{}
	for _, pref := range attributes {
		attr := Attribute{
			Name:  pref.Name,
			Value: pref.Value,
		}
		preferences = append(preferences, attr)
	}
	return preferences
}

func mapRoutingNumbers(routingnumbers []*ent.RoutingNumber) []RoutingNumber {
	rtns := []RoutingNumber{}
	for _, rtn := range routingnumbers {
		rn := RoutingNumber{
			Number: rtn.Number,
			Type:   string(rtn.Type),
		}
		rtns = append(rtns, rn)
	}
	return rtns
}
