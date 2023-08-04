package out

import "github.com/Goboolean/manager-cli/internal/domain/entity"

type DataTransmitterPort interface {
	//This method transmit data to remote storage
	TransmitDataToRemote(file []entity.FileManager) error
}
