package sentrysql

import (
	"context"
	"database/sql/driver"
)

var (
	_ driver.Driver        = (*sentryDriver)(nil)
	_ driver.DriverContext = (*sentryDriver)(nil)
)

type sentryDriver struct {
	drv driver.Driver
	cfg *config
}

// newDriver returns a driver wrapper.
func newDriver(drv driver.Driver, cfg *config) driver.Driver {
	_ = "STUB: not implemented"
	return *new(driver.Driver)
}

// implement only driver.Driver

// Open implements driver.Driver.
func (d *sentryDriver) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// OpenConnector implements driver.DriverContext.
func (d *sentryDriver) OpenConnector(name string) (driver.Connector, error) {
	_ = "STUB: not implemented"
	return *new(driver.Connector), nil
}

// dsnConnector is a connector for drivers that do not implement
// driver.DriverContext.
type dsnConnector struct {
	dsn string
	drv driver.Driver
}

func (c *dsnConnector) Connect(_ context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (c *dsnConnector) Driver() driver.Driver {
	_ = "STUB: not implemented"
	return *new(driver.Driver)
}
