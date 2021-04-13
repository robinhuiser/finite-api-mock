// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/fca-emu/ent/entitycontactpoint"
	"github.com/robinhuiser/fca-emu/ent/predicate"
)

// EntityContactPointUpdate is the builder for updating EntityContactPoint entities.
type EntityContactPointUpdate struct {
	config
	hooks    []Hook
	mutation *EntityContactPointMutation
}

// Where adds a new predicate for the EntityContactPointUpdate builder.
func (ecpu *EntityContactPointUpdate) Where(ps ...predicate.EntityContactPoint) *EntityContactPointUpdate {
	ecpu.mutation.predicates = append(ecpu.mutation.predicates, ps...)
	return ecpu
}

// SetPrefix sets the "prefix" field.
func (ecpu *EntityContactPointUpdate) SetPrefix(i int) *EntityContactPointUpdate {
	ecpu.mutation.ResetPrefix()
	ecpu.mutation.SetPrefix(i)
	return ecpu
}

// SetNillablePrefix sets the "prefix" field if the given value is not nil.
func (ecpu *EntityContactPointUpdate) SetNillablePrefix(i *int) *EntityContactPointUpdate {
	if i != nil {
		ecpu.SetPrefix(*i)
	}
	return ecpu
}

// AddPrefix adds i to the "prefix" field.
func (ecpu *EntityContactPointUpdate) AddPrefix(i int) *EntityContactPointUpdate {
	ecpu.mutation.AddPrefix(i)
	return ecpu
}

// ClearPrefix clears the value of the "prefix" field.
func (ecpu *EntityContactPointUpdate) ClearPrefix() *EntityContactPointUpdate {
	ecpu.mutation.ClearPrefix()
	return ecpu
}

// SetName sets the "name" field.
func (ecpu *EntityContactPointUpdate) SetName(s string) *EntityContactPointUpdate {
	ecpu.mutation.SetName(s)
	return ecpu
}

// SetType sets the "type" field.
func (ecpu *EntityContactPointUpdate) SetType(e entitycontactpoint.Type) *EntityContactPointUpdate {
	ecpu.mutation.SetType(e)
	return ecpu
}

// SetSuffix sets the "suffix" field.
func (ecpu *EntityContactPointUpdate) SetSuffix(i int) *EntityContactPointUpdate {
	ecpu.mutation.ResetSuffix()
	ecpu.mutation.SetSuffix(i)
	return ecpu
}

// SetNillableSuffix sets the "suffix" field if the given value is not nil.
func (ecpu *EntityContactPointUpdate) SetNillableSuffix(i *int) *EntityContactPointUpdate {
	if i != nil {
		ecpu.SetSuffix(*i)
	}
	return ecpu
}

// AddSuffix adds i to the "suffix" field.
func (ecpu *EntityContactPointUpdate) AddSuffix(i int) *EntityContactPointUpdate {
	ecpu.mutation.AddSuffix(i)
	return ecpu
}

// ClearSuffix clears the value of the "suffix" field.
func (ecpu *EntityContactPointUpdate) ClearSuffix() *EntityContactPointUpdate {
	ecpu.mutation.ClearSuffix()
	return ecpu
}

// SetValue sets the "value" field.
func (ecpu *EntityContactPointUpdate) SetValue(s string) *EntityContactPointUpdate {
	ecpu.mutation.SetValue(s)
	return ecpu
}

// Mutation returns the EntityContactPointMutation object of the builder.
func (ecpu *EntityContactPointUpdate) Mutation() *EntityContactPointMutation {
	return ecpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecpu *EntityContactPointUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ecpu.hooks) == 0 {
		if err = ecpu.check(); err != nil {
			return 0, err
		}
		affected, err = ecpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityContactPointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ecpu.check(); err != nil {
				return 0, err
			}
			ecpu.mutation = mutation
			affected, err = ecpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ecpu.hooks) - 1; i >= 0; i-- {
			mut = ecpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ecpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecpu *EntityContactPointUpdate) SaveX(ctx context.Context) int {
	affected, err := ecpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecpu *EntityContactPointUpdate) Exec(ctx context.Context) error {
	_, err := ecpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecpu *EntityContactPointUpdate) ExecX(ctx context.Context) {
	if err := ecpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecpu *EntityContactPointUpdate) check() error {
	if v, ok := ecpu.mutation.GetType(); ok {
		if err := entitycontactpoint.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (ecpu *EntityContactPointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitycontactpoint.Table,
			Columns: entitycontactpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitycontactpoint.FieldID,
			},
		},
	}
	if ps := ecpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecpu.mutation.Prefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldPrefix,
		})
	}
	if value, ok := ecpu.mutation.AddedPrefix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldPrefix,
		})
	}
	if ecpu.mutation.PrefixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: entitycontactpoint.FieldPrefix,
		})
	}
	if value, ok := ecpu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitycontactpoint.FieldName,
		})
	}
	if value, ok := ecpu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entitycontactpoint.FieldType,
		})
	}
	if value, ok := ecpu.mutation.Suffix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldSuffix,
		})
	}
	if value, ok := ecpu.mutation.AddedSuffix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldSuffix,
		})
	}
	if ecpu.mutation.SuffixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: entitycontactpoint.FieldSuffix,
		})
	}
	if value, ok := ecpu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitycontactpoint.FieldValue,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitycontactpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// EntityContactPointUpdateOne is the builder for updating a single EntityContactPoint entity.
