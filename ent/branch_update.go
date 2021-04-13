// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/finite-mock-server/ent/bank"
	"github.com/robinhuiser/finite-mock-server/ent/branch"
	"github.com/robinhuiser/finite-mock-server/ent/predicate"
)

// BranchUpdate is the builder for updating Branch entities.
type BranchUpdate struct {
	config
	hooks    []Hook
	mutation *BranchMutation
}

// Where adds a new predicate for the BranchUpdate builder.
func (bu *BranchUpdate) Where(ps ...predicate.Branch) *BranchUpdate {
	bu.mutation.predicates = append(bu.mutation.predicates, ps...)
	return bu
}

// SetBranchCode sets the "branchCode" field.
func (bu *BranchUpdate) SetBranchCode(s string) *BranchUpdate {
	bu.mutation.SetBranchCode(s)
	return bu
}

// SetStreetNumber sets the "streetNumber" field.
func (bu *BranchUpdate) SetStreetNumber(s string) *BranchUpdate {
	bu.mutation.SetStreetNumber(s)
	return bu
}

// SetStreetName sets the "streetName" field.
func (bu *BranchUpdate) SetStreetName(s string) *BranchUpdate {
	bu.mutation.SetStreetName(s)
	return bu
}

// SetCity sets the "city" field.
func (bu *BranchUpdate) SetCity(s string) *BranchUpdate {
	bu.mutation.SetCity(s)
	return bu
}

// SetState sets the "state" field.
func (bu *BranchUpdate) SetState(s string) *BranchUpdate {
	bu.mutation.SetState(s)
	return bu
}

// SetZip sets the "zip" field.
func (bu *BranchUpdate) SetZip(s string) *BranchUpdate {
	bu.mutation.SetZip(s)
	return bu
}

// SetLatitude sets the "latitude" field.
func (bu *BranchUpdate) SetLatitude(f float64) *BranchUpdate {
	bu.mutation.ResetLatitude()
	bu.mutation.SetLatitude(f)
	return bu
}

// AddLatitude adds f to the "latitude" field.
func (bu *BranchUpdate) AddLatitude(f float64) *BranchUpdate {
	bu.mutation.AddLatitude(f)
	return bu
}

// SetLongitude sets the "longitude" field.
func (bu *BranchUpdate) SetLongitude(f float64) *BranchUpdate {
	bu.mutation.ResetLongitude()
	bu.mutation.SetLongitude(f)
	return bu
}

// AddLongitude adds f to the "longitude" field.
func (bu *BranchUpdate) AddLongitude(f float64) *BranchUpdate {
	bu.mutation.AddLongitude(f)
	return bu
}

// SetBranchOwnerID sets the "branch_owner" edge to the Bank entity by ID.
func (bu *BranchUpdate) SetBranchOwnerID(id int) *BranchUpdate {
	bu.mutation.SetBranchOwnerID(id)
	return bu
}

// SetNillableBranchOwnerID sets the "branch_owner" edge to the Bank entity by ID if the given value is not nil.
func (bu *BranchUpdate) SetNillableBranchOwnerID(id *int) *BranchUpdate {
	if id != nil {
		bu = bu.SetBranchOwnerID(*id)
	}
	return bu
}

// SetBranchOwner sets the "branch_owner" edge to the Bank entity.
func (bu *BranchUpdate) SetBranchOwner(b *Bank) *BranchUpdate {
	return bu.SetBranchOwnerID(b.ID)
}

// Mutation returns the BranchMutation object of the builder.
func (bu *BranchUpdate) Mutation() *BranchMutation {
	return bu.mutation
}

