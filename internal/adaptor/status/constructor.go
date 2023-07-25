package status

import (
	"context"
	"errors"
	"os"

	grpcapi "github.com/Goboolean/manager-cli/infrastructure/grpc/props"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	errOnUpdating = errors.New("fetch-server: Fetch server returns errors on updating status")
	errOnGetting  = errors.New("fetch-server: Fetch server returns errors on getting status")
)

type StatusAdaptor struct {
	client grpcapi.StockConfiguratorClient
	ctx    context.Context
}

func New() (*StatusAdaptor, error) {

	os.Chdir("../../..")
	godotenv.Load()

	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")

	conn, err := grpc.Dial(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	client := grpcapi.NewStockConfiguratorClient(conn)
	return &StatusAdaptor{
		client: client,
		ctx:    context.Background(),
	}, nil

}