type EntityContactPointUpdateOne struct {
	config
	hooks    []Hook
	mutation *EntityContactPointMutation
}

// SetPrefix sets the "prefix" field.
func (ecpuo *EntityContactPointUpdateOne) SetPrefix(i int) *EntityContactPointUpdateOne {
	ecpuo.mutation.ResetPrefix()
	ecpuo.mutation.SetPrefix(i)
	return ecpuo
}

// SetNillablePrefix sets the "prefix" field if the given value is not nil.
func (ecpuo *EntityContactPointUpdateOne) SetNillablePrefix(i *int) *EntityContactPointUpdateOne {
	if i != nil {
		ecpuo.SetPrefix(*i)
	}
	return ecpuo
}

// AddPrefix adds i to the "prefix" field.
func (ecpuo *EntityContactPointUpdateOne) AddPrefix(i int) *EntityContactPointUpdateOne {
	ecpuo.mutation.AddPrefix(i)
	return ecpuo
}

// ClearPrefix clears the value of the "prefix" field.
func (ecpuo *EntityContactPointUpdateOne) ClearPrefix() *EntityContactPointUpdateOne {
	ecpuo.mutation.ClearPrefix()
	return ecpuo
}

// SetName sets the "name" field.
func (ecpuo *EntityContactPointUpdateOne) SetName(s string) *EntityContactPointUpdateOne {
	ecpuo.mutation.SetName(s)
	return ecpuo
}

// SetType sets the "type" field.
func (ecpuo *EntityContactPointUpdateOne) SetType(e entitycontactpoint.Type) *EntityContactPointUpdateOne {
	ecpuo.mutation.SetType(e)
	return ecpuo
}

// SetSuffix sets the "suffix" field.
func (ecpuo *EntityContactPointUpdateOne) SetSuffix(i int) *EntityContactPointUpdateOne {
	ecpuo.mutation.ResetSuffix()
	ecpuo.mutation.SetSuffix(i)
	return ecpuo
}

// SetNillableSuffix sets the "suffix" field if the given value is not nil.
func (ecpuo *EntityContactPointUpdateOne) SetNillableSuffix(i *int) *EntityContactPointUpdateOne {
	if i != nil {
		ecpuo.SetSuffix(*i)
	}
	return ecpuo
}

// AddSuffix adds i to the "suffix" field.
func (ecpuo *EntityContactPointUpdateOne) AddSuffix(i int) *EntityContactPointUpdateOne {
	ecpuo.mutation.AddSuffix(i)
	return ecpuo
}

// ClearSuffix clears the value of the "suffix" field.
func (ecpuo *EntityContactPointUpdateOne) ClearSuffix() *EntityContactPointUpdateOne {
	ecpuo.mutation.ClearSuffix()
	return ecpuo
}

// SetValue sets the "value" field.
func (ecpuo *EntityContactPointUpdateOne) SetValue(s string) *EntityContactPointUpdateOne {
	ecpuo.mutation.SetValue(s)
	return ecpuo
}

// Mutation returns the EntityContactPointMutation object of the builder.
func (ecpuo *EntityContactPointUpdateOne) Mutation() *EntityContactPointMutation {
	return ecpuo.mutation
}

// Save executes the query and returns the updated EntityContactPoint entity.
func (ecpuo *EntityContactPointUpdateOne) Save(ctx context.Context) (*EntityContactPoint, error) {
	var (
		err  error
		node *EntityContactPoint
	)
	if len(ecpuo.hooks) == 0 {
		if err = ecpuo.check(); err != nil {
			return nil, err
		}
		node, err = ecpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityContactPointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ecpuo.check(); err != nil {
				return nil, err
			}
			ecpuo.mutation = mutation
			node, err = ecpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ecpuo.hooks) - 1; i >= 0; i-- {
			mut = ecpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ecpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ecpuo *EntityContactPointUpdateOne) SaveX(ctx context.Context) *EntityContactPoint {
	node, err := ecpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecpuo *EntityContactPointUpdateOne) Exec(ctx context.Context) error {
	_, err := ecpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecpuo *EntityContactPointUpdateOne) ExecX(ctx context.Context) {
	if err := ecpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecpuo *EntityContactPointUpdateOne) check() error {
	if v, ok := ecpuo.mutation.GetType(); ok {
		if err := entitycontactpoint.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (ecpuo *EntityContactPointUpdateOne) sqlSave(ctx context.Context) (_node *EntityContactPoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entitycontactpoint.Table,
			Columns: entitycontactpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entitycontactpoint.FieldID,
			},
		},
	}
	id, ok := ecpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing EntityContactPoint.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ecpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecpuo.mutation.Prefix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldPrefix,
		})
	}
	if value, ok := ecpuo.mutation.AddedPrefix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldPrefix,
		})
	}
	if ecpuo.mutation.PrefixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: entitycontactpoint.FieldPrefix,
		})
	}
	if value, ok := ecpuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitycontactpoint.FieldName,
		})
	}
	if value, ok := ecpuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: entitycontactpoint.FieldType,
		})
	}
	if value, ok := ecpuo.mutation.Suffix(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldSuffix,
		})
	}
	if value, ok := ecpuo.mutation.AddedSuffix(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entitycontactpoint.FieldSuffix,
		})
	}
	if ecpuo.mutation.SuffixCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: entitycontactpoint.FieldSuffix,
		})
	}
	if value, ok := ecpuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entitycontactpoint.FieldValue,
		})
	}
	_node = &EntityContactPoint{config: ecpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitycontactpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
