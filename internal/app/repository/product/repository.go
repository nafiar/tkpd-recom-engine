package product

import m "github.com/nafiar/tkpd-recom-engine/internal/model/product"

// Repository for user info data source
type Repository interface {
	// GetData return user data based on user id in param
	GetProduct(param []m.Param) (m.Data, error)
}
