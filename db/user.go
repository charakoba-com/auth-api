package db

import (
	"bytes"
	"database/sql"
	"log"
	"time"

	"github.com/charakoba-com/auth-api/utils"
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

	// hash user's password
	hashed := utils.HashPassword(u.Password, u.ID + u.Name)

	_, err := tx.Exec(stmt.String(), u.ID, u.Name, hashed, now)
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

	// hash user's password
	hashed := utils.HashPassword(u.Password, u.ID + u.Name)

	_, err := tx.Exec(stmt.String(), u.Name, hashed, u.ID)

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

// Listup Users
func (l *UserList)Listup(tx *sql.Tx) error {
	log.Printf("db.User.Listup")

	stmt := bytes.Buffer{}
	stmt.WriteString(`SELECT `)
	stmt.WriteString(userSelectColumns)
	stmt.WriteString(` FROM `)
	stmt.WriteString(userTable)

	log.Printf("SQL QUERY: %s", stmt.String())

	rows, err := tx.Query(stmt.String())
	if err != nil {
		return errors.Wrap(err, `querying stmt`)
	}
	if err:= l.FromRows(rows); err != nil {
		return errors.Wrap(err, "scanning rows")
	}
	return nil
}

// FromRows scanning rows into user list
func (l *UserList) FromRows(rows *sql.Rows) error {
	log.Printf("db.User.FromRows")

	res := UserList{}

	for rows.Next() {
		user := User{}
		if err := user.Scan(rows); err != nil {
			return errors.Wrap(err, `scanning row`)
		}
		res = append(res, user)
	}
	*l = res
	return nil
}
