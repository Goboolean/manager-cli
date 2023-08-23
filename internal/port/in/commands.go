package in

import (
	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// abbreviation: command -> cmd

type RegCmdPort interface {
	// This method automatically registers a product by accepting the product's code and location information.
	RegisterProduct(meta entity.ProductMeta) error
}

type BackupCmdPort interface {
	// BackupTradeFull backs up all trade data to the local storage.
	BackupTradeFull() error
	// BackupTradeDiff backs up the differential trade data of last full backup to the local storage.
	BackupTradeDiff() error
	// BackupTradeFullToRemote backs up all trade data to a remote storage.
	BackupTradeFullToRemote() error
	// BackupTradeDiffToRemote backs up the differential trade data to a remote storage.
	BackupTradeDiffToRemote() error
	// BackupProductFull backs up all trade data
	// related to a specific product (identified by 'id') to the local storage.
	BackupProductFull(id string) error
	// BackupProductDiff backs up the differential trade data of last full backup
	//related to a specific product (identified by 'id') to the local storage.
	BackupProductDiff(id string) error
	// BackupProductFullToRemote backs up all trade data
	// related to a specific product (identified by 'id') to a remote storage.
	BackupProductFullToRemote(id string) error
	// BackupProductDiffToRemote backs up the differential trade data
	//related to a specific product (identified by 'id') to a remote storage.
	BackupProductDiffToRemote(id string) error
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
