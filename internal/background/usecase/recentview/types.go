package recentview

import (
	"time"
)

type RecentViewSignal struct {
	UserID    int
	ProductID int
	Timestamp time.Time
}
