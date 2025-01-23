package sqldb

import (
	"context"
	"database/sql"
)

type SqlBackend interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type SQLAdapter string

type DBConnectOptions struct {
	Adapter        SQLAdapter // [postgresql, sqlite3]
	DSN            string     // i.e postgresql://user:pass@host:5432/dbname
	VerboseLogging bool
}
