package in

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

// abbreviation: command -> cmd

type RegCmdPort interface {
	// This method automatically registers a product by accepting the product's code and location information.
	RegisterProduct(ctx context.Context, meta entity.ProductMeta) error
}

type BackupCmdPort interface {
	// BackupTradeFull backs up all trade data to the local storage.
	BackupTradeFull(ctx context.Context) error
	// BackupTradeDiff backs up the differential trade data of last full backup to the local storage.
	BackupTradeDiff(ctx context.Context) error
	// BackupTradeFullToRemote backs up all trade data to a remote storage.
	BackupTradeFullToRemote(ctx context.Context) error
	// BackupTradeDiffToRemote backs up the differential trade data to a remote storage.
	BackupTradeDiffToRemote(ctx context.Context) error
	// BackupProductFull backs up all trade data
	// related to a specific product (identified by 'id') to the local storage.
	BackupProductFull(ctx context.Context, id string) error
	// BackupProductDiff backs up the differential trade data of last full backup
	//related to a specific product (identified by 'id') to the local storage.
	BackupProductDiff(ctx context.Context, id string) error
	// BackupProductFullToRemote backs up all trade data
	// related to a specific product (identified by 'id') to a remote storage.
	BackupProductFullToRemote(ctx context.Context, id string) error
	// BackupProductDiffToRemote backs up the differential trade data
	//related to a specific product (identified by 'id') to a remote storage.
	BackupProductDiffToRemote(ctx context.Context, id string) error
}

// It has subscribe command
type StatusCmdPort interface {

	// SetStatus updates the status(es) of a product identified by its ID to desired status.
	SetStatus(ctx context.Context, id string, desired entity.ProductStatus) error
	// RemoveStatus removes status(es) from a product identified by its ID.
	RemoveStatus(ctx context.Context, id string, desired entity.ProductStatus) error
	// AddStatus adds a status(es) to a product identified by its ID.
	AddStatus(ctx context.Context, id string, desired entity.ProductStatus) error
	// get status of a product and returns status entity object
	GetStatus(ctx context.Context, id string) (entity.ProductStatus, error)
}
