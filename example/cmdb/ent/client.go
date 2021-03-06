// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"cmdb/ent/migrate"

	"cmdb/ent/alert"
	"cmdb/ent/message"
	"cmdb/ent/rolebinding"
	"cmdb/ent/server"
	"cmdb/ent/servicetree"
	"cmdb/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Alert is the client for interacting with the Alert builders.
	Alert *AlertClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// RoleBinding is the client for interacting with the RoleBinding builders.
	RoleBinding *RoleBindingClient
	// Server is the client for interacting with the Server builders.
	Server *ServerClient
	// ServiceTree is the client for interacting with the ServiceTree builders.
	ServiceTree *ServiceTreeClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Alert = NewAlertClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.RoleBinding = NewRoleBindingClient(c.config)
	c.Server = NewServerClient(c.config)
	c.ServiceTree = NewServiceTreeClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Alert:       NewAlertClient(cfg),
		Message:     NewMessageClient(cfg),
		RoleBinding: NewRoleBindingClient(cfg),
		Server:      NewServerClient(cfg),
		ServiceTree: NewServiceTreeClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:      cfg,
		Alert:       NewAlertClient(cfg),
		Message:     NewMessageClient(cfg),
		RoleBinding: NewRoleBindingClient(cfg),
		Server:      NewServerClient(cfg),
		ServiceTree: NewServiceTreeClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Alert.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Alert.Use(hooks...)
	c.Message.Use(hooks...)
	c.RoleBinding.Use(hooks...)
	c.Server.Use(hooks...)
	c.ServiceTree.Use(hooks...)
	c.User.Use(hooks...)
}

// AlertClient is a client for the Alert schema.
type AlertClient struct {
	config
}

