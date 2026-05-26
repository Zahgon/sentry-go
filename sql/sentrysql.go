package sentrysql

import (
	"database/sql"
	"database/sql/driver"
	"sync"
)

// Open is a wrapper over sql.Open that provides Sentry instrumentation.
func Open(driverName, dataSourceName string, opts ...Option) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefer driver.DriverContext for connection pooling semantics.

// OpenDB wraps an existing driver.Connector and returns a *sql.DB.
func OpenDB(c driver.Connector, opts ...Option) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WrapDriver wraps a driver.Driver so connections it hands out are instrumented.
func WrapDriver(drv driver.Driver, opts ...Option) (driver.Driver, error) {
	_ = "STUB: not implemented"
	return *new(driver.Driver), nil
}

// WrapConnector wraps a driver.Connector so connections it hands out are
// instrumented. WithDatabaseSystem is required.
func WrapConnector(c driver.Connector, opts ...Option) (driver.Connector, error) {
	_ = "STUB: not implemented"
	return *new(driver.Connector), nil
}

var (
	registerMu sync.Mutex
	registered = map[string]struct{}{}
)

// Register registers a wrapped version of drv under the name
// "sentrysql-<name>". It is safe to call once per process; subsequent calls
// with the same name return nil without re-registering.
func Register(name string, drv driver.Driver, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// getDriver opens and closes a driver to retrieve the driver implementation we need to wrap.
// We won't use the connection opened here, so we close it to avoid leaks.
func getDriver(name, dataSourceName string) (driver.Driver, error) {
	_ = "STUB: not implemented"
	return *new(driver.Driver), nil
}
