package registration

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method automatically registers a product by accepting the product's code and location information.
func (s RegistrationService) RegisterProduct(meta entity.ProductMeta) error {

	s.metaRepo.Begin(context.Background())

	err := s.metaRepo.StoreProductMeta(meta)

	if err != nil {
		s.metaRepo.Rollback()
		return err
	} else {
		s.metaRepo.Commit()
		return nil
	}
}
