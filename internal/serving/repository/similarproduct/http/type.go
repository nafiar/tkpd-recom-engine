package http

type Response struct {
	AnchorProductID int64     `json:"anchor_product_id"`
	RelatedProducts []Product `json:"related_products"`
}

type Product struct {
	ProductID int64 `json:"product_id"`
}
