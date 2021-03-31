/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abtract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/robinhuiser/finite-api-mock/go"
)

func main() {
	log.Printf("Server started")

	AccountsApiService := openapi.NewAccountsApiService()
	AccountsApiController := openapi.NewAccountsApiController(AccountsApiService)

	CacheApiService := openapi.NewCacheApiService()
	CacheApiController := openapi.NewCacheApiController(CacheApiService)

	CardsApiService := openapi.NewCardsApiService()
	CardsApiController := openapi.NewCardsApiController(CardsApiService)

	EntityApiService := openapi.NewEntityApiService()
	EntityApiController := openapi.NewEntityApiController(EntityApiService)

	ExchangeApiService := openapi.NewExchangeApiService()
	ExchangeApiController := openapi.NewExchangeApiController(ExchangeApiService)

	ProductsApiService := openapi.NewProductsApiService()
	ProductsApiController := openapi.NewProductsApiController(ProductsApiService)

	StatementApiService := openapi.NewStatementApiService()
	StatementApiController := openapi.NewStatementApiController(StatementApiService)

	StatementsApiService := openapi.NewStatementsApiService()
	StatementsApiController := openapi.NewStatementsApiController(StatementsApiService)

	TransactionsApiService := openapi.NewTransactionsApiService()
	TransactionsApiController := openapi.NewTransactionsApiController(TransactionsApiService)

	router := openapi.NewRouter(AccountsApiController, CacheApiController, CardsApiController, EntityApiController, ExchangeApiController, ProductsApiController, StatementApiController, StatementsApiController, TransactionsApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
