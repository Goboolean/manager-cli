package status

import (
	"errors"

	grpcapi "github.com/Goboolean/manager-cli/internal/infrastructure/grpc/props"
)

var (
	errOnUpdating = errors.New("fetch-server: Fetch server returns errors on updating status")
	errOnGetting  = errors.New("fetch-server: Fetch server returns errors on getting status")
)

type StatusAdaptor struct {
	client grpcapi.StockConfiguratorClient
}

func New(client grpcapi.StockConfiguratorClient) *StatusAdaptor {
	return &StatusAdaptor{
		client: client,
	}

}
