package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const driverName = `mysql`

var _db *sql.DB // global database connection

// Init database connection
func Init(c *mysql.Config) (err error) {
	if c == nil {
		c = &mysql.Config{
			User:      "root",
			Net:       "tcp",
			Addr:      "127.0.0.1:3306",
			DBName:    "apidb",
			ParseTime: true,
		}
	}

	dsn := c.FormatDSN()

	_db, err = sql.Open(driverName, dsn)
	if err != nil {
		return err
	}
	return nil
}

// BeginTx returns a transaction
func BeginTx() (*sql.Tx, error) {
	if _db == nil {
		return nil, errors.New(`database connection has not been initialized`)
	}
	tx, err := _db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, `begining transaction`)
	}
	return tx, nil
}
