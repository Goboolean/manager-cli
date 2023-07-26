package in

import (
	"time"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// abbreviation: command -> cmd

type RegCmdPort interface {

	// This method automatically registers a product by accepting the product's code and location information.
	RegisterProduct(meta entity.ProductMeta) error

	// This method automatically registers a product using the product's code and location information
	// completing the remaining metadata of the product.
	//TODO: Implement method below
	//AutoResisterProduct(id string) error
}

type BackupCmdPort interface {
	// This method backups all trade data to local
	BackupData() error
	// This method backups all trade data to local created before specific date
	BackupDataBefore(time time.Time) error
	// This method backups all trade data to Remote
	BackupDataToRemote() error
	// This method backups all trade data to local created before specific date
	BackupDataToRemoteBefore(time time.Time) error
	// This method backups data of a product to local.
	BackupProduct(id string) error
	// This method backups data of a product to local
	BackupProductToRemote(id string) error
}

// It has subscribe command
type StatusCmdPort interface {

	// SetStatus updates the status(es) of a product identified by its ID to desired status.
	SetStatus(id string, desired entity.ProductStatus) error
	// RemoveStatus removes status(es) from a product identified by its ID.
	RemoveStatus(id string, desired entity.ProductStatus) error
	// AddStatus adds a status(es) to a product identified by its ID.
	AddStatus(id string, desired entity.ProductStatus) error
	// get status of a product and returns status entity object
	GetStatus(id string) (entity.ProductStatus, error)
}
