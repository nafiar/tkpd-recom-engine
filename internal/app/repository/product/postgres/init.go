package postgres

import (
	"database/sql"

	"github.com/nafiar/tkpd-recom-engine/internal/app/repository/product"
)

// postgreProduct
type postgreProductInfo struct {
	dbCon *sql.DB
}

// New
func New(db *sql.DB) product.Repository {
	return &postgreProductInfo{
		dbCon: db,
	}
}
