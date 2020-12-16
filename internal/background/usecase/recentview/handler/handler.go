package handler

import (
	"context"
	"log"

	"github.com/nafiar/tkpd-recom-engine/internal/background/repository/recentView"
	"github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview"
)

type simpleUsecase struct {
	recentViewRepo recentView.Repository
}

func New(recentViewRepo recentView.Repository) recentview.Usecase {
	return &simpleUsecase{
		recentViewRepo: recentViewRepo,
	}
}

func (s *simpleUsecase) Handle(ctx context.Context, signal recentview.RecentViewSignal) error {
	s.recentViewRepo.SetRecentView(signal.UserID, []int{signal.ProductID})
	return nil
}
