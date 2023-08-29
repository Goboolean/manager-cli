package out

import (
	"context"

	"github.com/Goboolean/manager-cli/internal/domain/entity"
)

type DataTransmitterPort interface {
	//This method create remote directory
	CreateRemoteDir(ctx context.Context, dir string) error
	//This method transmit data to remote storage
	TransmitDataToRemote(ctx context.Context, localFile entity.File, remoteDir string) error
}
