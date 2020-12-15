package handler

import (
	"context"
	"log"

	"github.com/nafiar/tkpd-recom-engine/internal/background/usecase/recentview"
)

type simpleLogic struct{}

func New() recentview.Usecase {
	return &simpleLogic{}
}

func (s *simpleLogic) Handle(ctx context.Context, signal recentview.RecentViewSignal) error {
	log.Printf("executed")
	return nil
}
