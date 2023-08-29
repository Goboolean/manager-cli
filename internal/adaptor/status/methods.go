package status

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
	grpcapi "github.com/Goboolean/manager-cli/internal/infrastructure/grpc/props"
)

// This method returns status of a product
func (a *StatusAdaptor) GetStatus(ctx context.Context, id string) (entity.ProductStatus, error) {

	config, err := a.client.GetStockConfigOne(ctx, &grpcapi.StockId{
		StockId: id,
	})

	if err != nil {
		return entity.ProductStatus{}, err
	}

	if config.Relayable.OptionStatus > 1 || config.Storeable.OptionStatus > 1 || config.Transmittable.OptionStatus > 1 {

		return entity.ProductStatus{}, errOnGetting
	}

	return entity.ProductStatus{
		Relayable:   config.Relayable.OptionStatus == 1,
		Transmitted: config.Transmittable.OptionStatus == 1,
		Stored:      config.Storeable.OptionStatus == 1,
	}, nil
}

// This method changes status of a product by "status" val
func (a *StatusAdaptor) SetStatus(ctx context.Context, id string, status entity.ProductStatus) error {

	msg, err := a.client.UpdateStockConfigOne(ctx, &grpcapi.StockConfig{
		StockId:       id,
		Relayable:     &grpcapi.OptionStatus{OptionStatus: mapBoolToOption(status.Relayable)},
		Storeable:     &grpcapi.OptionStatus{OptionStatus: mapBoolToOption(status.Stored)},
		Transmittable: &grpcapi.OptionStatus{OptionStatus: mapBoolToOption(status.Transmitted)},
	})

	if err != nil {
		return err
	}

	if !msg.Status {
		return errOnUpdating
	}

	return nil
}

// This method maps bool to int32 type which gRPC api requires
func mapBoolToOption(in bool) int32 {
	if in {
		return 1
	} else {
		return 0
	}
}
