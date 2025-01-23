package sqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func UsePG(dsn string) (*sql.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("no database dsn provided")
	}

	d, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return d, nil
}
