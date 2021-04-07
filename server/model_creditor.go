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

// Creditor - Account or Entity that receives an exchange
type Creditor struct {
	CreditorType string `json:"creditorType"`

	// When crediting a loan-style account, where to target incoming funds.
	Target string `json:"target,omitempty"`
}
