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

// AccountDebtor - An Account That is Debited
type AccountDebtor struct {

	// Unique composite identifier of the account in the system of record
	Id string `json:"id"`

	// The type of account as enumerated by the system of record
	Type string `json:"type,omitempty"`

	// The account number as represented in the system of record (masked if mask=true)
	Number string `json:"number,omitempty"`

	// Unique identifier of the parent account, in the case of a sub account
	ParentId string `json:"parentId,omitempty"`

	// List of entities this account is associated with
	Entities []Entity `json:"entities,omitempty"`

	// The name of the primary account holder
	Name string `json:"name,omitempty"`

	// The given title of an account
	Title string `json:"title,omitempty"`

	// The international bank account number
	Iban string `json:"iban,omitempty"`

	// The routing numbers of this account
	Routingnumbers []RoutingNumber `json:"routingnumbers,omitempty"`

	Balances Balances `json:"balances,omitempty"`

	// ISO 6801 date when the account was created
	DateCreated string `json:"dateCreated,omitempty"`

	// ISO 6801 date when the account was opened or activated
	DateOpened string `json:"dateOpened,omitempty"`

	// ISO 6801 date when the account was last updated
	DateLastUpdated string `json:"dateLastUpdated,omitempty"`

	// ISO 6801 date when the account was closed or deactivated
	DateClosed string `json:"dateClosed,omitempty"`

	// The currency code of the account
	CurrencyCode string `json:"currencyCode,omitempty"`

	// The current status of the account
	Status string `json:"status,omitempty"`

	// The source of the account, where the account is located
	Source string `json:"source,omitempty"`

	// Indicates if interest is reportable on this account.
	InterestReporting bool `json:"interestReporting,omitempty"`

	CreditInformation CreditInformation `json:"creditInformation,omitempty"`

	LoanInterest Interest `json:"loanInterest,omitempty"`

	CreditInterest Interest `json:"creditInterest,omitempty"`

	// Unique preferences of the account
	Preferences []Attribute `json:"preferences,omitempty"`

	Bank Bank `json:"bank,omitempty"`

	// List of owners of the account
	Owners []Relationship `json:"owners,omitempty"`

	Product Product `json:"product,omitempty"`

	URI FiniteUri `json:"URI,omitempty"`
}
