package postgres

import (
	_ "github.com/lib/pq"
	"database/sql"
)

const (
	ConnString = "postgres://root:root12345@localhost:5432/tokopedia-product-recommendation?sslmode=disable"
)
// NewConnection creates a new postgres connection.
//
// During an error, it will perform a reconnect mechanism
// in he background until the connection is established.
func NewConnection() (tempPG *sql.DB, err error) {
	tempPG, err  = sql.Open("postgres", ConnString)
	if err != nil {
		return nil, err
	}
	return tempPG, nil
}
