package similarproduct

// Repository is interface for graph embedding model
type Repository interface {
	// GetSimilarProduct get similar products from a single product id
	GetSimilarProduct(id []int) (result []int, err error)
}
