package product

// Data store Product model data
type Data struct {
	ID   int
	Name string
	Price  float64
}

// Param store Product model data
type Param struct {
	ID         int
	CategoryID int
}
