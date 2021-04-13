// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/robinhuiser/fca-emu/ent/bank"
	"github.com/robinhuiser/fca-emu/ent/predicate"
)

// BankDelete is the builder for deleting a Bank entity.
type BankDelete struct {
	config
	hooks    []Hook
	mutation *BankMutation
}

// Where adds a new predicate to the BankDelete builder.
func (bd *BankDelete) Where(ps ...predicate.Bank) *BankDelete {
	bd.mutation.predicates = append(bd.mutation.predicates, ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BankDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bd.hooks) == 0 {
		affected, err = bd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bd.mutation = mutation
			affected, err = bd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bd.hooks) - 1; i >= 0; i-- {
			mut = bd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BankDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BankDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bank.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bank.FieldID,
			},
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
}

// BankDeleteOne is the builder for deleting a single Bank entity.
type BankDeleteOne struct {
	bd *BankDelete
}

// Exec executes the deletion query.
func (bdo *BankDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bank.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BankDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
