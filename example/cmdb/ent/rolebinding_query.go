// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cmdb/ent/predicate"
	"cmdb/ent/rolebinding"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleBindingQuery is the builder for querying RoleBinding entities.
type RoleBindingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RoleBinding
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleBindingQuery builder.
func (rbq *RoleBindingQuery) Where(ps ...predicate.RoleBinding) *RoleBindingQuery {
	rbq.predicates = append(rbq.predicates, ps...)
	return rbq
}

// Limit adds a limit step to the query.
func (rbq *RoleBindingQuery) Limit(limit int) *RoleBindingQuery {
	rbq.limit = &limit
	return rbq
}

// Offset adds an offset step to the query.
func (rbq *RoleBindingQuery) Offset(offset int) *RoleBindingQuery {
	rbq.offset = &offset
	return rbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rbq *RoleBindingQuery) Unique(unique bool) *RoleBindingQuery {
	rbq.unique = &unique
	return rbq
}

// Order adds an order step to the query.
func (rbq *RoleBindingQuery) Order(o ...OrderFunc) *RoleBindingQuery {
	rbq.order = append(rbq.order, o...)
	return rbq
}

// First returns the first RoleBinding entity from the query.
// Returns a *NotFoundError when no RoleBinding was found.
func (rbq *RoleBindingQuery) First(ctx context.Context) (*RoleBinding, error) {
	nodes, err := rbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rolebinding.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rbq *RoleBindingQuery) FirstX(ctx context.Context) *RoleBinding {
	node, err := rbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoleBinding ID from the query.
// Returns a *NotFoundError when no RoleBinding ID was found.
func (rbq *RoleBindingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rolebinding.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rbq *RoleBindingQuery) FirstIDX(ctx context.Context) int {
	id, err := rbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoleBinding entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RoleBinding entity is not found.
// Returns a *NotFoundError when no RoleBinding entities are found.
func (rbq *RoleBindingQuery) Only(ctx context.Context) (*RoleBinding, error) {
	nodes, err := rbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rolebinding.Label}
	default:
		return nil, &NotSingularError{rolebinding.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rbq *RoleBindingQuery) OnlyX(ctx context.Context) *RoleBinding {
	node, err := rbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoleBinding ID in the query.
// Returns a *NotSingularError when exactly one RoleBinding ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rbq *RoleBindingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = &NotSingularError{rolebinding.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rbq *RoleBindingQuery) OnlyIDX(ctx context.Context) int {
	id, err := rbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoleBindings.
func (rbq *RoleBindingQuery) All(ctx context.Context) ([]*RoleBinding, error) {
	if err := rbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rbq *RoleBindingQuery) AllX(ctx context.Context) []*RoleBinding {
	nodes, err := rbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoleBinding IDs.
func (rbq *RoleBindingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rbq.Select(rolebinding.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rbq *RoleBindingQuery) IDsX(ctx context.Context) []int {
	ids, err := rbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rbq *RoleBindingQuery) Count(ctx context.Context) (int, error) {
	if err := rbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rbq *RoleBindingQuery) CountX(ctx context.Context) int {
	count, err := rbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rbq *RoleBindingQuery) Exist(ctx context.Context) (bool, error) {
	if err := rbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rbq *RoleBindingQuery) ExistX(ctx context.Context) bool {
	exist, err := rbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleBindingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rbq *RoleBindingQuery) Clone() *RoleBindingQuery {
	if rbq == nil {
		return nil
	}
	return &RoleBindingQuery{
		config:     rbq.config,
		limit:      rbq.limit,
		offset:     rbq.offset,
		order:      append([]OrderFunc{}, rbq.order...),
		predicates: append([]predicate.RoleBinding{}, rbq.predicates...),
		// clone intermediate query.
		sql:  rbq.sql.Clone(),
		path: rbq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RoleName string `json:"role_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoleBinding.Query().
//		GroupBy(rolebinding.FieldRoleName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rbq *RoleBindingQuery) GroupBy(field string, fields ...string) *RoleBindingGroupBy {
	group := &RoleBindingGroupBy{config: rbq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rbq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RoleName string `json:"role_name,omitempty"`
//	}
//
//	client.RoleBinding.Query().
//		Select(rolebinding.FieldRoleName).
//		Scan(ctx, &v)
//
func (rbq *RoleBindingQuery) Select(fields ...string) *RoleBindingSelect {
	rbq.fields = append(rbq.fields, fields...)
	return &RoleBindingSelect{RoleBindingQuery: rbq}
}

func (rbq *RoleBindingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rbq.fields {
		if !rolebinding.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rbq.path != nil {
		prev, err := rbq.path(ctx)
		if err != nil {
			return err
		}
		rbq.sql = prev
	}
	return nil
}

func (rbq *RoleBindingQuery) sqlAll(ctx context.Context) ([]*RoleBinding, error) {
	var (
		nodes   = []*RoleBinding{}
		withFKs = rbq.withFKs
		_spec   = rbq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, rolebinding.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RoleBinding{config: rbq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, rbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rbq *RoleBindingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rbq.querySpec()
	return sqlgraph.CountNodes(ctx, rbq.driver, _spec)
}

func (rbq *RoleBindingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rbq *RoleBindingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rolebinding.Table,
			Columns: rolebinding.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rolebinding.FieldID,
			},
		},
		From:   rbq.sql,
		Unique: true,
	}
	if unique := rbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolebinding.FieldID)
		for i := range fields {
			if fields[i] != rolebinding.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rbq *RoleBindingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rbq.driver.Dialect())
	t1 := builder.Table(rolebinding.Table)
	columns := rbq.fields
	if len(columns) == 0 {
		columns = rolebinding.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rbq.sql != nil {
		selector = rbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range rbq.predicates {
		p(selector)
	}
	for _, p := range rbq.order {
		p(selector)
	}
	if offset := rbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoleBindingGroupBy is the group-by builder for RoleBinding entities.
type RoleBindingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rbgb *RoleBindingGroupBy) Aggregate(fns ...AggregateFunc) *RoleBindingGroupBy {
	rbgb.fns = append(rbgb.fns, fns...)
	return rbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rbgb *RoleBindingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rbgb.path(ctx)
	if err != nil {
		return err
	}
	rbgb.sql = query
	return rbgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rbgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rbgb.fields) > 1 {
		return nil, errors.New("ent: RoleBindingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) StringsX(ctx context.Context) []string {
	v, err := rbgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rbgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) StringX(ctx context.Context) string {
	v, err := rbgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rbgb.fields) > 1 {
		return nil, errors.New("ent: RoleBindingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) IntsX(ctx context.Context) []int {
	v, err := rbgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rbgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) IntX(ctx context.Context) int {
	v, err := rbgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rbgb.fields) > 1 {
		return nil, errors.New("ent: RoleBindingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rbgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rbgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rbgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rbgb.fields) > 1 {
		return nil, errors.New("ent: RoleBindingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rbgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rbgb *RoleBindingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rbgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rbgb *RoleBindingGroupBy) BoolX(ctx context.Context) bool {
	v, err := rbgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rbgb *RoleBindingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rbgb.fields {
		if !rolebinding.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rbgb *RoleBindingGroupBy) sqlQuery() *sql.Selector {
	selector := rbgb.sql.Select()
	aggregation := make([]string, 0, len(rbgb.fns))
	for _, fn := range rbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rbgb.fields)+len(rbgb.fns))
		for _, f := range rbgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rbgb.fields...)...)
}

// RoleBindingSelect is the builder for selecting fields of RoleBinding entities.
type RoleBindingSelect struct {
	*RoleBindingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rbs *RoleBindingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rbs.prepareQuery(ctx); err != nil {
		return err
	}
	rbs.sql = rbs.RoleBindingQuery.sqlQuery(ctx)
	return rbs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rbs *RoleBindingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rbs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rbs.fields) > 1 {
		return nil, errors.New("ent: RoleBindingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rbs *RoleBindingSelect) StringsX(ctx context.Context) []string {
	v, err := rbs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rbs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rbs *RoleBindingSelect) StringX(ctx context.Context) string {
	v, err := rbs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rbs.fields) > 1 {
		return nil, errors.New("ent: RoleBindingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rbs *RoleBindingSelect) IntsX(ctx context.Context) []int {
	v, err := rbs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rbs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rbs *RoleBindingSelect) IntX(ctx context.Context) int {
	v, err := rbs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rbs.fields) > 1 {
		return nil, errors.New("ent: RoleBindingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rbs *RoleBindingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rbs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rbs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rbs *RoleBindingSelect) Float64X(ctx context.Context) float64 {
	v, err := rbs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rbs.fields) > 1 {
		return nil, errors.New("ent: RoleBindingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rbs *RoleBindingSelect) BoolsX(ctx context.Context) []bool {
	v, err := rbs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rbs *RoleBindingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rbs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rolebinding.Label}
	default:
		err = fmt.Errorf("ent: RoleBindingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rbs *RoleBindingSelect) BoolX(ctx context.Context) bool {
	v, err := rbs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rbs *RoleBindingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rbs.sql.Query()
	if err := rbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
