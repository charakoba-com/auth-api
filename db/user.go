package db

import (
	"bytes"
	"database/sql"
	"log"
	"time"

	"github.com/pkg/errors"
)

// Scan raw database row to user
func (u *User) Scan(scanner interface {
	Scan(...interface{}) error
}) error {
	return scanner.Scan(&u.ID, &u.Name, &u.Password, &u.CreatedOn, &u.ModifiedOn)
}

// Create User
func (u *User) Create(tx *sql.Tx) error {
	log.Printf("db.User.Create %s", u.ID)

	now := time.Now()

	stmt := bytes.Buffer{}
	stmt.WriteString(`INSERT INTO `)
	stmt.WriteString(userTable)
	stmt.WriteString(` (id, username, password, created_on) VALUES (?, ?, ?, ?)`)

	log.Printf("SQL QUERY: %s: with values %s, %s, %s, %s", stmt.String(), u.ID, u.Name, u.Password, now)

	_, err := tx.Exec(stmt.String(), u.ID, u.Name, u.Password, now)
	return err
}

// Load user data by user ID
func (u *User) Load(tx *sql.Tx, id string) error {
	log.Printf("db.User.Load %s", id)

	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(userSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(userTable)
	stmt.WriteString(` WHERE id = ?`)

	log.Printf("SQL QUERY: %s: with values %s", stmt.String(), id)

	row := tx.QueryRow(stmt.String(), id)

	if err := u.Scan(row); err != nil {
		return errors.Wrap(err, "scanning row")
	}
	return nil
}

// Update user
func (u *User) Update(tx *sql.Tx) error {
	if u.ID == "" {
		return errors.New(`user ID is not valid`)
	}
	log.Printf("db.User.Update %s", u.ID)

	stmt := bytes.Buffer{}
	stmt.WriteString(`UPDATE `)
	stmt.WriteString(userTable)
	stmt.WriteString(` SET username = ?, password = ? WHERE id = ?`)
	log.Printf("SQL QUERY: %s: with values %s, %s, %s", stmt.String(), u.Name, u.Password, u.ID)

	_, err := tx.Exec(stmt.String(), u.Name, u.Password, u.ID)

	return err
}

// Delete user from DB by user ID
func (u *User) Delete(tx *sql.Tx) error {
	if u.ID == "" {
		return errors.New(`user ID is not valid`)
	}
	log.Printf("db.User.Delete %s", u.ID)

	stmt := bytes.Buffer{}
	stmt.WriteString(`DELETE FROM `)
	stmt.WriteString(userTable)
	stmt.WriteString(` WHERE id = ?`)
	log.Printf("SQL QUERY: %s: with values %s", stmt.String(), u.ID)

	_, err := tx.Exec(stmt.String(), u.ID)

	return err
}
