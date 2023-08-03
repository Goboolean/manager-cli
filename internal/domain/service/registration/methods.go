package registration

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method automatically registers a product by accepting the product's code and location information.
func (s RegistrationService) RegisterProduct(meta entity.ProductMeta) error {

	sess, err := s.metaRepo.CreateTxSession(context.TODO())

	if err != nil {
		return err
	}

	err = s.metaRepo.StoreProductMeta(sess, meta)

	if err != nil {
		s.metaRepo.Rollback(sess)
		return err
	} else {
		s.metaRepo.Commit(sess)
		return nil
	}
}
