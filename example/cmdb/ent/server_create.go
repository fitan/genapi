// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cmdb/ent/server"
	"cmdb/ent/servicetree"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServerCreate is the builder for creating a Server entity.
type ServerCreate struct {
	config
	mutation *ServerMutation
	hooks    []Hook
}

// SetIP sets the "ip" field.
func (sc *ServerCreate) SetIP(s string) *ServerCreate {
	sc.mutation.SetIP(s)
	return sc
}

// SetMachineType sets the "machine_type" field.
func (sc *ServerCreate) SetMachineType(st server.MachineType) *ServerCreate {
	sc.mutation.SetMachineType(st)
	return sc
}

// SetPlatformType sets the "platform_type" field.
func (sc *ServerCreate) SetPlatformType(st server.PlatformType) *ServerCreate {
	sc.mutation.SetPlatformType(st)
	return sc
}

// SetSystemType sets the "system_type" field.
func (sc *ServerCreate) SetSystemType(st server.SystemType) *ServerCreate {
	sc.mutation.SetSystemType(st)
	return sc
}

// SetOwnerID sets the "owner" edge to the ServiceTree entity by ID.
func (sc *ServerCreate) SetOwnerID(id int) *ServerCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetNillableOwnerID sets the "owner" edge to the ServiceTree entity by ID if the given value is not nil.
func (sc *ServerCreate) SetNillableOwnerID(id *int) *ServerCreate {
	if id != nil {
		sc = sc.SetOwnerID(*id)
	}
	return sc
}

// SetOwner sets the "owner" edge to the ServiceTree entity.
func (sc *ServerCreate) SetOwner(s *ServiceTree) *ServerCreate {
	return sc.SetOwnerID(s.ID)
}

// Mutation returns the ServerMutation object of the builder.
func (sc *ServerCreate) Mutation() *ServerMutation {
	return sc.mutation
}

// Save creates the Server in the database.
func (sc *ServerCreate) Save(ctx context.Context) (*Server, error) {
	var (
		err  error
		node *Server
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServerCreate) SaveX(ctx context.Context) *Server {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServerCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServerCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServerCreate) check() error {
	if _, ok := sc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "ip"`)}
	}
	if _, ok := sc.mutation.MachineType(); !ok {
		return &ValidationError{Name: "machine_type", err: errors.New(`ent: missing required field "machine_type"`)}
	}
	if v, ok := sc.mutation.MachineType(); ok {
		if err := server.MachineTypeValidator(v); err != nil {
			return &ValidationError{Name: "machine_type", err: fmt.Errorf(`ent: validator failed for field "machine_type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.PlatformType(); !ok {
		return &ValidationError{Name: "platform_type", err: errors.New(`ent: missing required field "platform_type"`)}
	}
	if v, ok := sc.mutation.PlatformType(); ok {
		if err := server.PlatformTypeValidator(v); err != nil {
			return &ValidationError{Name: "platform_type", err: fmt.Errorf(`ent: validator failed for field "platform_type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.SystemType(); !ok {
		return &ValidationError{Name: "system_type", err: errors.New(`ent: missing required field "system_type"`)}
	}
	if v, ok := sc.mutation.SystemType(); ok {
		if err := server.SystemTypeValidator(v); err != nil {
			return &ValidationError{Name: "system_type", err: fmt.Errorf(`ent: validator failed for field "system_type": %w`, err)}
		}
	}
	return nil
}

func (sc *ServerCreate) sqlSave(ctx context.Context) (*Server, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ServerCreate) createSpec() (*Server, *sqlgraph.CreateSpec) {
	var (
		_node = &Server{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: server.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: server.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: server.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := sc.mutation.MachineType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: server.FieldMachineType,
		})
		_node.MachineType = value
	}
	if value, ok := sc.mutation.PlatformType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: server.FieldPlatformType,
		})
		_node.PlatformType = value
	}
	if value, ok := sc.mutation.SystemType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: server.FieldSystemType,
		})
		_node.SystemType = value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   server.OwnerTable,
			Columns: []string{server.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicetree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.service_tree_servers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServerCreateBulk is the builder for creating many Server entities in bulk.
type ServerCreateBulk struct {
	config
	builders []*ServerCreate
}

// Save creates the Server entities in the database.
func (scb *ServerCreateBulk) Save(ctx context.Context) ([]*Server, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Server, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServerCreateBulk) SaveX(ctx context.Context) []*Server {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServerCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServerCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
