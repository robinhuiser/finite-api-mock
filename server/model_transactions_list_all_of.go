/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abstract API
 *
 * API version: v1.5.1
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

type TransactionsListAllOf struct {

	// The list of transactions
	Transactions []Transaction `json:"transactions,omitempty"`
}
