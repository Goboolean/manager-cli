//go:build wireinject
// +build wireinject

package inject

import (
	"os"

	backupMeta "github.com/Goboolean/manager-cli/internal/adaptor/backup-meta"
	"github.com/Goboolean/manager-cli/internal/adaptor/command"
	"github.com/Goboolean/manager-cli/internal/adaptor/file"
	productMetaRepo "github.com/Goboolean/manager-cli/internal/adaptor/product-meta-repo"
	statusAdaptor "github.com/Goboolean/manager-cli/internal/adaptor/status"
	tradeRepoDump "github.com/Goboolean/manager-cli/internal/adaptor/trade-repo-dump"
	transactionCreator "github.com/Goboolean/manager-cli/internal/adaptor/transaction-manager/transaction-creator"
	transmissionMock "github.com/Goboolean/manager-cli/internal/adaptor/transmission/mock"
	"github.com/Goboolean/manager-cli/internal/domain/service/backup"
	"github.com/Goboolean/manager-cli/internal/domain/service/registration"
	"github.com/Goboolean/manager-cli/internal/domain/service/status"
	fileInf "github.com/Goboolean/manager-cli/internal/infrastructure/file"
	grpcapi "github.com/Goboolean/manager-cli/internal/infrastructure/grpc/props"
	"github.com/Goboolean/manager-cli/internal/infrastructure/rdbms"
	"github.com/Goboolean/manager-cli/internal/port/in"
	"github.com/Goboolean/manager-cli/internal/port/out"
	"github.com/Goboolean/shared/pkg/mongo"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: separate get argus funcs to other struct or package
func getBackupOutDir() string {
	return os.Getenv("BACKUP_OUT_DIR")
}

func getMongoArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST":     os.Getenv("MONGO_HOST"),
		"USER":     os.Getenv("MONGO_USER"),
		"PORT":     os.Getenv("MONGO_PORT"),
		"PASSWORD": os.Getenv("MONGO_PASS"),
		"DATABASE": os.Getenv("MONGO_DATABASE"),
	}
}

func getPsqlArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST":     os.Getenv("PSQL_HOST"),
		"USER":     os.Getenv("PSQL_USER"),
		"PORT":     os.Getenv("PSQL_PORT"),
		"PASSWORD": os.Getenv("PSQL_PASS"),
		"DATABASE": os.Getenv("PSQL_DATABASE"),
	}
}

func getGrpcArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST": os.Getenv("STATUSAPI_HOST"),
		"PORT": os.Getenv("STATUSAPI_PORT"),
	}
}

func getBuycycleArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST": os.Getenv("BUYCYCLE_HOST"),
		"PORT": os.Getenv("BUYCYCLE_PORT"),
	}
}

func getPolygonArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"KEY": os.Getenv("POLYGON_API_KEY"),
	}
}

func getPrometheusArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"PORT": os.Getenv("METRIC_PORT"),
	}
}

func provideGrpcInfra() (grpcapi.StockConfiguratorClient, error) {

	c := getGrpcArgs()

	host, err := c.GetStringKey("HOST")
	if err != nil {
		return nil, err
	}

	port, err := c.GetStringKey("PORT")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return grpcapi.NewStockConfiguratorClient(conn), nil
}

func provideMongoInfra() (*mongo.DB, error) {
	c := getMongoArgs()

	return mongo.NewDB(c)
}

func provideRdbmsInfra() (*rdbms.PSQL, error) {
	c := getPsqlArgs()

	return rdbms.NewDB(c)
}

var InfraSet = wire.NewSet(
	fileInf.New,
	provideGrpcInfra,
	provideMongoInfra,
	provideRdbmsInfra,
)

func provideTradeRepoDumpAdaptor() (*tradeRepoDump.TradeDumpAdaptor, error) {
	c := getMongoArgs()
	return tradeRepoDump.New(c)
}

var AdaptorSet = wire.NewSet(
	backupMeta.New,
	command.New,
	file.New,
	productMetaRepo.New,
	statusAdaptor.New,
	provideTradeRepoDumpAdaptor,
	transactionCreator.New,
	transmissionMock.New,
	wire.Bind(new(out.BackupMetaPort), new(*backupMeta.BackupMetaAdaptor)),
	wire.Bind(new(out.DataTransmitterPort), new(*transmissionMock.TransmissionAdaptorMock)),
	wire.Bind(new(out.FileOperatorPort), new(*file.FileAdaptor)),
	wire.Bind(new(out.MetadataRepositoryPort), new(*productMetaRepo.MetadataRepositoryAdaptor)),
	wire.Bind(new(out.StatusPort), new(*statusAdaptor.StatusAdaptor)),
	wire.Bind(new(out.TradeDumperPort), new(*tradeRepoDump.TradeDumpAdaptor)),
	wire.Bind(new(out.TransactionCreator), new(*transactionCreator.TransactionFactory)),
)

var ServiceSet = wire.NewSet(
	backup.New,
	registration.New,
	status.New,
	getBackupOutDir,
	wire.Bind(new(in.BackupCmdPort), new(*backup.BackupService)),
	wire.Bind(new(in.RegCmdPort), new(*registration.RegistrationService)),
	wire.Bind(new(in.StatusCmdPort), new(*status.StatusService)),
)

func InitCommandAdaptor() (*command.CommandAdaptor, error) {
	wire.Build(InfraSet, AdaptorSet, ServiceSet)
	return &command.CommandAdaptor{}, nil
}
