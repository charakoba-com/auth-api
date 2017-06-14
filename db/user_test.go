package db_test

import (
	"testing"
	"time"

	"github.com/charakoba-com/auth-api/db"
)

func TestLoad(t *testing.T) {
	u := db.User{}
	db.Init(nil)
	tx, _ := db.BeginTx()
	if err := u.Load(tx, "lookupID"); err != nil {
		t.Errorf("%s", err)
		return
	}
	if u.ID != "lookupID" {
		t.Errorf("%s != lookupID", u.ID)
		return
	}
	if u.Name != "lookupuser" {
		t.Errorf("%s != lookupuser", u.Name)
		return
	}
	if u.Password != "testpasswd" {
		t.Errorf("%s != testpasswd", u.Password)
		return
	}
	loc, _ := time.LoadLocation("")
	exTime := time.Date(2017, 1, 1, 0, 0, 0, 0, loc)
	if u.CreatedOn != exTime {
		t.Errorf("%s != %s", u.CreatedOn, exTime)
		return
	}
	if !u.ModifiedOn.Valid {
		t.Errorf("%t", u.ModifiedOn.Valid)
		return
	}
	if u.ModifiedOn.Time != exTime {
		t.Errorf("%s != %s", u.ModifiedOn.Time, exTime)
		return
	}
}
