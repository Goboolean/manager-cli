package registration

import (
	"github.com/Goboolean/manager-cli/internal/port/out"
)

// define registration service struct
type RegistrationService struct {
	metaRepo out.MetadataRepositoryPort
	tx       out.TransactorPort
}

func New(metaPort out.MetadataRepositoryPort, tx out.TransactorPort) *RegistrationService {

	return &RegistrationService{
		metaRepo: metaPort,
		tx:       tx,
	}
}
