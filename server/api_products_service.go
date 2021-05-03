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
	"fmt"

	"github.com/robinhuiser/fca-emu/ent"
	"github.com/robinhuiser/fca-emu/ent/product"
)

// ProductsApiService is a service that implents the logic for the ProductsApiServicer
// This service should implement the business logic for every endpoint for the ProductsApi API.
// Include any external packages or services that will be required by this service.
type ProductsApiService struct {
}

// NewProductsApiService creates a default api service
func NewProductsApiService() ProductsApiServicer {
	return &ProductsApiService{}
}

// GetProducts - Return a list of products
func (s *ProductsApiService) GetProducts(ctx context.Context, productType string, limit int32, cursor string, enhance bool, xTRACEID string, xTOKEN string) (ImplResponse, error) {
	// Validate X-Token
	if !isValidSecret(xTOKEN) {
		return Response(401, setErrorResponse(INVALID_TOKEN_MSG)), nil
	}

	// Mandatory parameter
	if len(productType) == 0 {
		return Response(400, setErrorResponse("Mandatory parameter productType not specified")), nil
	}

	// Perform the search
	rs, err := clt.Product.
		Query().
		Where(product.TypeEQ(product.Type(productType))).
		All(ctx)
	if err != nil {
		return Response(500, setErrorResponse(fmt.Sprintf("%v", err))), nil
	}

	// So we have any products matching the specified productType?
	if len(rs) == 0 {
		return Response(404, setErrorResponse(fmt.Sprintf("No product found of type %s", productType))), nil
	}
	products := mapProducts(rs)

	p := ProductsList{
		Status:     true,
		TotalItems: int32(len(rs)),
		Products:   products,
	}

	return Response(200, p), nil
}

// Helper functions for mapping
func mapProducts(products []*ent.Product) []Product {
	prds := []Product{}
	for _, pr := range products {
		pr := Product{
			Number:      pr.Name,
			Type:        string(pr.Type),
			TypeName:    pr.TypeName,
			SubType:     pr.SubType,
			SubTypeName: pr.SubTypeName,
			URI: FiniteUri{
				URL: pr.URL,
			},
		}
		prds = append(prds, pr)
	}
	return prds
}
