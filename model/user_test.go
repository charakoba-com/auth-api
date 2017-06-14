package model_test

import (
	"testing"

	"github.com/charakoba-com/auth-api/db"
	"github.com/charakoba-com/auth-api/model"
)

func TestLoad(t *testing.T) {
	db.Init(nil)

	testID := "lookupID"
	testUsername := "lookupuser"
	testPassword := "testpasswd"

	u := model.User{}
	tx, _ := db.BeginTx()
	if err := u.Load(tx, testID); err != nil {
		t.Errorf("%s", err)
		return
	}
	if u.ID != testID {
		t.Errorf("%s != %s", u.ID, testID)
		return
	}
	if u.Name != testUsername {
		t.Errorf("%s != %s", u.Name, testUsername)
		return
	}
	if u.Password != testPassword {
		t.Errorf("%s != %s", u.Password, testPassword)
		return
	}
}

func TestFromDB(t *testing.T) {
	dbRow := db.User{
		ID:       "testID",
		Name:     "testName",
		Password: "testPasswd",
	}
	u := model.User{}
	if err := u.FromDB(&dbRow); err != nil {
		t.Errorf("FromDB does NOT return error...: %s", err)
		return
	}
	if u.ID != dbRow.ID {
		t.Errorf("%s != %s", u.ID, dbRow.ID)
		return
	}
	if u.Name != dbRow.Name {
		t.Errorf("%s != %s", u.Name, dbRow.Name)
		return
	}
	if u.Password != dbRow.Password {
		t.Errorf("%s != %s", u.Password, dbRow.ID)
		return
	}
}

func TestToDB(t *testing.T) {
	u := model.User{
		ID: "testID",
		Name: "testName",
		Password: "testPasswd",
	}
	du := db.User{}
	if err := u.ToDB(&du); err != nil {
		t.Errorf("FromDB does NOT return error...: %s", err)
		return
	}
	if du.ID != u.ID {
		t.Errorf("%s != %s", du.ID, u.ID)
		return
	}
	if du.Name != u.Name {
		t.Errorf("%s != %s", du.Name, u.Name)
		return
	}
	if du.Password != u.Password {
		t.Errorf("%s != %s", du.Password, u.ID)
		return
	}
}
