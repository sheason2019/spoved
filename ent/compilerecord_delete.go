// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sheason2019/spoved/ent/compilerecord"
	"github.com/sheason2019/spoved/ent/predicate"
)

// CompileRecordDelete is the builder for deleting a CompileRecord entity.
type CompileRecordDelete struct {
	config
	hooks    []Hook
	mutation *CompileRecordMutation
}

// Where appends a list predicates to the CompileRecordDelete builder.
func (crd *CompileRecordDelete) Where(ps ...predicate.CompileRecord) *CompileRecordDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CompileRecordDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, CompileRecordMutation](ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CompileRecordDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CompileRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: compilerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: compilerecord.FieldID,
			},
		},
	}
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// CompileRecordDeleteOne is the builder for deleting a single CompileRecord entity.
type CompileRecordDeleteOne struct {
	crd *CompileRecordDelete
}

// Exec executes the deletion query.
func (crdo *CompileRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{compilerecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CompileRecordDeleteOne) ExecX(ctx context.Context) {
	crdo.crd.ExecX(ctx)
}
