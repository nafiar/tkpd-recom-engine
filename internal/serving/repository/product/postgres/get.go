package postgres

import (
	"context"
	"log"
	"strconv"

	m "github.com/nafiar/tkpd-recom-engine/internal/model/product"
)

func (p *postgreProductInfo) GetProduct(param []m.Param) (result []m.Data, err error) {

	query := `select id, name, price from product where `
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

// GetProductByIDs get product detail data by list of ids
// intentionally remove it's implementation
// reference : slide postgre `Process Query` & `Scan Data`
func (p *postgreProductInfo) GetProductByIDs(ids []int) (result []m.Data, err error) {
	return
}

// generateListStmt generate list for list of statement for id
func generateListStmt(ids []int) string {
	result := ""
	for i, _ := range ids {
		if i == 0 {
			result += `$`
		} else {
			result += `,$`
		}
		result += strconv.Itoa(i + 1)
	}
	return `(` + result + `)`
}
