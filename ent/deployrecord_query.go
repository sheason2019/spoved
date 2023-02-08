// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sheason2019/spoved/ent/compilerecord"
	"github.com/sheason2019/spoved/ent/deployrecord"
	"github.com/sheason2019/spoved/ent/predicate"
	"github.com/sheason2019/spoved/ent/user"
)

// DeployRecordQuery is the builder for querying DeployRecord entities.
type DeployRecordQuery struct {
	config
	ctx               *QueryContext
	order             []OrderFunc
	inters            []Interceptor
	predicates        []predicate.DeployRecord
	withOperator      *UserQuery
	withCompileRecord *CompileRecordQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeployRecordQuery builder.
func (drq *DeployRecordQuery) Where(ps ...predicate.DeployRecord) *DeployRecordQuery {
	drq.predicates = append(drq.predicates, ps...)
	return drq
}

// Limit the number of records to be returned by this query.
func (drq *DeployRecordQuery) Limit(limit int) *DeployRecordQuery {
	drq.ctx.Limit = &limit
	return drq
}

// Offset to start from.
func (drq *DeployRecordQuery) Offset(offset int) *DeployRecordQuery {
	drq.ctx.Offset = &offset
	return drq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (drq *DeployRecordQuery) Unique(unique bool) *DeployRecordQuery {
	drq.ctx.Unique = &unique
	return drq
}

// Order specifies how the records should be ordered.
func (drq *DeployRecordQuery) Order(o ...OrderFunc) *DeployRecordQuery {
	drq.order = append(drq.order, o...)
	return drq
}

// QueryOperator chains the current query on the "operator" edge.
func (drq *DeployRecordQuery) QueryOperator() *UserQuery {
	query := (&UserClient{config: drq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deployrecord.Table, deployrecord.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, deployrecord.OperatorTable, deployrecord.OperatorPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCompileRecord chains the current query on the "compile_record" edge.
func (drq *DeployRecordQuery) QueryCompileRecord() *CompileRecordQuery {
	query := (&CompileRecordClient{config: drq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deployrecord.Table, deployrecord.FieldID, selector),
			sqlgraph.To(compilerecord.Table, compilerecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, deployrecord.CompileRecordTable, deployrecord.CompileRecordPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeployRecord entity from the query.
// Returns a *NotFoundError when no DeployRecord was found.
func (drq *DeployRecordQuery) First(ctx context.Context) (*DeployRecord, error) {
	nodes, err := drq.Limit(1).All(setContextOp(ctx, drq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deployrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drq *DeployRecordQuery) FirstX(ctx context.Context) *DeployRecord {
	node, err := drq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeployRecord ID from the query.
// Returns a *NotFoundError when no DeployRecord ID was found.
func (drq *DeployRecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(1).IDs(setContextOp(ctx, drq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deployrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (drq *DeployRecordQuery) FirstIDX(ctx context.Context) int {
	id, err := drq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeployRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeployRecord entity is found.
// Returns a *NotFoundError when no DeployRecord entities are found.
func (drq *DeployRecordQuery) Only(ctx context.Context) (*DeployRecord, error) {
	nodes, err := drq.Limit(2).All(setContextOp(ctx, drq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deployrecord.Label}
	default:
		return nil, &NotSingularError{deployrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drq *DeployRecordQuery) OnlyX(ctx context.Context) *DeployRecord {
	node, err := drq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeployRecord ID in the query.
// Returns a *NotSingularError when more than one DeployRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (drq *DeployRecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(2).IDs(setContextOp(ctx, drq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deployrecord.Label}
	default:
		err = &NotSingularError{deployrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drq *DeployRecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := drq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeployRecords.
func (drq *DeployRecordQuery) All(ctx context.Context) ([]*DeployRecord, error) {
	ctx = setContextOp(ctx, drq.ctx, "All")
	if err := drq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DeployRecord, *DeployRecordQuery]()
	return withInterceptors[[]*DeployRecord](ctx, drq, qr, drq.inters)
}

// AllX is like All, but panics if an error occurs.
func (drq *DeployRecordQuery) AllX(ctx context.Context) []*DeployRecord {
	nodes, err := drq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeployRecord IDs.
func (drq *DeployRecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, drq.ctx, "IDs")
	if err := drq.Select(deployrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drq *DeployRecordQuery) IDsX(ctx context.Context) []int {
	ids, err := drq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drq *DeployRecordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, drq.ctx, "Count")
	if err := drq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, drq, querierCount[*DeployRecordQuery](), drq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (drq *DeployRecordQuery) CountX(ctx context.Context) int {
	count, err := drq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drq *DeployRecordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, drq.ctx, "Exist")
	switch _, err := drq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (drq *DeployRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := drq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeployRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drq *DeployRecordQuery) Clone() *DeployRecordQuery {
	if drq == nil {
		return nil
	}
	return &DeployRecordQuery{
		config:            drq.config,
		ctx:               drq.ctx.Clone(),
		order:             append([]OrderFunc{}, drq.order...),
		inters:            append([]Interceptor{}, drq.inters...),
		predicates:        append([]predicate.DeployRecord{}, drq.predicates...),
		withOperator:      drq.withOperator.Clone(),
		withCompileRecord: drq.withCompileRecord.Clone(),
		// clone intermediate query.
		sql:  drq.sql.Clone(),
		path: drq.path,
	}
}

// WithOperator tells the query-builder to eager-load the nodes that are connected to
// the "operator" edge. The optional arguments are used to configure the query builder of the edge.
func (drq *DeployRecordQuery) WithOperator(opts ...func(*UserQuery)) *DeployRecordQuery {
	query := (&UserClient{config: drq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	drq.withOperator = query
	return drq
}

// WithCompileRecord tells the query-builder to eager-load the nodes that are connected to
// the "compile_record" edge. The optional arguments are used to configure the query builder of the edge.
func (drq *DeployRecordQuery) WithCompileRecord(opts ...func(*CompileRecordQuery)) *DeployRecordQuery {
	query := (&CompileRecordClient{config: drq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	drq.withCompileRecord = query
	return drq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Image string `json:"image,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeployRecord.Query().
//		GroupBy(deployrecord.FieldImage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (drq *DeployRecordQuery) GroupBy(field string, fields ...string) *DeployRecordGroupBy {
	drq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeployRecordGroupBy{build: drq}
	grbuild.flds = &drq.ctx.Fields
	grbuild.label = deployrecord.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Image string `json:"image,omitempty"`
//	}
//
//	client.DeployRecord.Query().
//		Select(deployrecord.FieldImage).
//		Scan(ctx, &v)
func (drq *DeployRecordQuery) Select(fields ...string) *DeployRecordSelect {
	drq.ctx.Fields = append(drq.ctx.Fields, fields...)
	sbuild := &DeployRecordSelect{DeployRecordQuery: drq}
	sbuild.label = deployrecord.Label
	sbuild.flds, sbuild.scan = &drq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeployRecordSelect configured with the given aggregations.
func (drq *DeployRecordQuery) Aggregate(fns ...AggregateFunc) *DeployRecordSelect {
	return drq.Select().Aggregate(fns...)
}

func (drq *DeployRecordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range drq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, drq); err != nil {
				return err
			}
		}
	}
	for _, f := range drq.ctx.Fields {
		if !deployrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if drq.path != nil {
		prev, err := drq.path(ctx)
		if err != nil {
			return err
		}
		drq.sql = prev
	}
	return nil
}

func (drq *DeployRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeployRecord, error) {
	var (
		nodes       = []*DeployRecord{}
		_spec       = drq.querySpec()
		loadedTypes = [2]bool{
			drq.withOperator != nil,
			drq.withCompileRecord != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DeployRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DeployRecord{config: drq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, drq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := drq.withOperator; query != nil {
		if err := drq.loadOperator(ctx, query, nodes,
			func(n *DeployRecord) { n.Edges.Operator = []*User{} },
			func(n *DeployRecord, e *User) { n.Edges.Operator = append(n.Edges.Operator, e) }); err != nil {
			return nil, err
		}
	}
	if query := drq.withCompileRecord; query != nil {
		if err := drq.loadCompileRecord(ctx, query, nodes,
			func(n *DeployRecord) { n.Edges.CompileRecord = []*CompileRecord{} },
			func(n *DeployRecord, e *CompileRecord) { n.Edges.CompileRecord = append(n.Edges.CompileRecord, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (drq *DeployRecordQuery) loadOperator(ctx context.Context, query *UserQuery, nodes []*DeployRecord, init func(*DeployRecord), assign func(*DeployRecord, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DeployRecord)
	nids := make(map[int]map[*DeployRecord]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(deployrecord.OperatorTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(deployrecord.OperatorPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(deployrecord.OperatorPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(deployrecord.OperatorPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*DeployRecord]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "operator" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (drq *DeployRecordQuery) loadCompileRecord(ctx context.Context, query *CompileRecordQuery, nodes []*DeployRecord, init func(*DeployRecord), assign func(*DeployRecord, *CompileRecord)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*DeployRecord)
	nids := make(map[int]map[*DeployRecord]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(deployrecord.CompileRecordTable)
		s.Join(joinT).On(s.C(compilerecord.FieldID), joinT.C(deployrecord.CompileRecordPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(deployrecord.CompileRecordPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(deployrecord.CompileRecordPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*DeployRecord]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "compile_record" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (drq *DeployRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drq.querySpec()
	_spec.Node.Columns = drq.ctx.Fields
	if len(drq.ctx.Fields) > 0 {
		_spec.Unique = drq.ctx.Unique != nil && *drq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, drq.driver, _spec)
}

func (drq *DeployRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deployrecord.Table,
			Columns: deployrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: deployrecord.FieldID,
			},
		},
		From:   drq.sql,
		Unique: true,
	}
	if unique := drq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := drq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deployrecord.FieldID)
		for i := range fields {
			if fields[i] != deployrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := drq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (drq *DeployRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(drq.driver.Dialect())
	t1 := builder.Table(deployrecord.Table)
	columns := drq.ctx.Fields
	if len(columns) == 0 {
		columns = deployrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if drq.sql != nil {
		selector = drq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if drq.ctx.Unique != nil && *drq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range drq.predicates {
		p(selector)
	}
	for _, p := range drq.order {
		p(selector)
	}
	if offset := drq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeployRecordGroupBy is the group-by builder for DeployRecord entities.
type DeployRecordGroupBy struct {
	selector
	build *DeployRecordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drgb *DeployRecordGroupBy) Aggregate(fns ...AggregateFunc) *DeployRecordGroupBy {
	drgb.fns = append(drgb.fns, fns...)
	return drgb
}

// Scan applies the selector query and scans the result into the given value.
func (drgb *DeployRecordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, drgb.build.ctx, "GroupBy")
	if err := drgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeployRecordQuery, *DeployRecordGroupBy](ctx, drgb.build, drgb, drgb.build.inters, v)
}

func (drgb *DeployRecordGroupBy) sqlScan(ctx context.Context, root *DeployRecordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(drgb.fns))
	for _, fn := range drgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*drgb.flds)+len(drgb.fns))
		for _, f := range *drgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*drgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeployRecordSelect is the builder for selecting fields of DeployRecord entities.
type DeployRecordSelect struct {
	*DeployRecordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (drs *DeployRecordSelect) Aggregate(fns ...AggregateFunc) *DeployRecordSelect {
	drs.fns = append(drs.fns, fns...)
	return drs
}

// Scan applies the selector query and scans the result into the given value.
func (drs *DeployRecordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, drs.ctx, "Select")
	if err := drs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeployRecordQuery, *DeployRecordSelect](ctx, drs.DeployRecordQuery, drs, drs.inters, v)
}

func (drs *DeployRecordSelect) sqlScan(ctx context.Context, root *DeployRecordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(drs.fns))
	for _, fn := range drs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*drs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}