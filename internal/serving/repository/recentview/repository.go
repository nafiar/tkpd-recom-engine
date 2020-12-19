package recentview

// Repository is an interface for recent view user
type Repository interface {
	GetRecentView(userID int) (result []int, err error)
}
