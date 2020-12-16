package postgres

import (
	"context"
	"log"

	m "github.com/nafiar/tkpd-recom-engine/internal/model/product"
)

func (p *postgreProductInfo) GetProduct(param []m.Param) (result []m.Data, err error) {

	query := `select id, name, price from product`
	ctx := context.Background()
	rows, err := p.dbCon.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[getData] err executing scan : %+v\n", err)
	}

	for rows.Next() {
		var id int
		var name string
		var price float64

		err = rows.Scan(&id, &name, &price)
		if err != nil {
			log.Printf("[getData] err executing scan : %+v\n", err)
		}

		result = append(result, m.Data{
			ID:    id,
			Name:  name,
			Price: price,
		})
	}

	return
}
