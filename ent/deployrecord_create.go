// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sheason2019/spoved/ent/compilerecord"
	"github.com/sheason2019/spoved/ent/deployrecord"
	"github.com/sheason2019/spoved/ent/user"
)

// DeployRecordCreate is the builder for creating a DeployRecord entity.
type DeployRecordCreate struct {
	config
	mutation *DeployRecordMutation
	hooks    []Hook
}

// SetImage sets the "image" field.
func (drc *DeployRecordCreate) SetImage(s string) *DeployRecordCreate {
	drc.mutation.SetImage(s)
	return drc
}

// SetCreatedAt sets the "created_at" field.
func (drc *DeployRecordCreate) SetCreatedAt(t time.Time) *DeployRecordCreate {
	drc.mutation.SetCreatedAt(t)
	return drc
}

// SetStatusCode sets the "status_code" field.
func (drc *DeployRecordCreate) SetStatusCode(i int) *DeployRecordCreate {
	drc.mutation.SetStatusCode(i)
	return drc
}

// SetContainerHash sets the "container_hash" field.
func (drc *DeployRecordCreate) SetContainerHash(s string) *DeployRecordCreate {
	drc.mutation.SetContainerHash(s)
	return drc
}

// AddOperatorIDs adds the "operator" edge to the User entity by IDs.
func (drc *DeployRecordCreate) AddOperatorIDs(ids ...int) *DeployRecordCreate {
	drc.mutation.AddOperatorIDs(ids...)
	return drc
}

// AddOperator adds the "operator" edges to the User entity.
func (drc *DeployRecordCreate) AddOperator(u ...*User) *DeployRecordCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return drc.AddOperatorIDs(ids...)
}

// AddCompileRecordIDs adds the "compile_record" edge to the CompileRecord entity by IDs.
func (drc *DeployRecordCreate) AddCompileRecordIDs(ids ...int) *DeployRecordCreate {
	drc.mutation.AddCompileRecordIDs(ids...)
	return drc
}

// AddCompileRecord adds the "compile_record" edges to the CompileRecord entity.
func (drc *DeployRecordCreate) AddCompileRecord(c ...*CompileRecord) *DeployRecordCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return drc.AddCompileRecordIDs(ids...)
}

// Mutation returns the DeployRecordMutation object of the builder.
func (drc *DeployRecordCreate) Mutation() *DeployRecordMutation {
	return drc.mutation
}

// Save creates the DeployRecord in the database.
func (drc *DeployRecordCreate) Save(ctx context.Context) (*DeployRecord, error) {
	return withHooks[*DeployRecord, DeployRecordMutation](ctx, drc.sqlSave, drc.mutation, drc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (drc *DeployRecordCreate) SaveX(ctx context.Context) *DeployRecord {
	v, err := drc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drc *DeployRecordCreate) Exec(ctx context.Context) error {
	_, err := drc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drc *DeployRecordCreate) ExecX(ctx context.Context) {
	if err := drc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drc *DeployRecordCreate) check() error {
	if _, ok := drc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "DeployRecord.image"`)}
	}
	if _, ok := drc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeployRecord.created_at"`)}
	}
	if _, ok := drc.mutation.StatusCode(); !ok {
		return &ValidationError{Name: "status_code", err: errors.New(`ent: missing required field "DeployRecord.status_code"`)}
	}
	if _, ok := drc.mutation.ContainerHash(); !ok {
		return &ValidationError{Name: "container_hash", err: errors.New(`ent: missing required field "DeployRecord.container_hash"`)}
	}
	return nil
}

func (drc *DeployRecordCreate) sqlSave(ctx context.Context) (*DeployRecord, error) {
	if err := drc.check(); err != nil {
		return nil, err
	}
	_node, _spec := drc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	drc.mutation.id = &_node.ID
	drc.mutation.done = true
	return _node, nil
}

func (drc *DeployRecordCreate) createSpec() (*DeployRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &DeployRecord{config: drc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deployrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deployrecord.FieldID,
			},
		}
	)
	if value, ok := drc.mutation.Image(); ok {
		_spec.SetField(deployrecord.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := drc.mutation.CreatedAt(); ok {
		_spec.SetField(deployrecord.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := drc.mutation.StatusCode(); ok {
		_spec.SetField(deployrecord.FieldStatusCode, field.TypeInt, value)
		_node.StatusCode = value
	}
	if value, ok := drc.mutation.ContainerHash(); ok {
		_spec.SetField(deployrecord.FieldContainerHash, field.TypeString, value)
		_node.ContainerHash = value
	}
	if nodes := drc.mutation.OperatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   deployrecord.OperatorTable,
			Columns: deployrecord.OperatorPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := drc.mutation.CompileRecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   deployrecord.CompileRecordTable,
			Columns: deployrecord.CompileRecordPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: compilerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeployRecordCreateBulk is the builder for creating many DeployRecord entities in bulk.
type DeployRecordCreateBulk struct {
	config
	builders []*DeployRecordCreate
}

// Save creates the DeployRecord entities in the database.
func (drcb *DeployRecordCreateBulk) Save(ctx context.Context) ([]*DeployRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(drcb.builders))
	nodes := make([]*DeployRecord, len(drcb.builders))
	mutators := make([]Mutator, len(drcb.builders))
	for i := range drcb.builders {
		func(i int, root context.Context) {
			builder := drcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeployRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, drcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, drcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, drcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (drcb *DeployRecordCreateBulk) SaveX(ctx context.Context) []*DeployRecord {
	v, err := drcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drcb *DeployRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := drcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drcb *DeployRecordCreateBulk) ExecX(ctx context.Context) {
	if err := drcb.Exec(ctx); err != nil {
		panic(err)
	}
}