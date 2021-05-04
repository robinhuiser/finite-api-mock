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

// Interest - Interest object
type Interest struct {

	// The effective interest rate % of a loan or saving account
	Rate float32 `json:"rate,omitempty"`

	// The measure of terms
	TermsUnit string `json:"termsUnit,omitempty"`

	// The number of terms agreed
	Terms int32 `json:"terms,omitempty"`
}
