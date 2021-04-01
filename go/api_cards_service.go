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

// CardsApiService is a service that implents the logic for the CardsApiServicer
// This service should implement the business logic for every endpoint for the CardsApi API. 
// Include any external packages or services that will be required by this service.
type CardsApiService struct {
}

// NewCardsApiService creates a default api service
func NewCardsApiService() CardsApiServicer {
	return &CardsApiService{}
}

// GetAccountCards - Return a accounts cards
func (s *CardsApiService) GetAccountCards(ctx context.Context, accountId string, mask bool, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// TODO - update GetAccountCards with the required logic for this service method.
	// Add api_cards_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CardsList{}) or use other options such as http.Ok ...
	//return Response(200, CardsList{}), nil

	//TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(401, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(400, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(404, ErrorResponse{}), nil

	//TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	//return Response(500, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAccountCards method not implemented")
}
