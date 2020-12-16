package recentview

import (
	"context"
)

// Usecase represent recent view business logic
type Usecase interface {
	Handle(ctx context.Context, signal RecentViewSignal) error
}