// NewAlertClient returns a client for the Alert from the given config.
func NewAlertClient(c config) *AlertClient {
	return &AlertClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `alert.Hooks(f(g(h())))`.
func (c *AlertClient) Use(hooks ...Hook) {
	c.hooks.Alert = append(c.hooks.Alert, hooks...)
}

// Create returns a create builder for Alert.
func (c *AlertClient) Create() *AlertCreate {
	mutation := newAlertMutation(c.config, OpCreate)
	return &AlertCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Alert entities.
func (c *AlertClient) CreateBulk(builders ...*AlertCreate) *AlertCreateBulk {
	return &AlertCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Alert.
func (c *AlertClient) Update() *AlertUpdate {
	mutation := newAlertMutation(c.config, OpUpdate)
	return &AlertUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlertClient) UpdateOne(a *Alert) *AlertUpdateOne {
	mutation := newAlertMutation(c.config, OpUpdateOne, withAlert(a))
	return &AlertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlertClient) UpdateOneID(id int) *AlertUpdateOne {
	mutation := newAlertMutation(c.config, OpUpdateOne, withAlertID(id))
	return &AlertUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Alert.
func (c *AlertClient) Delete() *AlertDelete {
	mutation := newAlertMutation(c.config, OpDelete)
	return &AlertDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AlertClient) DeleteOne(a *Alert) *AlertDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AlertClient) DeleteOneID(id int) *AlertDeleteOne {
	builder := c.Delete().Where(alert.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlertDeleteOne{builder}
}

// Query returns a query builder for Alert.
func (c *AlertClient) Query() *AlertQuery {
	return &AlertQuery{
		config: c.config,
	}
}

// Get returns a Alert entity by its id.
func (c *AlertClient) Get(ctx context.Context, id int) (*Alert, error) {
	return c.Query().Where(alert.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlertClient) GetX(ctx context.Context, id int) *Alert {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Alert.
func (c *AlertClient) QueryUser(a *Alert) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(alert.Table, alert.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, alert.UserTable, alert.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AlertClient) Hooks() []Hook {
	return c.hooks.Alert
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Create returns a create builder for Message.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id int) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MessageClient) DeleteOneID(id int) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id int) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id int) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Message.
func (c *MessageClient) QueryUser(m *Message) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, message.UserTable, message.UserColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	return c.hooks.Message
}

// RoleBindingClient is a client for the RoleBinding schema.
type RoleBindingClient struct {
	config
}

// NewRoleBindingClient returns a client for the RoleBinding from the given config.
func NewRoleBindingClient(c config) *RoleBindingClient {
	return &RoleBindingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `rolebinding.Hooks(f(g(h())))`.
func (c *RoleBindingClient) Use(hooks ...Hook) {
	c.hooks.RoleBinding = append(c.hooks.RoleBinding, hooks...)
}

// Create returns a create builder for RoleBinding.
func (c *RoleBindingClient) Create() *RoleBindingCreate {
	mutation := newRoleBindingMutation(c.config, OpCreate)
	return &RoleBindingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RoleBinding entities.
func (c *RoleBindingClient) CreateBulk(builders ...*RoleBindingCreate) *RoleBindingCreateBulk {
	return &RoleBindingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RoleBinding.
func (c *RoleBindingClient) Update() *RoleBindingUpdate {
	mutation := newRoleBindingMutation(c.config, OpUpdate)
	return &RoleBindingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleBindingClient) UpdateOne(rb *RoleBinding) *RoleBindingUpdateOne {
	mutation := newRoleBindingMutation(c.config, OpUpdateOne, withRoleBinding(rb))
	return &RoleBindingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleBindingClient) UpdateOneID(id int) *RoleBindingUpdateOne {
	mutation := newRoleBindingMutation(c.config, OpUpdateOne, withRoleBindingID(id))
	return &RoleBindingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RoleBinding.
func (c *RoleBindingClient) Delete() *RoleBindingDelete {
	mutation := newRoleBindingMutation(c.config, OpDelete)
	return &RoleBindingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleBindingClient) DeleteOne(rb *RoleBinding) *RoleBindingDeleteOne {
	return c.DeleteOneID(rb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleBindingClient) DeleteOneID(id int) *RoleBindingDeleteOne {
	builder := c.Delete().Where(rolebinding.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleBindingDeleteOne{builder}
}

// Query returns a query builder for RoleBinding.
func (c *RoleBindingClient) Query() *RoleBindingQuery {
	return &RoleBindingQuery{
		config: c.config,
	}
}

// Get returns a RoleBinding entity by its id.
func (c *RoleBindingClient) Get(ctx context.Context, id int) (*RoleBinding, error) {
	return c.Query().Where(rolebinding.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleBindingClient) GetX(ctx context.Context, id int) *RoleBinding {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RoleBindingClient) Hooks() []Hook {
	return c.hooks.RoleBinding
}

// ServerClient is a client for the Server schema.
type ServerClient struct {
	config
}

// NewServerClient returns a client for the Server from the given config.
func NewServerClient(c config) *ServerClient {
	return &ServerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `server.Hooks(f(g(h())))`.
func (c *ServerClient) Use(hooks ...Hook) {
	c.hooks.Server = append(c.hooks.Server, hooks...)
}

// Create returns a create builder for Server.
func (c *ServerClient) Create() *ServerCreate {
	mutation := newServerMutation(c.config, OpCreate)
	return &ServerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Server entities.
func (c *ServerClient) CreateBulk(builders ...*ServerCreate) *ServerCreateBulk {
	return &ServerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Server.
func (c *ServerClient) Update() *ServerUpdate {
	mutation := newServerMutation(c.config, OpUpdate)
	return &ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerClient) UpdateOne(s *Server) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServer(s))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerClient) UpdateOneID(id int) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServerID(id))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Server.
func (c *ServerClient) Delete() *ServerDelete {
	mutation := newServerMutation(c.config, OpDelete)
	return &ServerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ServerClient) DeleteOne(s *Server) *ServerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ServerClient) DeleteOneID(id int) *ServerDeleteOne {
	builder := c.Delete().Where(server.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerDeleteOne{builder}
}

// Query returns a query builder for Server.
func (c *ServerClient) Query() *ServerQuery {
	return &ServerQuery{
		config: c.config,
	}
}

// Get returns a Server entity by its id.
func (c *ServerClient) Get(ctx context.Context, id int) (*Server, error) {
	return c.Query().Where(server.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerClient) GetX(ctx context.Context, id int) *Server {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Server.
func (c *ServerClient) QueryOwner(s *Server) *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(servicetree.Table, servicetree.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, server.OwnerTable, server.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerClient) Hooks() []Hook {
	return c.hooks.Server
}

// ServiceTreeClient is a client for the ServiceTree schema.
type ServiceTreeClient struct {
	config
}

// NewServiceTreeClient returns a client for the ServiceTree from the given config.
func NewServiceTreeClient(c config) *ServiceTreeClient {
	return &ServiceTreeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `servicetree.Hooks(f(g(h())))`.
func (c *ServiceTreeClient) Use(hooks ...Hook) {
	c.hooks.ServiceTree = append(c.hooks.ServiceTree, hooks...)
}

// Create returns a create builder for ServiceTree.
func (c *ServiceTreeClient) Create() *ServiceTreeCreate {
	mutation := newServiceTreeMutation(c.config, OpCreate)
	return &ServiceTreeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ServiceTree entities.
func (c *ServiceTreeClient) CreateBulk(builders ...*ServiceTreeCreate) *ServiceTreeCreateBulk {
	return &ServiceTreeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ServiceTree.
func (c *ServiceTreeClient) Update() *ServiceTreeUpdate {
	mutation := newServiceTreeMutation(c.config, OpUpdate)
	return &ServiceTreeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServiceTreeClient) UpdateOne(st *ServiceTree) *ServiceTreeUpdateOne {
	mutation := newServiceTreeMutation(c.config, OpUpdateOne, withServiceTree(st))
	return &ServiceTreeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServiceTreeClient) UpdateOneID(id int) *ServiceTreeUpdateOne {
	mutation := newServiceTreeMutation(c.config, OpUpdateOne, withServiceTreeID(id))
	return &ServiceTreeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ServiceTree.
func (c *ServiceTreeClient) Delete() *ServiceTreeDelete {
	mutation := newServiceTreeMutation(c.config, OpDelete)
	return &ServiceTreeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ServiceTreeClient) DeleteOne(st *ServiceTree) *ServiceTreeDeleteOne {
	return c.DeleteOneID(st.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ServiceTreeClient) DeleteOneID(id int) *ServiceTreeDeleteOne {
	builder := c.Delete().Where(servicetree.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServiceTreeDeleteOne{builder}
}

// Query returns a query builder for ServiceTree.
func (c *ServiceTreeClient) Query() *ServiceTreeQuery {
	return &ServiceTreeQuery{
		config: c.config,
	}
}

// Get returns a ServiceTree entity by its id.
func (c *ServiceTreeClient) Get(ctx context.Context, id int) (*ServiceTree, error) {
	return c.Query().Where(servicetree.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServiceTreeClient) GetX(ctx context.Context, id int) *ServiceTree {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProject queries the project edge of a ServiceTree.
func (c *ServiceTreeClient) QueryProject(st *ServiceTree) *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := st.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(servicetree.Table, servicetree.FieldID, id),
			sqlgraph.To(servicetree.Table, servicetree.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, servicetree.ProjectTable, servicetree.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(st.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryService queries the service edge of a ServiceTree.
func (c *ServiceTreeClient) QueryService(st *ServiceTree) *ServiceTreeQuery {
	query := &ServiceTreeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := st.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(servicetree.Table, servicetree.FieldID, id),
			sqlgraph.To(servicetree.Table, servicetree.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, servicetree.ServiceTable, servicetree.ServiceColumn),
		)
		fromV = sqlgraph.Neighbors(st.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryServers queries the servers edge of a ServiceTree.
func (c *ServiceTreeClient) QueryServers(st *ServiceTree) *ServerQuery {
	query := &ServerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := st.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(servicetree.Table, servicetree.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, servicetree.ServersTable, servicetree.ServersColumn),
		)
		fromV = sqlgraph.Neighbors(st.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServiceTreeClient) Hooks() []Hook {
	return c.hooks.ServiceTree
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoleBind queries the role_bind edge of a User.
func (c *UserClient) QueryRoleBind(u *User) *RoleBindingQuery {
	query := &RoleBindingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(rolebinding.Table, rolebinding.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.RoleBindTable, user.RoleBindColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAlert queries the alert edge of a User.
func (c *UserClient) QueryAlert(u *User) *AlertQuery {
	query := &AlertQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(alert.Table, alert.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.AlertTable, user.AlertPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMsg queries the msg edge of a User.
func (c *UserClient) QueryMsg(u *User) *MessageQuery {
	query := &MessageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.MsgTable, user.MsgColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
