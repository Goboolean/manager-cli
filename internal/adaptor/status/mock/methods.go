package statusMock

import (
	"context"
	"errors"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method returns status of a product
func (a *StatusAdaptorMock) GetStatus(ctx context.Context, id string) (entity.ProductStatus, error) {
	res, ok := a.idToStatus[id]
	if !ok {
		return entity.ProductStatus{}, errors.New("product is not exist")
	}
	return res, nil
}

// This method changes status of a product by "status" val
func (a *StatusAdaptorMock) SetStatus(ctx context.Context, id string, status entity.ProductStatus) error {

	a.idToStatus[id] = status
	return nil
}
