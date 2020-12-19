package product

import (
	"fmt"
)

// safeguard validate unexpected conditions
func (p *productRecommendation) safeguard() error {
	if p.recentViewRepo == nil {
		return fmt.Errorf("user recent view activity data source is not available")
	}

	if p.similarProductRepo == nil {
		return fmt.Errorf("similar product data source is not available")
	}

	if p.productRepo == nil {
		return fmt.Errorf("product detail data source available")
	}

	return nil
}
