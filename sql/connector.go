package sentrysql

import (
	"context"
	"database/sql/driver"
)

// sentryConnector wraps a driver.Connector so that returned connections are
// wrapped with sentryConn.
type sentryConnector struct {
	connector driver.Connector
	drv       driver.Driver
	cfg       *config
}

func newConnector(c driver.Connector, cfg *config) *sentryConnector {
	_ = "STUB: not implemented"
	return nil
}

// Connect implements driver.Connector.
func (c *sentryConnector) Connect(ctx context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// Driver implements driver.Connector.
func (c *sentryConnector) Driver() driver.Driver {
	_ = "STUB: not implemented"

	// Close checks if underlying connector implements io.Closer to Close
	// the connection.
	return *new(driver.Driver)
}

func (c *sentryConnector) Close() error { _ = "STUB: not implemented"; return nil }
