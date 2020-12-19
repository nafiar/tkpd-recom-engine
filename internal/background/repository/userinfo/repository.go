package userinfo

import m "github.com/nafiar/tkpd-recom-engine/internal/model/user"

// Repository for user info data source
type Repository interface {
	// GetData return user data based on user id in param
	GetData(userID string) (m.Data, error)
}
