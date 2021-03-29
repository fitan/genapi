// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"ent_samp/ent/car"
	"ent_samp/ent/predicate"
	"ent_samp/ent/user"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCar  = "Car"
	TypeUser = "User"
)

// CarMutation represents an operation that mutates the Car nodes in the graph.
type CarMutation struct {
	config
	op            Op
	typ           string
	id            *int
	model         *string
	registered_at *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Car, error)
	predicates    []predicate.Car
}

var _ ent.Mutation = (*CarMutation)(nil)

// carOption allows management of the mutation configuration using functional options.
type carOption func(*CarMutation)

// newCarMutation creates new mutation for the Car entity.
func newCarMutation(c config, op Op, opts ...carOption) *CarMutation {
	m := &CarMutation{
		config:        c,
		op:            op,
		typ:           TypeCar,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCarID sets the ID field of the mutation.
func withCarID(id int) carOption {
	return func(m *CarMutation) {
		var (
			err   error
			once  sync.Once
			value *Car
		)
		m.oldValue = func(ctx context.Context) (*Car, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Car.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCar sets the old Car of the mutation.
func withCar(node *Car) carOption {
	return func(m *CarMutation) {
		m.oldValue = func(context.Context) (*Car, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CarMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CarMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *CarMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetModel sets the "model" field.
func (m *CarMutation) SetModel(s string) {
	m.model = &s
}

// Model returns the value of the "model" field in the mutation.
func (m *CarMutation) Model() (r string, exists bool) {
	v := m.model
	if v == nil {
		return
	}
	return *v, true
}

// OldModel returns the old "model" field's value of the Car entity.
// If the Car object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CarMutation) OldModel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldModel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldModel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModel: %w", err)
	}
	return oldValue.Model, nil
}

// ResetModel resets all changes to the "model" field.
func (m *CarMutation) ResetModel() {
	m.model = nil
}

// SetRegisteredAt sets the "registered_at" field.
func (m *CarMutation) SetRegisteredAt(t time.Time) {
	m.registered_at = &t
}

// RegisteredAt returns the value of the "registered_at" field in the mutation.
func (m *CarMutation) RegisteredAt() (r time.Time, exists bool) {
	v := m.registered_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRegisteredAt returns the old "registered_at" field's value of the Car entity.
// If the Car object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CarMutation) OldRegisteredAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRegisteredAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRegisteredAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRegisteredAt: %w", err)
	}
	return oldValue.RegisteredAt, nil
}

// ResetRegisteredAt resets all changes to the "registered_at" field.
func (m *CarMutation) ResetRegisteredAt() {
	m.registered_at = nil
}

// Op returns the operation name.
func (m *CarMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.model != nil {
		fields = append(fields, car.FieldModel)
	}
	if m.registered_at != nil {
		fields = append(fields, car.FieldRegisteredAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case car.FieldModel:
		return m.Model()
	case car.FieldRegisteredAt:
		return m.RegisteredAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CarMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case car.FieldModel:
		return m.OldModel(ctx)
	case car.FieldRegisteredAt:
		return m.OldRegisteredAt(ctx)
	}
	return nil, fmt.Errorf("unknown Car field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	case car.FieldModel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModel(v)
		return nil
	case car.FieldRegisteredAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRegisteredAt(v)
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	switch name {
	case car.FieldModel:
		m.ResetModel()
		return nil
	case car.FieldRegisteredAt:
		m.ResetRegisteredAt()
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Car edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	age1          *int
	addage1       *int
	en            *user.En
	clearedFields map[string]struct{}
	cars          map[int]struct{}
	removedcars   map[int]struct{}
	clearedcars   bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetAge1 sets the "age1" field.
func (m *UserMutation) SetAge1(i int) {
	m.age1 = &i
	m.addage1 = nil
}

// Age1 returns the value of the "age1" field in the mutation.
func (m *UserMutation) Age1() (r int, exists bool) {
	v := m.age1
	if v == nil {
		return
	}
	return *v, true
}

// OldAge1 returns the old "age1" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge1(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge1 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAge1 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge1: %w", err)
	}
	return oldValue.Age1, nil
}

// AddAge1 adds i to the "age1" field.
func (m *UserMutation) AddAge1(i int) {
	if m.addage1 != nil {
		*m.addage1 += i
	} else {
		m.addage1 = &i
	}
}

// AddedAge1 returns the value that was added to the "age1" field in this mutation.
func (m *UserMutation) AddedAge1() (r int, exists bool) {
	v := m.addage1
	if v == nil {
		return
	}
	return *v, true
}

// ClearAge1 clears the value of the "age1" field.
func (m *UserMutation) ClearAge1() {
	m.age1 = nil
	m.addage1 = nil
	m.clearedFields[user.FieldAge1] = struct{}{}
}

// Age1Cleared returns if the "age1" field was cleared in this mutation.
func (m *UserMutation) Age1Cleared() bool {
	_, ok := m.clearedFields[user.FieldAge1]
	return ok
}

// ResetAge1 resets all changes to the "age1" field.
func (m *UserMutation) ResetAge1() {
	m.age1 = nil
	m.addage1 = nil
	delete(m.clearedFields, user.FieldAge1)
}

// SetEn sets the "en" field.
func (m *UserMutation) SetEn(u user.En) {
	m.en = &u
}

// En returns the value of the "en" field in the mutation.
func (m *UserMutation) En() (r user.En, exists bool) {
	v := m.en
	if v == nil {
		return
	}
	return *v, true
}

// OldEn returns the old "en" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEn(ctx context.Context) (v user.En, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEn is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEn requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEn: %w", err)
	}
	return oldValue.En, nil
}

// ClearEn clears the value of the "en" field.
func (m *UserMutation) ClearEn() {
	m.en = nil
	m.clearedFields[user.FieldEn] = struct{}{}
}

// EnCleared returns if the "en" field was cleared in this mutation.
func (m *UserMutation) EnCleared() bool {
	_, ok := m.clearedFields[user.FieldEn]
	return ok
}

// ResetEn resets all changes to the "en" field.
func (m *UserMutation) ResetEn() {
	m.en = nil
	delete(m.clearedFields, user.FieldEn)
}

// AddCarIDs adds the "cars" edge to the Car entity by ids.
func (m *UserMutation) AddCarIDs(ids ...int) {
	if m.cars == nil {
		m.cars = make(map[int]struct{})
	}
	for i := range ids {
		m.cars[ids[i]] = struct{}{}
	}
}

// ClearCars clears the "cars" edge to the Car entity.
func (m *UserMutation) ClearCars() {
	m.clearedcars = true
}

// CarsCleared returns if the "cars" edge to the Car entity was cleared.
func (m *UserMutation) CarsCleared() bool {
	return m.clearedcars
}

// RemoveCarIDs removes the "cars" edge to the Car entity by IDs.
func (m *UserMutation) RemoveCarIDs(ids ...int) {
	if m.removedcars == nil {
		m.removedcars = make(map[int]struct{})
	}
	for i := range ids {
		m.removedcars[ids[i]] = struct{}{}
	}
}

// RemovedCars returns the removed IDs of the "cars" edge to the Car entity.
func (m *UserMutation) RemovedCarsIDs() (ids []int) {
	for id := range m.removedcars {
		ids = append(ids, id)
	}
	return
}

// CarsIDs returns the "cars" edge IDs in the mutation.
func (m *UserMutation) CarsIDs() (ids []int) {
	for id := range m.cars {
		ids = append(ids, id)
	}
	return
}

// ResetCars resets all changes to the "cars" edge.
func (m *UserMutation) ResetCars() {
	m.cars = nil
	m.clearedcars = false
	m.removedcars = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.age1 != nil {
		fields = append(fields, user.FieldAge1)
	}
	if m.en != nil {
		fields = append(fields, user.FieldEn)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldAge1:
		return m.Age1()
	case user.FieldEn:
		return m.En()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldAge1:
		return m.OldAge1(ctx)
	case user.FieldEn:
		return m.OldEn(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldAge1:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge1(v)
		return nil
	case user.FieldEn:
		v, ok := value.(user.En)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEn(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage1 != nil {
		fields = append(fields, user.FieldAge1)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge1:
		return m.AddedAge1()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge1:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge1(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldAge1) {
		fields = append(fields, user.FieldAge1)
	}
	if m.FieldCleared(user.FieldEn) {
		fields = append(fields, user.FieldEn)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldAge1:
		m.ClearAge1()
		return nil
	case user.FieldEn:
		m.ClearEn()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldAge1:
		m.ResetAge1()
		return nil
	case user.FieldEn:
		m.ResetEn()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cars != nil {
		edges = append(edges, user.EdgeCars)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCars:
		ids := make([]ent.Value, 0, len(m.cars))
		for id := range m.cars {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcars != nil {
		edges = append(edges, user.EdgeCars)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCars:
		ids := make([]ent.Value, 0, len(m.removedcars))
		for id := range m.removedcars {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcars {
		edges = append(edges, user.EdgeCars)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeCars:
		return m.clearedcars
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCars:
		m.ResetCars()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
