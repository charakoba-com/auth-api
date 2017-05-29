package db_test

import (
	"testing"

	"github.com/builderscon/octav/octav/db"
)

func TestDBConnection(t *testing.T) {
	err := db.Init(nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if err = db.DB.Ping(); err != nil {
		t.Errorf("%s", err)
		return
	}
}
