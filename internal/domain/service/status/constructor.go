package status

import "github.com/Goboolean/manager-cli/internal/port/out"

type StatusService struct {
	status out.StatusPort
}

func New(statusPort out.StatusPort) *StatusService {
	return &StatusService{
		status: statusPort,
	}
}
