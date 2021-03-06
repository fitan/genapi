// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cmdb/ent/rolebinding"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleBindingCreate is the builder for creating a RoleBinding entity.
type RoleBindingCreate struct {
	config
	mutation *RoleBindingMutation
	hooks    []Hook
}

// SetRoleName sets the "role_name" field.
func (rbc *RoleBindingCreate) SetRoleName(s string) *RoleBindingCreate {
	rbc.mutation.SetRoleName(s)
	return rbc
}

// SetRoleID sets the "role_id" field.
func (rbc *RoleBindingCreate) SetRoleID(s string) *RoleBindingCreate {
	rbc.mutation.SetRoleID(s)
	return rbc
}

// SetStatus sets the "status" field.
func (rbc *RoleBindingCreate) SetStatus(b bool) *RoleBindingCreate {
	rbc.mutation.SetStatus(b)
	return rbc
}

// SetCreatedAt sets the "created_at" field.
func (rbc *RoleBindingCreate) SetCreatedAt(t time.Time) *RoleBindingCreate {
	rbc.mutation.SetCreatedAt(t)
	return rbc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rbc *RoleBindingCreate) SetNillableCreatedAt(t *time.Time) *RoleBindingCreate {
	if t != nil {
		rbc.SetCreatedAt(*t)
	}
	return rbc
}

// SetNote sets the "note" field.
func (rbc *RoleBindingCreate) SetNote(s string) *RoleBindingCreate {
	rbc.mutation.SetNote(s)
	return rbc
}

// SetPermissions sets the "permissions" field.
func (rbc *RoleBindingCreate) SetPermissions(s []string) *RoleBindingCreate {
	rbc.mutation.SetPermissions(s)
	return rbc
}

// Mutation returns the RoleBindingMutation object of the builder.
func (rbc *RoleBindingCreate) Mutation() *RoleBindingMutation {
	return rbc.mutation
}

// Save creates the RoleBinding in the database.
func (rbc *RoleBindingCreate) Save(ctx context.Context) (*RoleBinding, error) {
	var (
		err  error
		node *RoleBinding
	)
	rbc.defaults()
	if len(rbc.hooks) == 0 {
		if err = rbc.check(); err != nil {
			return nil, err
		}
		node, err = rbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleBindingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rbc.check(); err != nil {
				return nil, err
			}
			rbc.mutation = mutation
			if node, err = rbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rbc.hooks) - 1; i >= 0; i-- {
			if rbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rbc *RoleBindingCreate) SaveX(ctx context.Context) *RoleBinding {
	v, err := rbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rbc *RoleBindingCreate) Exec(ctx context.Context) error {
	_, err := rbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbc *RoleBindingCreate) ExecX(ctx context.Context) {
	if err := rbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rbc *RoleBindingCreate) defaults() {
	if _, ok := rbc.mutation.CreatedAt(); !ok {
		v := rolebinding.DefaultCreatedAt()
		rbc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rbc *RoleBindingCreate) check() error {
	if _, ok := rbc.mutation.RoleName(); !ok {
		return &ValidationError{Name: "role_name", err: errors.New(`ent: missing required field "role_name"`)}
	}
	if _, ok := rbc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "role_id"`)}
	}
	if _, ok := rbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if _, ok := rbc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`ent: missing required field "note"`)}
	}
	if _, ok := rbc.mutation.Permissions(); !ok {
		return &ValidationError{Name: "permissions", err: errors.New(`ent: missing required field "permissions"`)}
	}
	return nil
}

func (rbc *RoleBindingCreate) sqlSave(ctx context.Context) (*RoleBinding, error) {
	_node, _spec := rbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rbc *RoleBindingCreate) createSpec() (*RoleBinding, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleBinding{config: rbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rolebinding.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rolebinding.FieldID,
			},
		}
	)
	if value, ok := rbc.mutation.RoleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rolebinding.FieldRoleName,
		})
		_node.RoleName = value
	}
	if value, ok := rbc.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rolebinding.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := rbc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rolebinding.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := rbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rolebinding.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	if value, ok := rbc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rolebinding.FieldNote,
		})
		_node.Note = value
	}
	if value, ok := rbc.mutation.Permissions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: rolebinding.FieldPermissions,
		})
		_node.Permissions = value
	}
	return _node, _spec
}

// RoleBindingCreateBulk is the builder for creating many RoleBinding entities in bulk.
type RoleBindingCreateBulk struct {
	config
	builders []*RoleBindingCreate
}

// Save creates the RoleBinding entities in the database.
func (rbcb *RoleBindingCreateBulk) Save(ctx context.Context) ([]*RoleBinding, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rbcb.builders))
	nodes := make([]*RoleBinding, len(rbcb.builders))
	mutators := make([]Mutator, len(rbcb.builders))
	for i := range rbcb.builders {
		func(i int, root context.Context) {
			builder := rbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleBindingMutation)
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
					_, err = mutators[i+1].Mutate(root, rbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rbcb *RoleBindingCreateBulk) SaveX(ctx context.Context) []*RoleBinding {
	v, err := rbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rbcb *RoleBindingCreateBulk) Exec(ctx context.Context) error {
	_, err := rbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbcb *RoleBindingCreateBulk) ExecX(ctx context.Context) {
	if err := rbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
