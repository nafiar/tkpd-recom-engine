package recentView

type Repository interface {
	GetRecentView(userID int) (result []int, err error)
	SetRecentView(userID int, productIDs []int) (err error)
}
