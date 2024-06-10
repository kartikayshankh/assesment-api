// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"api/ent/migrate"

	"api/ent/employee"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Employee is the client for interacting with the Employee builders.
	Employee *EmployeeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Employee = NewEmployeeClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Employee: NewEmployeeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:      ctx,
		config:   cfg,
		Employee: NewEmployeeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Employee.
//		Query().
//		Count(ctx)
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
	c.Employee.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Employee.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *EmployeeMutation:
		return c.Employee.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// EmployeeClient is a client for the Employee schema.
type EmployeeClient struct {
	config
}

// NewEmployeeClient returns a client for the Employee from the given config.
func NewEmployeeClient(c config) *EmployeeClient {
	return &EmployeeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employee.Hooks(f(g(h())))`.
func (c *EmployeeClient) Use(hooks ...Hook) {
	c.hooks.Employee = append(c.hooks.Employee, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `employee.Intercept(f(g(h())))`.
func (c *EmployeeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Employee = append(c.inters.Employee, interceptors...)
}

// Create returns a builder for creating a Employee entity.
func (c *EmployeeClient) Create() *EmployeeCreate {
	mutation := newEmployeeMutation(c.config, OpCreate)
	return &EmployeeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Employee entities.
func (c *EmployeeClient) CreateBulk(builders ...*EmployeeCreate) *EmployeeCreateBulk {
	return &EmployeeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EmployeeClient) MapCreateBulk(slice any, setFunc func(*EmployeeCreate, int)) *EmployeeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EmployeeCreateBulk{err: fmt.Errorf("calling to EmployeeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EmployeeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EmployeeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Employee.
func (c *EmployeeClient) Update() *EmployeeUpdate {
	mutation := newEmployeeMutation(c.config, OpUpdate)
	return &EmployeeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmployeeClient) UpdateOne(e *Employee) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployee(e))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmployeeClient) UpdateOneID(id int) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployeeID(id))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Employee.
func (c *EmployeeClient) Delete() *EmployeeDelete {
	mutation := newEmployeeMutation(c.config, OpDelete)
	return &EmployeeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EmployeeClient) DeleteOne(e *Employee) *EmployeeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EmployeeClient) DeleteOneID(id int) *EmployeeDeleteOne {
	builder := c.Delete().Where(employee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmployeeDeleteOne{builder}
}

// Query returns a query builder for Employee.
func (c *EmployeeClient) Query() *EmployeeQuery {
	return &EmployeeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEmployee},
		inters: c.Interceptors(),
	}
}

// Get returns a Employee entity by its id.
func (c *EmployeeClient) Get(ctx context.Context, id int) (*Employee, error) {
	return c.Query().Where(employee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmployeeClient) GetX(ctx context.Context, id int) *Employee {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmployeeClient) Hooks() []Hook {
	return c.hooks.Employee
}

// Interceptors returns the client interceptors.
func (c *EmployeeClient) Interceptors() []Interceptor {
	return c.inters.Employee
}

func (c *EmployeeClient) mutate(ctx context.Context, m *EmployeeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EmployeeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EmployeeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EmployeeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Employee mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Employee []ent.Hook
	}
	inters struct {
		Employee []ent.Interceptor
	}
)