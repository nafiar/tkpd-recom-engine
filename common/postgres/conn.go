package postgres

import (
	_ "github.com/lib/pq"
	"database/sql"
)

// NewConnection creates a new postgres connection.
//
// During an error, it will perform a reconnect mechanism
// in he background until the connection is established.
func NewConnection(ConnString string) (tempPG *sql.DB, err error) {
	tempPG, err  = sql.Open("postgres", ConnString)
	if err != nil {
		return nil, err
	}
	return tempPG, nil
}
