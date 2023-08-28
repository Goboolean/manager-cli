package statusMock

import (
	"errors"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

var (
	errOnUpdating = errors.New("fetch-server: Fetch server returns errors on updating status")
	errOnGetting  = errors.New("fetch-server: Fetch server returns errors on getting status")
)

type StatusAdaptorMock struct {
	idToStatus map[string]entity.ProductStatus
}

func New() (*StatusAdaptorMock, error) {
	return &StatusAdaptorMock{
		idToStatus: make(map[string]entity.ProductStatus),
	}, nil
}
