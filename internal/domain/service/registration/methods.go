package registration

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// This method automatically registers a product by accepting the product's code and location information.
func (s *RegistrationService) RegisterProduct(meta entity.ProductMeta) error {

	var err error
	transactor := s.txCreator.CreateTransaction(context.TODO())

	err = s.metaRepo.StoreProductMeta(transactor.TransactionExtractor(), meta)

	if err != nil {
		err = transactor.Rollback()
		return err
	}

	err = transactor.Commit()
	return err

}
