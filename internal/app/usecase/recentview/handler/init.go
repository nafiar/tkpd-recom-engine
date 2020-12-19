package handler

import (
	"github.com/nafiar/tkpd-recom-engine/internal/app/repository/recentView"
	"github.com/nafiar/tkpd-recom-engine/internal/app/usecase/userdata"
)

type recentViewUseCase struct {
	recentViewRepo recentView.Repository
}

// New create new object handler for user recent view retrieval
func New(recentView recentView.Repository) userdata.UseCase {
	return &recentViewUseCase{
		recentViewRepo: recentView,
	}
}
