package view

import (
	"github.com/nafiar/tkpd-recom-engine/internal/background/repository/recentView"
	"github.com/nafiar/tkpd-recom-engine/internal/background/usecase/useractivity"
)

type viewUserActivity struct {
	recentViewRepo recentView.Repository
}

// New create new object for user activity view handler
func New(recentViewRepo recentView.Repository) useractivity.UseCase {
	return &viewUserActivity{
		recentViewRepo: recentViewRepo,
	}
}
