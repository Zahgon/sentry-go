package sentrysql

import "database/sql/driver"

// sentryTx wraps a driver.Tx.
type sentryTx struct {
	tx driver.Tx
}

// Commit implements driver.Tx.
func (t *sentryTx) Commit() error { _ = "STUB: not implemented"; return nil }

// Rollback implements driver.Tx.
func (t *sentryTx) Rollback() error { _ = "STUB: not implemented"; return nil }
