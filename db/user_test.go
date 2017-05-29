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
	err = u.Create(tx)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
}
