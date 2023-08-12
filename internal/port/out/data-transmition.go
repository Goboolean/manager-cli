package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type DataTransmitterPort interface {
	//This method create remote directory
	CreateRemoteDir(dir string) error
	//This method transmit data to remote storage
	TransmitDataToRemote(localFile entity.File, remoteDir string) error
}
