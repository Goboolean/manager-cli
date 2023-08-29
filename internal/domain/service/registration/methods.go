package registration

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method automatically registers a product by accepting the product's code and location information.
func (s *RegistrationService) RegisterProduct(ctx context.Context, meta entity.ProductMeta) error {

	var err error
	transactor, err := s.txCreator.CreateTransaction(ctx)
	if err != nil {
		return err
	}

	err = s.metaRepo.StoreProductMeta(ctx, transactor.TransactionExtractor(), meta)

	if err != nil {
		err = transactor.Rollback()
		return err
	}

	err = transactor.Commit()
	return err

}
