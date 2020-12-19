package postgres

import (
	"database/sql"

	"github.com/nafiar/tkpd-recom-engine/internal/serving/repository/product"
)

// postgreProduct
type postgreProductInfo struct {
	dbCon *sql.DB
}

var (
	stmtGetProductByIDs *sql.Stmt
)

// New
func New(db *sql.DB) product.Repository {
	return &postgreProductInfo{
		dbCon: db,
	}
}
