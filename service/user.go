package service

import (
	"database/sql"
	"log"

	"github.com/charakoba-com/auth-api/db"
	"github.com/charakoba-com/auth-api/model"
	"github.com/pkg/errors"
)

// Create User
func (v *UserService) Create(tx *sql.Tx, du *db.User) error {
	log.Printf("service.User.Create %s", du.ID)

	if err := du.Create(tx); err != nil {
		return errors.Wrap(err, `creating db.User`)
	}
	return nil
}

// Lookup User
func (v *UserService) Lookup(tx *sql.Tx, id string) (*model.User, error) {
	log.Printf("service.User.Lookup %s", id)

	var mu model.User
	if err := mu.Load(tx, id); err != nil {
		return nil, errors.Wrap(err, `loading model.User`)
	}
	log.Printf("DONE: service.User.Loopkup")
	return &mu, nil
}

// Update User
func (v *UserService) Update(tx *sql.Tx, du *db.User) error {
	log.Printf("service.User.Update %s", du.ID)
	if err := du.Update(tx); err != nil {
		return errors.Wrap(err, `updating db.User`)
	}
	return nil
}

// Delete User
func (v *UserService) Delete(tx *sql.Tx, id string) error {
	log.Printf("service.User.Delete %s", id)

	du := db.User{ID: id}
	if err := du.Delete(tx); err != nil {
		return errors.Wrap(err, `deleting db.User`)
	}
	return nil
}

// Listup User
func (v *UserService) Listup(tx *sql.Tx) (model.UserList, error) {
	log.Printf("service.User.Listup")

	var userList db.UserList
	if err := userList.Listup(tx); err != nil {
		return nil, errors.Wrap(err, `loading user list`)
	}
	l := make(model.UserList, len(userList))
	for i, user := range userList {
		if err := l[i].FromDB(&user); err != nil {
			return nil, errors.Wrap(err, `converting db.User to model.User`)
		}
	}

	return l, nil
}
