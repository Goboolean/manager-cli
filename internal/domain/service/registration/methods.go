package registration

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method automatically registers a product by accepting the product's code and location information.
func (s RegistrationService) RegisterProduct(meta entity.ProductMeta) error {

	ctx, _ := context.WithCancel(context.Background())

	s.tx.Begin(ctx)
	err := s.metaRepo.StoreProductMeta(meta)

	if err != nil {
		s.tx.Rollback()
		return err
	} else {
		s.tx.Commit()
		return nil
	}
}
