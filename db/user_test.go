package db_test

import (
	"testing"

	"github.com/charakoba-com/auth-api/db"
)

func TestCreateUser(t *testing.T) {
	u := db.User{
		Name:     "Taro",
		Password: "none",
	}
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if err = u.Create(tx); err != nil {
		t.Errorf("%s", err)
		return
	}
	tx, err = db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	row := tx.QueryRow(`SELECT username, password, created_on, modified_on FROM users WHERE username='Taro'`)
	var scu db.User
	if err := scu.Scan(row); err != nil {
		t.Errorf("%s", err)
		return
	}
}

func TestLookupUser(t *testing.T) {
	u := db.User{}
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	err = u.Lookup(tx, "testuser")
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if u.Name != "testuser" {
		t.Errorf("%s != testuser", u.Name)
		return
	}
	if u.Password != "testpasswd" {
		t.Errorf("%s != testpasswd", u.Password)
		return
	}
}

func TestUpdateUser(t *testing.T) {
	u := db.User{}
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	err = u.Lookup(tx, "updateuser")
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	names := []string{"jiro", "saburo"}
	for _, name := range names {
		u.Name = name
		err = u.Update(tx)
		if err != nil {
			t.Errorf("%s", err)
			return
		}
		u = db.User{}
		tx, err = db.BeginTx()
		if err != nil {
			t.Errorf("%s", err)
			return
		}
		err = u.Lookup(tx, name)
		if err != nil {
			t.Errorf("%s", err)
			return
		}
		if u.Name != name {
			t.Errorf("%s != %s", u.Name, name)
			return
		}
	}
}

func TestDeleteUser(t *testing.T) {
	u := db.User{}
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	err = u.Lookup(tx, "deleteuser")
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	err = u.Delete(tx)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	err = u.Lookup(tx, "deleteuser")
	if err == nil {
		t.Errorf("%s should not be", u.Name)
		return
	}
}
