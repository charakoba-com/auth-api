package db

import (
	"bytes"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// UserTable name
const UserTable = `users`

// UserSelectColumns is a list of select columns
const UserSelectColumns = `username, password, created_on, modified_on`

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
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, `commiting transaction`)
	}

	return nil
}

// Lookup user by username
func (u *User) Lookup(tx *sql.Tx, username string) error {
	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(UserSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(UserTable)
	stmt.WriteString(` WHERE username=?`)
	row := tx.QueryRow(stmt.String(), username)
	if err := u.Scan(row); err != nil {
		return errors.Wrap(err, `scanning user record`)
	}
	u.key = u.Name

	return nil
}

// Update user data
func (u *User) Update(tx *sql.Tx) error {
	stmt := bytes.Buffer{}
	stmt.WriteString(`UPDATE `)
	stmt.WriteString(UserTable)
	stmt.WriteString(` SET username = ?, password = ? WHERE username = ?`)
	_, err := tx.Exec(stmt.String(), u.Name, u.Password, u.key)
	if err != nil {
		return errors.Wrap(err, `updating user record`)
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, `cammitting transaction`)
	}
	u.key = u.Name
	return nil
}

// Delete user data
func (u *User) Delete(tx *sql.Tx) error {
	stmt := bytes.Buffer{}
	stmt.WriteString(`DELETE FROM `)
	stmt.WriteString(UserTable)
	stmt.WriteString(` WHERE username=?`)
	_, err := tx.Exec(stmt.String(), u.Name)
	if err != nil {
		return errors.Wrap(err, `deleting user record`)
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, `committing transaction`)
	}
	u.key = ""
	u.Name = ""
	u.Password = ""
	u.CreatedOn = time.Time{}
	u.ModifiedOn = time.Time{}
	return nil
}
