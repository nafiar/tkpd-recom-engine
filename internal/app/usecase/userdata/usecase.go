package userdata

import (
	m "github.com/nafiar/tkpd-recom-engine/internal/model/user"
)

// UseCase interface for user data logic
type UseCase interface {
	GetData(userID string) (m.Data, error)
}
