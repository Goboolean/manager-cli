package registration

import (
	"github.com/Goboolean/manager-cli/internal/port/out"
)

// define registration service struct
type RegistrationService struct {
	metaRepo out.MetadataRepositoryPort
}

func New(metaPort out.MetadataRepositoryPort) *RegistrationService {

	return &RegistrationService{
		metaRepo: metaPort,
	}
}
