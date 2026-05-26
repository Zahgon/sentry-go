package sentrysql

import (
	"context"
	"database/sql/driver"
)

// sentryStmt wraps a driver.Stmt.
type sentryStmt struct {
	stmt  driver.Stmt
	conn  *sentryConn
	cfg   *config
	query string
}

func newStmt(s driver.Stmt, conn *sentryConn, cfg *config, query string) driver.Stmt {
	_ = "STUB: not implemented"
	return *new(driver.Stmt)
}

// Close implements driver.Stmt.
func (s *sentryStmt) Close() error { _ = "STUB: not implemented"; return nil }

// NumInput implements driver.Stmt.
func (s *sentryStmt) NumInput() int { _ = "STUB: not implemented"; return 0 }

// Exec implements driver.Stmt.
func (s *sentryStmt) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *
	//nolint:staticcheck // required by driver.Stmt; ExecContext covers the modern path.
	new(driver.Result), nil
}

// Query implements driver.Stmt.
func (s *sentryStmt) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *
	//nolint:staticcheck // required by driver.Stmt; QueryContext covers the modern path.
	new(driver.Rows), nil
}

// ExecContext implements driver.StmtExecContext with fallback to Exec.
func (s *sentryStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

//nolint:staticcheck // legacy driver.Stmt.Exec fallback is intentional.

// QueryContext implements driver.StmtQueryContext with fallback to Query.
func (s *sentryStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

//nolint:staticcheck // legacy driver.Stmt.Query fallback is intentional.

// CheckNamedValue implements driver.NamedValueChecker when the underlying
// statement supports it.
func (s *sentryStmt) CheckNamedValue(nv *driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback to sentryConn.CheckNamedValue
// The `database/sql` package checks whether the stmt or conn implement this method
// and calls the first one. Since our implementation satisfies both, we need to manually
// follow the same fallback logic.

// ColumnConverter implements driver.ColumnConverter when the underlying
// statement supports it.
func (s *sentryStmt) ColumnConverter(idx int) driver.ValueConverter {
	_ = "STUB: not implemented"
	return *new(driver.ValueConverter)
}

//nolint:staticcheck // ColumnConverter is deprecated but still honored by the stdlib.
