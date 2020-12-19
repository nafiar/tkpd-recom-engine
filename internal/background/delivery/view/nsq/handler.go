package nsq

import (
	"context"
	"log"
	"strings"
)

// Run will open NSQ consummer connection and start consuming data
// initially it will Split lookupaddress into list
// then use it as connection address
func (v *viewListener) Run(ctx context.Context) {
	// safeguard for nil consumer
	if v.consumer == nil {
		log.Println("consumer cannot be nil")
		return
	}

	lookupds := strings.Split(v.lookupdAddress, ",")
	err := v.consumer.ConnectToNSQLookupds(lookupds)
	if err != nil {
		log.Println(err)
	}
}
