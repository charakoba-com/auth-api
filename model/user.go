package model

import (
	"database/sql"
	"log"

	"github.com/charakoba-com/auth-api/db"
	"github.com/pkg/errors"
)

// Load with user ID
func (u *User) Load(tx *sql.Tx, id string) (err error) {
	log.Printf("model.User.Load %s", id)

	du := db.User{}
	if err := du.Load(tx, id); err != nil {
		return errors.Wrap(err, "loading db.User")
	}

	if err := u.FromDB(&du); err != nil {
		return errors.Wrap(err, "scanning db.User")
	}
	return nil
}

// FromDB binds db.User to model.User
func (u *User) FromDB(du *db.User) error {
	log.Printf("model.User.FromDB")
	u.ID = du.ID
	u.Name = du.Name
	u.Password = du.Password
	return nil
}

// ToDB binds model.User to db.User
func (u *User) ToDB(du *db.User) error {
	log.Printf("model.User.ToDB")
	du.ID = u.ID
	du.Name = u.Name
	du.Password = u.Password
	return nil
}

// sort.Interface implementation

// Len returns the number of elements
func (ul UserList) Len() int {
	return len(ul)
}

// Swap two elements
func (ul UserList) Swap(i, j int) {
	ul[i], ul[j] = ul[j], ul[i]
}

// Less for sort
func (ul UserList) Less(i, j int) bool {
	return ul[i].ID < ul[j].ID
}
