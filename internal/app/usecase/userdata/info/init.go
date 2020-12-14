package info

import (
	"github.com/nafiar/tkpd-recom-engine/internal/app/repository/userinfo"
	"github.com/nafiar/tkpd-recom-engine/internal/app/usecase/userdata"
)

// infoUseCase handler for user info retrieval
type infoUseCase struct {
	userInfoRepo userinfo.Repository
}

// New create new object handler for user info retrieval
func New(userInfo userinfo.Repository) userdata.UseCase {
	return &infoUseCase{
		userInfoRepo: userInfo,
	}
}
