// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"ent_samp/ent/car"
	"ent_samp/ent/predicate"
	"ent_samp/ent/user"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetAge1 sets the "age1" field.
func (uu *UserUpdate) SetAge1(i int) *UserUpdate {
	uu.mutation.ResetAge1()
	uu.mutation.SetAge1(i)
	return uu
}

// SetNillableAge1 sets the "age1" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge1(i *int) *UserUpdate {
	if i != nil {
		uu.SetAge1(*i)
	}
	return uu
}

// AddAge1 adds i to the "age1" field.
func (uu *UserUpdate) AddAge1(i int) *UserUpdate {
	uu.mutation.AddAge1(i)
	return uu
}

// ClearAge1 clears the value of the "age1" field.
func (uu *UserUpdate) ClearAge1() *UserUpdate {
	uu.mutation.ClearAge1()
	return uu
}

// SetEn sets the "en" field.
func (uu *UserUpdate) SetEn(u user.En) *UserUpdate {
	uu.mutation.SetEn(u)
	return uu
}

// SetNillableEn sets the "en" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEn(u *user.En) *UserUpdate {
	if u != nil {
		uu.SetEn(*u)
	}
	return uu
}

// ClearEn clears the value of the "en" field.
func (uu *UserUpdate) ClearEn() *UserUpdate {
	uu.mutation.ClearEn()
	return uu
}

// AddCarIDs adds the "cars" edge to the Car entity by IDs.
func (uu *UserUpdate) AddCarIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCarIDs(ids...)
	return uu
}

// AddCars adds the "cars" edges to the Car entity.
func (uu *UserUpdate) AddCars(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCars clears all "cars" edges to the Car entity.
func (uu *UserUpdate) ClearCars() *UserUpdate {
	uu.mutation.ClearCars()
	return uu
}

// RemoveCarIDs removes the "cars" edge to Car entities by IDs.
func (uu *UserUpdate) RemoveCarIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCarIDs(ids...)
	return uu
}

// RemoveCars removes "cars" edges to Car entities.
func (uu *UserUpdate) RemoveCars(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.En(); ok {
		if err := user.EnValidator(v); err != nil {
			return &ValidationError{Name: "en", err: fmt.Errorf("ent: validator failed for field \"en\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Age1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge1,
		})
	}
	if value, ok := uu.mutation.AddedAge1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge1,
		})
	}
	if uu.mutation.Age1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldAge1,
		})
	}
	if value, ok := uu.mutation.En(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldEn,
		})
	}
	if uu.mutation.EnCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldEn,
		})
	}
	if uu.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarsTable,
			Columns: []string{user.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCarsIDs(); len(nodes) > 0 && !uu.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarsTable,
			Columns: []string{user.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarsTable,
			Columns: []string{user.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetAge1 sets the "age1" field.
func (uuo *UserUpdateOne) SetAge1(i int) *UserUpdateOne {
	uuo.mutation.ResetAge1()
	uuo.mutation.SetAge1(i)
	return uuo
}

// SetNillableAge1 sets the "age1" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge1(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetAge1(*i)
	}
	return uuo
}

// AddAge1 adds i to the "age1" field.
func (uuo *UserUpdateOne) AddAge1(i int) *UserUpdateOne {
	uuo.mutation.AddAge1(i)
	return uuo
}

// ClearAge1 clears the value of the "age1" field.
func (uuo *UserUpdateOne) ClearAge1() *UserUpdateOne {
	uuo.mutation.ClearAge1()
	return uuo
}

// SetEn sets the "en" field.
func (uuo *UserUpdateOne) SetEn(u user.En) *UserUpdateOne {
	uuo.mutation.SetEn(u)
	return uuo
}

// SetNillableEn sets the "en" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEn(u *user.En) *UserUpdateOne {
	if u != nil {
		uuo.SetEn(*u)
	}
	return uuo
}

// ClearEn clears the value of the "en" field.
func (uuo *UserUpdateOne) ClearEn() *UserUpdateOne {
	uuo.mutation.ClearEn()
	return uuo
}

// AddCarIDs adds the "cars" edge to the Car entity by IDs.
func (uuo *UserUpdateOne) AddCarIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCarIDs(ids...)
	return uuo
}

// AddCars adds the "cars" edges to the Car entity.
func (uuo *UserUpdateOne) AddCars(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCars clears all "cars" edges to the Car entity.
func (uuo *UserUpdateOne) ClearCars() *UserUpdateOne {
	uuo.mutation.ClearCars()
	return uuo
}

// RemoveCarIDs removes the "cars" edge to Car entities by IDs.
func (uuo *UserUpdateOne) RemoveCarIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCarIDs(ids...)
	return uuo
}

// RemoveCars removes "cars" edges to Car entities.
func (uuo *UserUpdateOne) RemoveCars(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarIDs(ids...)
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.En(); ok {
		if err := user.EnValidator(v); err != nil {
			return &ValidationError{Name: "en", err: fmt.Errorf("ent: validator failed for field \"en\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Age1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge1,
		})
	}
	if value, ok := uuo.mutation.AddedAge1(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge1,
		})
	}
	if uuo.mutation.Age1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldAge1,
		})
	}
	if value, ok := uuo.mutation.En(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldEn,
		})
	}
	if uuo.mutation.EnCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldEn,
		})
	}
	if uuo.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarsTable,
			Columns: []string{user.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCarsIDs(); len(nodes) > 0 && !uuo.mutation.CarsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarsTable,
			Columns: []string{user.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CarsTable,
			Columns: []string{user.CarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
