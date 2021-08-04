// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cmdb/ent/predicate"
	"cmdb/ent/server"
	"cmdb/ent/servicetree"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServiceTreeQuery is the builder for querying ServiceTree entities.
type ServiceTreeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.ServiceTree
	// eager-loading edges.
	withProject *ServiceTreeQuery
	withService *ServiceTreeQuery
	withServers *ServerQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServiceTreeQuery builder.
func (stq *ServiceTreeQuery) Where(ps ...predicate.ServiceTree) *ServiceTreeQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *ServiceTreeQuery) Limit(limit int) *ServiceTreeQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *ServiceTreeQuery) Offset(offset int) *ServiceTreeQuery {
	stq.offset = &offset
	return stq
}

// Order adds an order step to the query.
func (stq *ServiceTreeQuery) Order(o ...OrderFunc) *ServiceTreeQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// QueryProject chains the current query on the "project" edge.
func (stq *ServiceTreeQuery) QueryProject() *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servicetree.Table, servicetree.FieldID, selector),
			sqlgraph.To(servicetree.Table, servicetree.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, servicetree.ProjectTable, servicetree.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryService chains the current query on the "service" edge.
func (stq *ServiceTreeQuery) QueryService() *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servicetree.Table, servicetree.FieldID, selector),
			sqlgraph.To(servicetree.Table, servicetree.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, servicetree.ServiceTable, servicetree.ServiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryServers chains the current query on the "servers" edge.
func (stq *ServiceTreeQuery) QueryServers() *ServerQuery {
	query := &ServerQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(servicetree.Table, servicetree.FieldID, selector),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, servicetree.ServersTable, servicetree.ServersColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServiceTree entity from the query.
// Returns a *NotFoundError when no ServiceTree was found.
func (stq *ServiceTreeQuery) First(ctx context.Context) (*ServiceTree, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{servicetree.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *ServiceTreeQuery) FirstX(ctx context.Context) *ServiceTree {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServiceTree ID from the query.
// Returns a *NotFoundError when no ServiceTree ID was found.
func (stq *ServiceTreeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{servicetree.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *ServiceTreeQuery) FirstIDX(ctx context.Context) int {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServiceTree entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ServiceTree entity is not found.
// Returns a *NotFoundError when no ServiceTree entities are found.
func (stq *ServiceTreeQuery) Only(ctx context.Context) (*ServiceTree, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{servicetree.Label}
	default:
		return nil, &NotSingularError{servicetree.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *ServiceTreeQuery) OnlyX(ctx context.Context) *ServiceTree {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServiceTree ID in the query.
// Returns a *NotSingularError when exactly one ServiceTree ID is not found.
// Returns a *NotFoundError when no entities are found.
func (stq *ServiceTreeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = &NotSingularError{servicetree.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *ServiceTreeQuery) OnlyIDX(ctx context.Context) int {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServiceTrees.
func (stq *ServiceTreeQuery) All(ctx context.Context) ([]*ServiceTree, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *ServiceTreeQuery) AllX(ctx context.Context) []*ServiceTree {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServiceTree IDs.
func (stq *ServiceTreeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := stq.Select(servicetree.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *ServiceTreeQuery) IDsX(ctx context.Context) []int {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *ServiceTreeQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *ServiceTreeQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *ServiceTreeQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *ServiceTreeQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServiceTreeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *ServiceTreeQuery) Clone() *ServiceTreeQuery {
	if stq == nil {
		return nil
	}
	return &ServiceTreeQuery{
		config:      stq.config,
		limit:       stq.limit,
		offset:      stq.offset,
		order:       append([]OrderFunc{}, stq.order...),
		predicates:  append([]predicate.ServiceTree{}, stq.predicates...),
		withProject: stq.withProject.Clone(),
		withService: stq.withService.Clone(),
		withServers: stq.withServers.Clone(),
		// clone intermediate query.
		sql:  stq.sql.Clone(),
		path: stq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ServiceTreeQuery) WithProject(opts ...func(*ServiceTreeQuery)) *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withProject = query
	return stq
}

// WithService tells the query-builder to eager-load the nodes that are connected to
// the "service" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ServiceTreeQuery) WithService(opts ...func(*ServiceTreeQuery)) *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withService = query
	return stq
}

// WithServers tells the query-builder to eager-load the nodes that are connected to
// the "servers" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ServiceTreeQuery) WithServers(opts ...func(*ServerQuery)) *ServiceTreeQuery {
	query := &ServerQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withServers = query
	return stq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServiceTree.Query().
//		GroupBy(servicetree.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (stq *ServiceTreeQuery) GroupBy(field string, fields ...string) *ServiceTreeGroupBy {
	group := &ServiceTreeGroupBy{config: stq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ServiceTree.Query().
//		Select(servicetree.FieldName).
//		Scan(ctx, &v)
//
func (stq *ServiceTreeQuery) Select(field string, fields ...string) *ServiceTreeSelect {
	stq.fields = append([]string{field}, fields...)
	return &ServiceTreeSelect{ServiceTreeQuery: stq}
}

func (stq *ServiceTreeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !servicetree.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *ServiceTreeQuery) sqlAll(ctx context.Context) ([]*ServiceTree, error) {
	var (
		nodes       = []*ServiceTree{}
		withFKs     = stq.withFKs
		_spec       = stq.querySpec()
		loadedTypes = [3]bool{
			stq.withProject != nil,
			stq.withService != nil,
			stq.withServers != nil,
		}
	)
	if stq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, servicetree.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ServiceTree{config: stq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := stq.withProject; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ServiceTree)
		for i := range nodes {
			fk := nodes[i].service_tree_service
			if fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(servicetree.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "service_tree_service" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Project = n
			}
		}
	}

	if query := stq.withService; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ServiceTree)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Service = []*ServiceTree{}
		}
		query.withFKs = true
		query.Where(predicate.ServiceTree(func(s *sql.Selector) {
			s.Where(sql.InValues(servicetree.ServiceColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.service_tree_service
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "service_tree_service" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "service_tree_service" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Service = append(node.Edges.Service, n)
		}
	}

	if query := stq.withServers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ServiceTree)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Servers = []*Server{}
		}
		query.withFKs = true
		query.Where(predicate.Server(func(s *sql.Selector) {
			s.Where(sql.InValues(servicetree.ServersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.service_tree_servers
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "service_tree_servers" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "service_tree_servers" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Servers = append(node.Edges.Servers, n)
		}
	}

	return nodes, nil
}

func (stq *ServiceTreeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *ServiceTreeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := stq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (stq *ServiceTreeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicetree.Table,
			Columns: servicetree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicetree.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicetree.FieldID)
		for i := range fields {
			if fields[i] != servicetree.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, servicetree.ValidColumn)
			}
		}
	}
	return _spec
}

func (stq *ServiceTreeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(servicetree.Table)
	selector := builder.Select(t1.Columns(servicetree.Columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(servicetree.Columns...)...)
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector, servicetree.ValidColumn)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceTreeGroupBy is the group-by builder for ServiceTree entities.
type ServiceTreeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *ServiceTreeGroupBy) Aggregate(fns ...AggregateFunc) *ServiceTreeGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *ServiceTreeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := stgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) StringsX(ctx context.Context) []string {
	v, err := stgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = stgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) StringX(ctx context.Context) string {
	v, err := stgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) IntsX(ctx context.Context) []int {
	v, err := stgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = stgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) IntX(ctx context.Context) int {
	v, err := stgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := stgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = stgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := stgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := stgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ServiceTreeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = stgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (stgb *ServiceTreeGroupBy) BoolX(ctx context.Context) bool {
	v, err := stgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (stgb *ServiceTreeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range stgb.fields {
		if !servicetree.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *ServiceTreeGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql
	columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
	columns = append(columns, stgb.fields...)
	for _, fn := range stgb.fns {
		columns = append(columns, fn(selector, servicetree.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(stgb.fields...)
}

// ServiceTreeSelect is the builder for selecting fields of ServiceTree entities.
type ServiceTreeSelect struct {
	*ServiceTreeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *ServiceTreeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.ServiceTreeQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sts *ServiceTreeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sts *ServiceTreeSelect) StringsX(ctx context.Context) []string {
	v, err := sts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sts *ServiceTreeSelect) StringX(ctx context.Context) string {
	v, err := sts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sts *ServiceTreeSelect) IntsX(ctx context.Context) []int {
	v, err := sts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sts *ServiceTreeSelect) IntX(ctx context.Context) int {
	v, err := sts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sts *ServiceTreeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sts *ServiceTreeSelect) Float64X(ctx context.Context) float64 {
	v, err := sts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ServiceTreeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sts *ServiceTreeSelect) BoolsX(ctx context.Context) []bool {
	v, err := sts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sts *ServiceTreeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{servicetree.Label}
	default:
		err = fmt.Errorf("ent: ServiceTreeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sts *ServiceTreeSelect) BoolX(ctx context.Context) bool {
	v, err := sts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sts *ServiceTreeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sts.sqlQuery().Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sts *ServiceTreeSelect) sqlQuery() sql.Querier {
	selector := sts.sql
	selector.Select(selector.Columns(sts.fields...)...)
	return selector
}
