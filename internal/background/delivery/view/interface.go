package nsq

import (
	"context"
)

type Delivery interface {
	Run(context.Context)
}
