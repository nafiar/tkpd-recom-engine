package recentview

type UseCase interface {
	SetRecentView(messageBody []byte) (err error)
}
