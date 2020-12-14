package info

import (
	"fmt"
)

var (
	// ErrNilRepository - error type for uninitialized user info repository
	ErrNilRepository = fmt.Errorf("User info data source is not initialized")

	// ErrEmptyRequest - error type for empty request
	ErrEmptyRequest = fmt.Errorf("User id can not be empty")
)
