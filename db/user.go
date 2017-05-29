package db

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// UserTable name
const UserTable = `users`

// Scan user data from database row
func (u *User) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&u.Name, &u.Password, &u.CreatedOn, &u.ModifiedOn)
}

// Create saves user into database
func (u *User) Create(tx *sql.Tx) error {
	u.CreatedOn = time.Now()
	u.ModifiedOn = time.Now()

	stmt := bytes.Buffer{}
	stmt.WriteString(`INSERT INTO `)
	stmt.WriteString(UserTable)
	stmt.WriteString(` (username, password, created_on, modified_on) VALUES (?, ?, ?, ?)`)
	_, err := tx.Exec(stmt.String(), u.Name, u.Password, u.CreatedOn, u.ModifiedOn)
	if err != nil {
		return errors.Wrap(err, `creating user record`)
	}

	return nil
}
