package info

import (
	"fmt"

	m "github.com/nafiar/tkpd-recom-engine/internal/model/user"
)

// GetData get user info data based on user id param
func (i *infoUseCase) GetData(userID string) (data m.Data, err error) {
	if i.userInfoRepo == nil {
		err = ErrNilRepository
		return
	}

	if userID == "" {
		err = ErrEmptyRequest
		return
	}

	data, err = i.userInfoRepo.GetData(userID)
	if err != nil {
		err = fmt.Errorf("failed to retrieve data, Err: %v", err.Error())
		return
	}

	return
}
