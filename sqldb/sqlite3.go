package sqldb

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func UseSQLite(dbname string, options SQLite3Options) (*sql.DB, error) {
	if dbname == "" {
		return nil, fmt.Errorf("no database name provided")
	}

	d, err := sql.Open("sqlite3", fmt.Sprintf("%s?%s", dbname, strings.Join(options.Flags, "&")))
	if err != nil {
		return nil, err
	}

	if len(options.Pragmas) >= 1 {
		pragmaStmts := strings.Join(options.Pragmas, "; ")
		d.Exec(pragmaStmts)
	}

	if options.EnableLogging {
		loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
		d = sqldblogger.OpenDriver(dbname, d.Driver(), loggerAdapter)

		if d == nil {
			return d, fmt.Errorf("database logger failed")
		}
	}

	return d, nil
}

type SQLite3Options struct {
	Flags         []string
	Pragmas       []string
	EnableLogging bool
}
