package sentrysql

import (
	"context"
	"database/sql/driver"
)

// sentryConn wraps a driver.Conn.
type sentryConn struct {
	conn driver.Conn
	cfg  *config
}

func newConn(c driver.Conn, cfg *config) driver.Conn {
	_ = "STUB: not implemented"
	return *new(driver.Conn)
}

// Ping implements driver.Pinger when the underlying connection does.
func (c *sentryConn) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// QueryContext implements driver.QueryerContext with fallback to the legacy
// driver.Queryer path.
func (c *sentryConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

//nolint:staticcheck // legacy driver.Queryer fallback is intentional.

// ExecContext implements driver.ExecerContext with fallback to the legacy
// driver.Execer path.
func (c *sentryConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

//nolint:staticcheck // legacy driver.Execer fallback is intentional.

// PrepareContext implements driver.ConnPrepareContext with fallback to
// Prepare when the underlying connection does not support it.
func (c *sentryConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

// Prepare implements driver.Conn.
func (c *sentryConn) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

// Close implements driver.Conn.
func (c *sentryConn) Close() error { _ = "STUB: not implemented"; return nil }

// Begin implements driver.Conn.
func (c *sentryConn) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *
	//nolint:staticcheck // required by driver.Conn; BeginTx covers the modern path.
	new(driver.Tx), nil
}

// BeginTx implements driver.ConnBeginTx with fallback to Begin.
func (c *sentryConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

// Mirror stdlib ctxDriverBegin: reject non-default TxOptions that can't be
// expressed through the legacy Begin().

// ResetSession implements driver.SessionResetter.
func (c *sentryConn) ResetSession(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// IsValid implements driver.Validator.
func (c *sentryConn) IsValid() bool { _ = "STUB: not implemented"; return false }

// CheckNamedValue implements driver.NamedValueChecker when the underlying
// connection supports it; otherwise it returns driver.ErrSkip so the standard
// library falls back to default value conversion.
func (c *sentryConn) CheckNamedValue(nv *driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

// Raw returns the underlying driver connection. Useful for type-assertions.
func (c *sentryConn) Raw() driver.Conn {
	_ = "STUB: not implemented"

	// namedValuesToValues converts []driver.NamedValue to []driver.Value for
	// fallback calls to the legacy driver.Execer and driver.Queryer interfaces.
	return *new(driver.Conn)
}

func namedValuesToValues(named []driver.NamedValue) ([]driver.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
