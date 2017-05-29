package db

import "time"

// User represents API user including admin and regular user
type User struct {
	key        string
	Name       string
	Password   string
	CreatedOn  time.Time
	ModifiedOn time.Time
}
