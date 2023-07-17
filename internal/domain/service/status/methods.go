package status

import (
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

func (s StatusService) SetStatus(id string, desired entity.ProductStatus) error {
	//TODO: validate id
	return s.status.SetStatus(id, desired)
}

// RemoveStatus removes status(es) from a product identified by its ID.
func (s StatusService) RemoveStatus(id string, desired entity.ProductStatus) error {
	//TODO: validate id
	//TODO: validate desired

	current, err := s.status.GetStatus(id)

	if err != nil {
		return err
	}

	//  Find the state changed to through logical operation.
	//
	newStatus := entity.ProductStatus{
		Relayable:   current.Relayable && !desired.Relayable,
		Stored:      current.Stored && !desired.Stored,
		Transmitted: current.Transmitted && !desired.Transmitted,
	}

	return s.status.SetStatus(id, newStatus)
}

// AddStatus adds a status(es) to a product identified by its ID.
func (s StatusService) AddStatus(id string, desired entity.ProductStatus) error {
	//TODO: validate id

	current, err := s.status.GetStatus(id)

	if err != nil {
		return err
	}

	// Find the state to be changed through logical OR operation.
	newStatus := entity.ProductStatus{
		Relayable:   current.Relayable || desired.Relayable,
		Stored:      current.Stored || desired.Stored,
		Transmitted: current.Transmitted || desired.Transmitted,
	}

	return s.status.SetStatus(id, newStatus)

}

// get status of a product and returns status entity object
func (s StatusService) GetStatus(id string) (entity.ProductStatus, error) {
	return s.status.GetStatus(id)
}
