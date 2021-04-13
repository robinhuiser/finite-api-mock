// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent/entitycontactpoint"
)

// EntityContactPoint is the model entity for the EntityContactPoint schema.
type EntityContactPoint struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Prefix holds the value of the "prefix" field.
	Prefix int `json:"prefix,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type entitycontactpoint.Type `json:"type,omitempty"`
	// Suffix holds the value of the "suffix" field.
	Suffix int `json:"suffix,omitempty"`
	// Value holds the value of the "value" field.
	Value                        string `json:"value,omitempty"`
	entity_entity_contact_points *uuid.UUID
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntityContactPoint) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case entitycontactpoint.FieldID, entitycontactpoint.FieldPrefix, entitycontactpoint.FieldSuffix:
			values[i] = &sql.NullInt64{}
		case entitycontactpoint.FieldName, entitycontactpoint.FieldType, entitycontactpoint.FieldValue:
			values[i] = &sql.NullString{}
		case entitycontactpoint.ForeignKeys[0]: // entity_entity_contact_points
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type EntityContactPoint", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntityContactPoint fields.
func (ecp *EntityContactPoint) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entitycontactpoint.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ecp.ID = int(value.Int64)
		case entitycontactpoint.FieldPrefix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prefix", values[i])
			} else if value.Valid {
				ecp.Prefix = int(value.Int64)
			}
		case entitycontactpoint.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ecp.Name = value.String
			}
		case entitycontactpoint.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ecp.Type = entitycontactpoint.Type(value.String)
			}
		case entitycontactpoint.FieldSuffix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field suffix", values[i])
			} else if value.Valid {
				ecp.Suffix = int(value.Int64)
			}
		case entitycontactpoint.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ecp.Value = value.String
			}
		case entitycontactpoint.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field entity_entity_contact_points", values[i])
			} else if value != nil {
				ecp.entity_entity_contact_points = value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this EntityContactPoint.
// Note that you need to call EntityContactPoint.Unwrap() before calling this method if this EntityContactPoint
// was returned from a transaction, and the transaction was committed or rolled back.
func (ecp *EntityContactPoint) Update() *EntityContactPointUpdateOne {
	return (&EntityContactPointClient{config: ecp.config}).UpdateOne(ecp)
}

// Unwrap unwraps the EntityContactPoint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ecp *EntityContactPoint) Unwrap() *EntityContactPoint {
	tx, ok := ecp.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntityContactPoint is not a transactional entity")
	}
	ecp.config.driver = tx.drv
	return ecp
}

// String implements the fmt.Stringer.
func (ecp *EntityContactPoint) String() string {
	var builder strings.Builder
	builder.WriteString("EntityContactPoint(")
	builder.WriteString(fmt.Sprintf("id=%v", ecp.ID))
	builder.WriteString(", prefix=")
	builder.WriteString(fmt.Sprintf("%v", ecp.Prefix))
	builder.WriteString(", name=")
	builder.WriteString(ecp.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", ecp.Type))
	builder.WriteString(", suffix=")
	builder.WriteString(fmt.Sprintf("%v", ecp.Suffix))
	builder.WriteString(", value=")
	builder.WriteString(ecp.Value)
	builder.WriteByte(')')
	return builder.String()
}

// EntityContactPoints is a parsable slice of EntityContactPoint.
type EntityContactPoints []*EntityContactPoint

func (ecp EntityContactPoints) config(cfg config) {
	for _i := range ecp {
		ecp[_i].config = cfg
	}
}
