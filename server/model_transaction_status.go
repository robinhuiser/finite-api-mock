/*
 * Cloud API
 *
 * The public facing API through which connectors are exposed as a single abstract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finite

// TransactionStatus : The status of the transaction
type TransactionStatus string

// List of TransactionStatus
const (
	PENDING TransactionStatus = "PENDING"
	POSTED  TransactionStatus = "POSTED"
)
