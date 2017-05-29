package db_test

import (
	"testing"

	"github.com/charakoba-com/auth-api/db"
)

func TestDBConnection(t *testing.T) {
	err := db.Init(nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	_, err = db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
}