// ClearBranchOwner clears the "branch_owner" edge to the Bank entity.
func (bu *BranchUpdate) ClearBranchOwner() *BranchUpdate {
	bu.mutation.ClearBranchOwner()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BranchUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BranchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BranchUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BranchUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BranchUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BranchUpdate) check() error {
	if v, ok := bu.mutation.State(); ok {
		if err := branch.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (bu *BranchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   branch.Table,
			Columns: branch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: branch.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BranchCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldBranchCode,
		})
	}
	if value, ok := bu.mutation.StreetNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldStreetNumber,
		})
	}
	if value, ok := bu.mutation.StreetName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldStreetName,
		})
	}
	if value, ok := bu.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldCity,
		})
	}
	if value, ok := bu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldState,
		})
	}
	if value, ok := bu.mutation.Zip(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldZip,
		})
	}
	if value, ok := bu.mutation.Latitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLatitude,
		})
	}
	if value, ok := bu.mutation.AddedLatitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLatitude,
		})
	}
	if value, ok := bu.mutation.Longitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLongitude,
		})
	}
	if value, ok := bu.mutation.AddedLongitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLongitude,
		})
	}
	if bu.mutation.BranchOwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BranchOwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{branch.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BranchUpdateOne is the builder for updating a single Branch entity.
type BranchUpdateOne struct {
	config
	hooks    []Hook
	mutation *BranchMutation
}

// SetBranchCode sets the "branchCode" field.
func (buo *BranchUpdateOne) SetBranchCode(s string) *BranchUpdateOne {
	buo.mutation.SetBranchCode(s)
	return buo
}

// SetStreetNumber sets the "streetNumber" field.
func (buo *BranchUpdateOne) SetStreetNumber(s string) *BranchUpdateOne {
	buo.mutation.SetStreetNumber(s)
	return buo
}

// SetStreetName sets the "streetName" field.
func (buo *BranchUpdateOne) SetStreetName(s string) *BranchUpdateOne {
	buo.mutation.SetStreetName(s)
	return buo
}

// SetCity sets the "city" field.
func (buo *BranchUpdateOne) SetCity(s string) *BranchUpdateOne {
	buo.mutation.SetCity(s)
	return buo
}

// SetState sets the "state" field.
func (buo *BranchUpdateOne) SetState(s string) *BranchUpdateOne {
	buo.mutation.SetState(s)
	return buo
}

// SetZip sets the "zip" field.
func (buo *BranchUpdateOne) SetZip(s string) *BranchUpdateOne {
	buo.mutation.SetZip(s)
	return buo
}

// SetLatitude sets the "latitude" field.
func (buo *BranchUpdateOne) SetLatitude(f float64) *BranchUpdateOne {
	buo.mutation.ResetLatitude()
	buo.mutation.SetLatitude(f)
	return buo
}

// AddLatitude adds f to the "latitude" field.
func (buo *BranchUpdateOne) AddLatitude(f float64) *BranchUpdateOne {
	buo.mutation.AddLatitude(f)
	return buo
}

// SetLongitude sets the "longitude" field.
func (buo *BranchUpdateOne) SetLongitude(f float64) *BranchUpdateOne {
	buo.mutation.ResetLongitude()
	buo.mutation.SetLongitude(f)
	return buo
}

// AddLongitude adds f to the "longitude" field.
func (buo *BranchUpdateOne) AddLongitude(f float64) *BranchUpdateOne {
	buo.mutation.AddLongitude(f)
	return buo
}

// SetBranchOwnerID sets the "branch_owner" edge to the Bank entity by ID.
func (buo *BranchUpdateOne) SetBranchOwnerID(id int) *BranchUpdateOne {
	buo.mutation.SetBranchOwnerID(id)
	return buo
}

// SetNillableBranchOwnerID sets the "branch_owner" edge to the Bank entity by ID if the given value is not nil.
func (buo *BranchUpdateOne) SetNillableBranchOwnerID(id *int) *BranchUpdateOne {
	if id != nil {
		buo = buo.SetBranchOwnerID(*id)
	}
	return buo
}

// SetBranchOwner sets the "branch_owner" edge to the Bank entity.
func (buo *BranchUpdateOne) SetBranchOwner(b *Bank) *BranchUpdateOne {
	return buo.SetBranchOwnerID(b.ID)
}

// Mutation returns the BranchMutation object of the builder.
func (buo *BranchUpdateOne) Mutation() *BranchMutation {
	return buo.mutation
}

// ClearBranchOwner clears the "branch_owner" edge to the Bank entity.
func (buo *BranchUpdateOne) ClearBranchOwner() *BranchUpdateOne {
	buo.mutation.ClearBranchOwner()
	return buo
}

// Save executes the query and returns the updated Branch entity.
func (buo *BranchUpdateOne) Save(ctx context.Context) (*Branch, error) {
	var (
		err  error
		node *Branch
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BranchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BranchUpdateOne) SaveX(ctx context.Context) *Branch {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BranchUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BranchUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BranchUpdateOne) check() error {
	if v, ok := buo.mutation.State(); ok {
		if err := branch.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (buo *BranchUpdateOne) sqlSave(ctx context.Context) (_node *Branch, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   branch.Table,
			Columns: branch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: branch.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Branch.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.BranchCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldBranchCode,
		})
	}
	if value, ok := buo.mutation.StreetNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldStreetNumber,
		})
	}
	if value, ok := buo.mutation.StreetName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldStreetName,
		})
	}
	if value, ok := buo.mutation.City(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldCity,
		})
	}
	if value, ok := buo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldState,
		})
	}
	if value, ok := buo.mutation.Zip(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: branch.FieldZip,
		})
	}
	if value, ok := buo.mutation.Latitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLatitude,
		})
	}
	if value, ok := buo.mutation.AddedLatitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLatitude,
		})
	}
	if value, ok := buo.mutation.Longitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLongitude,
		})
	}
	if value, ok := buo.mutation.AddedLongitude(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: branch.FieldLongitude,
		})
	}
	if buo.mutation.BranchOwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BranchOwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Branch{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{branch.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}