package registration

import (
	"github.com/Goboolean/manager-cli/internal/port/out"
)

// define registration service struct
type RegistrationService struct {
	txCreator out.TransactionCreator
	metaRepo  out.MetadataRepositoryPort
}

func New(
	transactionCreator out.TransactionCreator,
	metaPort out.MetadataRepositoryPort) *RegistrationService {

	return &RegistrationService{
		txCreator: transactionCreator,
		metaRepo:  metaPort,
	}
}
