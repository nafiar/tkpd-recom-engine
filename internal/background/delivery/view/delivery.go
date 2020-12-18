package nsq

import (
	"context"
)

// // Delivery interface Message Consumer
type Delivery interface {
	Run(context.Context)
}
