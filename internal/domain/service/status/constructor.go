package status

import (
	"errors"

	"github.com/Goboolean/manager-cli/internal/port/out"
)

var InvalidStatus = errors.New("update status: invalid status")

type StatusService struct {
	status out.StatusPort
}

func New(statusPort out.StatusPort) *StatusService {
	return &StatusService{
		status: statusPort,
	}
}
