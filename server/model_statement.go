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

// Statement - Account Statement
type Statement struct {

	// The ID of this Statement.
	Id string `json:"id,omitempty"`

	// The ID of the Account for this Statement.
	AccountId string `json:"accountId,omitempty"`

	// The name of the document for this Statement.
	Name string `json:"name,omitempty"`

	// List of attributes for a binary item, e.g description, creation time...
	Attributes []Attribute `json:"attributes,omitempty"`

	// ISO 8601 date of creation
	DateCreated string `json:"dateCreated,omitempty"`

	// ISO 8601 date of revision
	DateRevised string `json:"dateRevised,omitempty"`

	Items BinaryItemList `json:"items,omitempty"`
}
