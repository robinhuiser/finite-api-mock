// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/bank"
	"github.com/robinhuiser/finite-mock-server/ent/branch"
)

// BranchCreate is the builder for creating a Branch entity.
type BranchCreate struct {
	config
	mutation *BranchMutation
	hooks    []Hook
}

// SetBranchCode sets the "branchCode" field.
func (bc *BranchCreate) SetBranchCode(s string) *BranchCreate {
	bc.mutation.SetBranchCode(s)
	return bc
}

// SetStreetNumber sets the "streetNumber" field.
func (bc *BranchCreate) SetStreetNumber(s string) *BranchCreate {
	bc.mutation.SetStreetNumber(s)
	return bc
}

// SetStreetName sets the "streetName" field.
func (bc *BranchCreate) SetStreetName(s string) *BranchCreate {
	bc.mutation.SetStreetName(s)
	return bc
}

// SetCity sets the "city" field.
func (bc *BranchCreate) SetCity(s string) *BranchCreate {
	bc.mutation.SetCity(s)
	return bc
}

// SetState sets the "state" field.
func (bc *BranchCreate) SetState(s string) *BranchCreate {
	bc.mutation.SetState(s)
	return bc
}

// SetZip sets the "zip" field.
func (bc *BranchCreate) SetZip(s string) *BranchCreate {
	bc.mutation.SetZip(s)
	return bc
}

// SetLatitude sets the "latitude" field.
func (bc *BranchCreate) SetLatitude(f float64) *BranchCreate {
	bc.mutation.SetLatitude(f)
	return bc
}

// SetLongitude sets the "longitude" field.
func (bc *BranchCreate) SetLongitude(f float64) *BranchCreate {
	bc.mutation.SetLongitude(f)
	return bc
}

// SetBranchOwnerID sets the "branch_owner" edge to the Bank entity by ID.
func (bc *BranchCreate) SetBranchOwnerID(id int) *BranchCreate {
	bc.mutation.SetBranchOwnerID(id)
	return bc
}

// SetNillableBranchOwnerID sets the "branch_owner" edge to the Bank entity by ID if the given value is not nil.
func (bc *BranchCreate) SetNillableBranchOwnerID(id *int) *BranchCreate {
	if id != nil {
		bc = bc.SetBranchOwnerID(*id)
	}
	return bc
}

// SetBranchOwner sets the "branch_owner" edge to the Bank entity.
func (bc *BranchCreate) SetBranchOwner(b *Bank) *BranchCreate {
	return bc.SetBranchOwnerID(b.ID)
}

// Mutation returns the BranchMutation object of the builder.
func (bc *BranchCreate) Mutation() *BranchMutation {
	return bc.mutation
}

// Save creates the Branch in the database.
func (bc *BranchCreate) Save(ctx context.Context) (*Branch, error) {
	var (
		err  error
		node *Branch
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BranchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BranchCreate) SaveX(ctx context.Context) *Branch {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (bc *BranchCreate) check() error {
	if _, ok := bc.mutation.BranchCode(); !ok {
		return &ValidationError{Name: "branchCode", err: errors.New("ent: missing required field \"branchCode\"")}
	}
	if _, ok := bc.mutation.StreetNumber(); !ok {
		return &ValidationError{Name: "streetNumber", err: errors.New("ent: missing required field \"streetNumber\"")}
	}
	if _, ok := bc.mutation.StreetName(); !ok {
		return &ValidationError{Name: "streetName", err: errors.New("ent: missing required field \"streetName\"")}
	}
	if _, ok := bc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New("ent: missing required field \"city\"")}
	}
	if _, ok := bc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New("ent: missing required field \"state\"")}
	}
	if v, ok := bc.mutation.State(); ok {
		if err := branch.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Zip(); !ok {
		return &ValidationError{Name: "zip", err: errors.New("ent: missing required field \"zip\"")}
	}
	if _, ok := bc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New("ent: missing required field \"latitude\"")}
	}
	if _, ok := bc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New("ent: missing required field \"longitude\"")}
	}
	return nil
}

func (bc *BranchCreate) sqlSave(ctx context.Context) (*Branch, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BranchCreate) createSpec() (*Branch, *sqlgraph.CreateSpec) {
	var (
		_node = &Branch{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: branch.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: branch.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.BranchCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldBranchCode,
		})
		_node.BranchCode = value
	}
	if value, ok := bc.mutation.StreetNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldStreetNumber,
		})
		_node.StreetNumber = value
	}
	if value, ok := bc.mutation.StreetName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldStreetName,
		})
		_node.StreetName = value
	}
	if value, ok := bc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldCity,
		})
		_node.City = value
	}
	if value, ok := bc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldState,
		})
		_node.State = value
	}
	if value, ok := bc.mutation.Zip(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldZip,
		})
		_node.Zip = value
	}
	if value, ok := bc.mutation.Latitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLatitude,
		})
		_node.Latitude = value
	}
	if value, ok := bc.mutation.Longitude(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLongitude,
		})
		_node.Longitude = value
	}
	if nodes := bc.mutation.BranchOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   branch.BranchOwnerTable,
			Columns: []string{branch.BranchOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.bank_branches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BranchCreateBulk is the builder for creating many Branch entities in bulk.
type BranchCreateBulk struct {
	config
	builders []*BranchCreate
}

// Save creates the Branch entities in the database.
func (bcb *BranchCreateBulk) Save(ctx context.Context) ([]*Branch, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Branch, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BranchMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BranchCreateBulk) SaveX(ctx context.Context) []*Branch {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}