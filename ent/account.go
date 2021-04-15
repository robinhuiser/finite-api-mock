// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent/account"
	"github.com/robinhuiser/fca-emu/ent/branch"
	"github.com/robinhuiser/fca-emu/ent/product"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// ParentId holds the value of the "parentId" field.
	ParentId uuid.UUID `json:"parentId,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// DateCreated holds the value of the "dateCreated" field.
	DateCreated time.Time `json:"dateCreated,omitempty"`
	// DateOpened holds the value of the "dateOpened" field.
	DateOpened time.Time `json:"dateOpened,omitempty"`
	// DateLastUpdated holds the value of the "dateLastUpdated" field.
	DateLastUpdated time.Time `json:"dateLastUpdated,omitempty"`
	// DateClosed holds the value of the "dateClosed" field.
	DateClosed time.Time `json:"dateClosed,omitempty"`
	// CurrencyCode holds the value of the "currencyCode" field.
	CurrencyCode string `json:"currencyCode,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// InterestReporting holds the value of the "interestReporting" field.
	InterestReporting bool `json:"interestReporting,omitempty"`
	// CurrentBalance holds the value of the "currentBalance" field.
	CurrentBalance float32 `json:"currentBalance,omitempty"`
	// AvailableBalance holds the value of the "availableBalance" field.
	AvailableBalance float32 `json:"availableBalance,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges           AccountEdges `json:"edges"`
	account_branch  *int
	account_product *int
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Branch holds the value of the branch edge.
	Branch *Branch `json:"branch,omitempty"`
	// Owner holds the value of the owner edge.
	Owner []*Entity `json:"owner,omitempty"`
	// Preference holds the value of the preference edge.
	Preference []*Preference `json:"preference,omitempty"`
	// Routingnumber holds the value of the routingnumber edge.
	Routingnumber []*RoutingNumber `json:"routingnumber,omitempty"`
	// Product holds the value of the product edge.
	Product *Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// BranchOrErr returns the Branch value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) BranchOrErr() (*Branch, error) {
	if e.loadedTypes[0] {
		if e.Branch == nil {
			// The edge branch was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: branch.Label}
		}
		return e.Branch, nil
	}
	return nil, &NotLoadedError{edge: "branch"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) OwnerOrErr() ([]*Entity, error) {
	if e.loadedTypes[1] {
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// PreferenceOrErr returns the Preference value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) PreferenceOrErr() ([]*Preference, error) {
	if e.loadedTypes[2] {
		return e.Preference, nil
	}
	return nil, &NotLoadedError{edge: "preference"}
}

// RoutingnumberOrErr returns the Routingnumber value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) RoutingnumberOrErr() ([]*RoutingNumber, error) {
	if e.loadedTypes[3] {
		return e.Routingnumber, nil
	}
	return nil, &NotLoadedError{edge: "routingnumber"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[4] {
		if e.Product == nil {
			// The edge product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldInterestReporting:
			values[i] = &sql.NullBool{}
		case account.FieldCurrentBalance, account.FieldAvailableBalance:
			values[i] = &sql.NullFloat64{}
		case account.FieldType, account.FieldNumber, account.FieldName, account.FieldTitle, account.FieldCurrencyCode, account.FieldStatus, account.FieldSource, account.FieldURL:
			values[i] = &sql.NullString{}
		case account.FieldDateCreated, account.FieldDateOpened, account.FieldDateLastUpdated, account.FieldDateClosed:
			values[i] = &sql.NullTime{}
		case account.FieldID, account.FieldParentId:
			values[i] = &uuid.UUID{}
		case account.ForeignKeys[0]: // account_branch
			values[i] = &sql.NullInt64{}
		case account.ForeignKeys[1]: // account_product
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case account.FieldNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				a.Number = value.String
			}
		case account.FieldParentId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parentId", values[i])
			} else if value != nil {
				a.ParentId = *value
			}
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		case account.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateCreated", values[i])
			} else if value.Valid {
				a.DateCreated = value.Time
			}
		case account.FieldDateOpened:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateOpened", values[i])
			} else if value.Valid {
				a.DateOpened = value.Time
			}
		case account.FieldDateLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateLastUpdated", values[i])
			} else if value.Valid {
				a.DateLastUpdated = value.Time
			}
		case account.FieldDateClosed:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateClosed", values[i])
			} else if value.Valid {
				a.DateClosed = value.Time
			}
		case account.FieldCurrencyCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currencyCode", values[i])
			} else if value.Valid {
				a.CurrencyCode = value.String
			}
		case account.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = value.String
			}
		case account.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				a.Source = value.String
			}
		case account.FieldInterestReporting:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field interestReporting", values[i])
			} else if value.Valid {
				a.InterestReporting = value.Bool
			}
		case account.FieldCurrentBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field currentBalance", values[i])
			} else if value.Valid {
				a.CurrentBalance = float32(value.Float64)
			}
		case account.FieldAvailableBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field availableBalance", values[i])
			} else if value.Valid {
				a.AvailableBalance = float32(value.Float64)
			}
		case account.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_branch", value)
			} else if value.Valid {
				a.account_branch = new(int)
				*a.account_branch = int(value.Int64)
			}
		case account.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field account_product", value)
			} else if value.Valid {
				a.account_product = new(int)
				*a.account_product = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBranch queries the "branch" edge of the Account entity.
func (a *Account) QueryBranch() *BranchQuery {
	return (&AccountClient{config: a.config}).QueryBranch(a)
}

// QueryOwner queries the "owner" edge of the Account entity.
func (a *Account) QueryOwner() *EntityQuery {
	return (&AccountClient{config: a.config}).QueryOwner(a)
}

// QueryPreference queries the "preference" edge of the Account entity.
func (a *Account) QueryPreference() *PreferenceQuery {
	return (&AccountClient{config: a.config}).QueryPreference(a)
}

// QueryRoutingnumber queries the "routingnumber" edge of the Account entity.
func (a *Account) QueryRoutingnumber() *RoutingNumberQuery {
	return (&AccountClient{config: a.config}).QueryRoutingnumber(a)
}

// QueryProduct queries the "product" edge of the Account entity.
func (a *Account) QueryProduct() *ProductQuery {
	return (&AccountClient{config: a.config}).QueryProduct(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", type=")
	builder.WriteString(a.Type)
	builder.WriteString(", number=")
	builder.WriteString(a.Number)
	builder.WriteString(", parentId=")
	builder.WriteString(fmt.Sprintf("%v", a.ParentId))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", title=")
	builder.WriteString(a.Title)
	builder.WriteString(", dateCreated=")
	builder.WriteString(a.DateCreated.Format(time.ANSIC))
	builder.WriteString(", dateOpened=")
	builder.WriteString(a.DateOpened.Format(time.ANSIC))
	builder.WriteString(", dateLastUpdated=")
	builder.WriteString(a.DateLastUpdated.Format(time.ANSIC))
	builder.WriteString(", dateClosed=")
	builder.WriteString(a.DateClosed.Format(time.ANSIC))
	builder.WriteString(", currencyCode=")
	builder.WriteString(a.CurrencyCode)
	builder.WriteString(", status=")
	builder.WriteString(a.Status)
	builder.WriteString(", source=")
	builder.WriteString(a.Source)
	builder.WriteString(", interestReporting=")
	builder.WriteString(fmt.Sprintf("%v", a.InterestReporting))
	builder.WriteString(", currentBalance=")
	builder.WriteString(fmt.Sprintf("%v", a.CurrentBalance))
	builder.WriteString(", availableBalance=")
	builder.WriteString(fmt.Sprintf("%v", a.AvailableBalance))
	builder.WriteString(", url=")
	builder.WriteString(a.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
