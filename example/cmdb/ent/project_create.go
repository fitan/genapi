// Code generated by entc, DO NOT EDIT.

package ent

import (
	"cmdb/ent/project"
	"cmdb/ent/rolebinding"
	"cmdb/ent/service"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *ProjectCreate) SetCreateTime(t time.Time) *ProjectCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreateTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProjectCreate) SetUpdateTime(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdateTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// AddRoleBindingIDs adds the "role_bindings" edge to the RoleBinding entity by IDs.
func (pc *ProjectCreate) AddRoleBindingIDs(ids ...int) *ProjectCreate {
	pc.mutation.AddRoleBindingIDs(ids...)
	return pc
}

// AddRoleBindings adds the "role_bindings" edges to the RoleBinding entity.
func (pc *ProjectCreate) AddRoleBindings(r ...*RoleBinding) *ProjectCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRoleBindingIDs(ids...)
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (pc *ProjectCreate) AddServiceIDs(ids ...int) *ProjectCreate {
	pc.mutation.AddServiceIDs(ids...)
	return pc
}

// AddServices adds the "services" edges to the Service entity.
func (pc *ProjectCreate) AddServices(s ...*Service) *ProjectCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddServiceIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := project.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := project.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: project.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: project.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
		_node.Name = value
	}
	if nodes := pc.mutation.RoleBindingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RoleBindingsTable,
			Columns: []string{project.RoleBindingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rolebinding.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ServicesTable,
			Columns: []string{project.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
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

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
