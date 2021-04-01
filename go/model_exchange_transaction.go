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

// ExchangeTransaction - Exchange Transaction
type ExchangeTransaction struct {

	// Unique identifier for this exchange transaction
	Id string `json:"id,omitempty"`

	// The type of exchange or money movement
	Type string `json:"type,omitempty"`

	Debtor Debtor `json:"debtor,omitempty"`

	Creditor Creditor `json:"creditor,omitempty"`

	IntermediateInstitutions []IntermediateInstitution `json:"intermediateInstitutions,omitempty"`

	RemittanceInformation RemittanceInformation `json:"remittanceInformation,omitempty"`

	// Indicate if this transaction is part of a batch
	IsBatch bool `json:"isBatch,omitempty"`

	// Indicates the priority of this transaction
	Priority string `json:"priority,omitempty"`

	// ISO 8601 date of execution
	ExecutionDate string `json:"executionDate,omitempty"`

	// Indicates if this is a recurring transaction
	IsRecurring bool `json:"isRecurring,omitempty"`

	RecurringSchedule Schedule `json:"recurringSchedule,omitempty"`

	// The amount of the transaction
	Amount float32 `json:"amount,omitempty"`

	Fee Fee `json:"fee,omitempty"`

	URI FiniteUri `json:"URI,omitempty"`
}